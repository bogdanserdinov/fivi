package app

import (
	"context"
	"database/sql"
	pb_posts "fivi/gen/go/posts/v1"
	store2 "fivi/lib/store"
	comments_client_grpc "fivi/services/comments/v1/client/grpc"
	followers_client_grpc "fivi/services/followers/v1/client/grpc"
	likes_client_grpc "fivi/services/likes/v1/client/grpc"
	v1 "fivi/services/posts/v1"
	"fivi/services/posts/v1/repository"
	profile_client_grpc "fivi/services/profile/v1/client/grpc"
	"fmt"
	"net"

	"github.com/caarlos0/env/v6"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const defaultGrpcMessageSize = 100 * 1024 * 1024

type Config struct {
	GrpcServerPort      int    `env:"GRPC_SERVER_PORT,notEmpty"`
	DBConnString        string `env:"DATABASE_URL,notEmpty"`
	ImagesDir           string `env:"DATABASE_URL,notEmpty"`
	ProfileServerAddr   string `env:"DATABASE_URL,notEmpty"`
	CommentsServerAddr  string `env:"DATABASE_URL,notEmpty"`
	LikesServerAddr     string `env:"DATABASE_URL,notEmpty"`
	FollowersServerAddr string
}

func (cfg *Config) ToMap() map[string]string {
	return map[string]string{
		"GRPC_SERVER_PORT": fmt.Sprintf("%d", cfg.GrpcServerPort),
		"DATABASE_URL":     cfg.DBConnString,
	}
}

type app struct {
	cfg      *Config
	shutdown bool
	done     chan struct{}
}

func New() (*app, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, errors.Wrap(err, "can't get Config from env vars")
	}
	return WithConfig(cfg), nil
}

func WithConfig(cfg *Config) *app {
	return &app{
		cfg:  cfg,
		done: make(chan struct{}),
	}
}

func (a *app) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	log.Info("The posts-service starting")
	grpcServerListen := fmt.Sprintf("0.0.0.0:%v", a.cfg.GrpcServerPort)
	var g run.Group
	db, err := sql.Open("postgres", a.cfg.DBConnString)
	if err != nil {
		log.Fatalf("can't open connection to postgres: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Errorf("can't close connection to postgres: %v", err)
		}
	}()
	// GRPC endpoints
	{
		grpcServer := grpc.NewServer(
			grpc.MaxRecvMsgSize(defaultGrpcMessageSize),
			grpc.MaxSendMsgSize(defaultGrpcMessageSize),
		)

		repository, err := repository.Prepare(context.Background(), db)
		if err != nil {
			log.Fatalf("can't prepare posts repository: %v", err)
		}

		store := store2.NewStore(store2.Config{
			OutputPath: a.cfg.ImagesDir,
		})

		profilesClient, err := profile_client_grpc.NewProfileServiceClient(a.cfg.ProfileServerAddr)
		if err != nil {
			log.Fatalf("can't connect to profiles server: %v", err)
		}

		commentsClient, err := comments_client_grpc.NewCommentsServiceClient(a.cfg.CommentsServerAddr)
		if err != nil {
			log.Fatalf("can't connect to comments server: %v", err)
		}

		likesClient, err := likes_client_grpc.NewLikesServiceClient(a.cfg.LikesServerAddr)
		if err != nil {
			log.Fatalf("can't connect to likes server: %v", err)
		}

		followersClient, err := followers_client_grpc.NewFollowersServiceClient(a.cfg.FollowersServerAddr)
		if err != nil {
			log.Fatalf("can't connect to followers server: %v", err)
		}

		likesServer := v1.New(repository, store, profilesClient, commentsClient, likesClient, followersClient)
		if err != nil {
			log.Fatalf("can't create posts server: %v\n", err)
		}

		pb_posts.RegisterServiceServer(grpcServer, likesServer)

		g.Add(func() error {
			log.Info("Start GRPC endpoints")

			lis, err := net.Listen("tcp", grpcServerListen)
			if err != nil {
				return fmt.Errorf("failed to listen: %v", err)
			}
			return grpcServer.Serve(lis)
		}, func(err error) {
			log.Info("Stop GRPC endpoints")
			grpcServer.GracefulStop()
			cancel()
		})
	}

	{
		g.Add(func() error {
			<-a.done
			return nil
		}, func(err error) {
			a.Shutdown()
		})
	}
	log.Infof("posts-service was terminated with: %v", g.Run())
}

func (a *app) Shutdown() {
	if a.shutdown {
		return
	}
	a.shutdown = true
	close(a.done)
}
