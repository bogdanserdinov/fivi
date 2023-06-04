package app

import (
	"context"
	"database/sql"
	"encoding/base64"
	pb_profile "fivi/gen/go/profile/v1"
	"fivi/lib/jwt"
	store2 "fivi/lib/store"
	"fivi/services/profile/v1"
	"fivi/services/profile/v1/repository"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)

const defaultGrpcMessageSize = 100 * 1024 * 1024

type Config struct {
	GrpcServerPort int    `env:"GRPC_SERVER_PORT,notEmpty"`
	DBConnString   string `env:"DATABASE_URL,notEmpty"`
	ImagesDir      string `env:"DATABASE_URL,notEmpty"`
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

	log.Info("The profile-service starting")
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
			log.Fatalf("can't prepare profile repository: %v", err)
		}

		jwtSigningKey, err := base64.StdEncoding.DecodeString("Uv38ByGCZU8WP18PmmIdcpVmx00QA3xNe7sEB9Hixkk=")
		if err != nil {
			log.Fatalf("can't base64 decode jwt signing key: %v", err)
		}
		ji := jwt.NewInteractor(jwtSigningKey, 24*time.Hour)

		store := store2.NewStore(store2.Config{
			OutputPath: a.cfg.ImagesDir,
		})

		likesServer := v1.New(ji, repository, store)
		if err != nil {
			log.Fatalf("can't create posts server: %v\n", err)
		}
		pb_profile.RegisterProfileServiceServer(grpcServer, likesServer)

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
	log.Infof("profile-service was terminated with: %v", g.Run())
}

func (a *app) Shutdown() {
	if a.shutdown {
		return
	}
	a.shutdown = true
	close(a.done)
}
