package app

import (
	"context"
	"encoding/base64"
	pb_comments "fivi/gen/go/comments/v1"
	pb_followers "fivi/gen/go/followers/v1"
	pb_likes "fivi/gen/go/likes/v1"
	pb_posts "fivi/gen/go/posts/v1"
	profilepb "fivi/gen/go/profile/v1"
	"fivi/lib/jwt"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/oklog/run"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultGrpcMessageSize = 100 * 1024 * 1024

type Config struct {
	HttpPublicServerPort int `env:"HTTP_PUBLIC_SERVER_PORT,notEmpty"`

	ProfilesGrpcServerAddr  string `env:"USERS_GRPC_SERVER_ADDR,notEmpty"`
	CommentsGrpcServerAddr  string `env:"API_TOKENS_GRPC_SERVER_ADDR,notEmpty"`
	PostsGrpcServerAddr     string `env:"WEB_HOOKS_GRPC_SERVER_ADDR,notEmpty"`
	FollowersGrpcServerAddr string `env:"LIGHTNING_GRPC_SERVER_ADDR,notEmpty"`
	LikesGrpcServerAddr     string `env:"BITCOIN_GRPC_SERVER_ADDR,notEmpty"`

	StripPrefix   string        `env:"STRIP_PREFIX,required"`
	JwtSigningKey string        `env:"JWT_SIGNING_KEY,notEmpty"`
	JwtTTL        time.Duration `env:"JWT_TTL,notEmpty"`
	StaticDir     string
	ImagesDir     string
}

type App struct {
	cfg *Config

	ji *jwt.JWT

	templates struct {
		index *template.Template
	}

	shutdown bool
	done     chan struct{}
}

func New() (*App, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, errors.Wrap(err, "can't get Config from env vars")
	}

	app, err := WithConfig(cfg)

	return app, err
}

func WithConfig(cfg *Config) (*App, error) {
	jwtSigningKey, err := base64.StdEncoding.DecodeString(cfg.JwtSigningKey)
	if err != nil {
		return nil, errors.Errorf("can't base64 decode jwt signing key: %v", err)
	}

	ji := jwt.NewInteractor(jwtSigningKey, cfg.JwtTTL)

	return &App{
		cfg:  cfg,
		ji:   ji,
		done: make(chan struct{}),
	}, nil
}

func (a *App) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := a.initializeTemplates()
	if err != nil {
		log.Println("TEMPLATES")
		return
	}

	log.Info("The gateway-service starting")

	httpServerListen := fmt.Sprintf("127.0.0.1:%v", a.cfg.HttpPublicServerPort)

	var g run.Group

	{ // HTTP endpoints.
		grpcMux := runtime.NewServeMux(
			a.ji.AuthMiddleware(),
		)
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(defaultGrpcMessageSize),
				grpc.MaxCallSendMsgSize(defaultGrpcMessageSize),
			),
		}

		if err := profilepb.RegisterProfileServiceHandlerFromEndpoint(ctx, grpcMux, a.cfg.ProfilesGrpcServerAddr, opts); err != nil {
			log.Println("failed to start profiles service: %v", err)
			return
		}

		if err := pb_comments.RegisterServiceHandlerFromEndpoint(ctx, grpcMux, a.cfg.CommentsGrpcServerAddr, opts); err != nil {
			log.Println("failed to start profiles service: %v", err)
			return
		}

		if err := pb_likes.RegisterServiceHandlerFromEndpoint(ctx, grpcMux, a.cfg.LikesGrpcServerAddr, opts); err != nil {
			log.Println("failed to start profiles service: %v", err)
			return
		}

		if err := pb_followers.RegisterFollowersServiceHandlerFromEndpoint(ctx, grpcMux, a.cfg.FollowersGrpcServerAddr, opts); err != nil {
			log.Println("failed to start profiles service: %v", err)
			return
		}

		if err := pb_posts.RegisterServiceHandlerFromEndpoint(ctx, grpcMux, a.cfg.PostsGrpcServerAddr, opts); err != nil {
			log.Println("failed to start profiles service: %v", err)
			return
		}

		router := mux.NewRouter()

		router.PathPrefix("/gateway/").Handler(
			http.StripPrefix(a.cfg.StripPrefix, grpcMux),
		)

		router.PathPrefix("/docs/").Handler(
			http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs/"))),
		)

		imagesServer := http.FileServer(http.Dir(a.cfg.ImagesDir))
		router.PathPrefix("/images/").Handler(http.StripPrefix("/images", imagesServer))

		fs := http.FileServer(http.Dir(a.cfg.StaticDir))
		router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
		router.PathPrefix("/").HandlerFunc(a.appHandler)

		// Start HTTP server (and proxy calls to gRPC server endpoint)
		httpServer := &http.Server{
			Addr:    httpServerListen,
			Handler: allowCORS(router),
		}

		g.Add(func() error {
			log.Info("Start HTTP endpoints")
			if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}
			return nil
		}, func(err error) {
			log.Info("Stop HTTP endpoints\n")
			e := httpServer.Shutdown(context.Background())
			if e != nil {
				log.Errorf("Could not gracefully shut down the server: %v", e)
			}
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

	log.Infof("The gateway-service was terminated with: %v", g.Run())
}

func (a *App) Shutdown() {
	if a.shutdown {
		return
	}
	a.shutdown = true

	close(a.done)
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, OPTIONS, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == http.MethodOptions {
			return
		}

		h.ServeHTTP(w, r)
	})
}

// appHandler is web app http handler function.
func (a *App) appHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Set("Content-Type", "text/html; charset=UTF-8")
	// header.Set("X-Content-Type-Options", "nosniff").
	header.Set("Referrer-Policy", "same-origin")

	if a.templates.index == nil {
		log.Println("index template is not set", nil)
		return
	}

	if err := a.templates.index.Execute(w, nil); err != nil {
		log.Println("index template could not be executed", err)
		return
	}
}

// initializeTemplates is used to initialize all templates.
func (a *App) initializeTemplates() (err error) {
	a.templates.index, err = template.ParseFiles(filepath.Join(a.cfg.StaticDir, "dist", "index.html"))
	if err != nil {
		log.Println("dist folder is not generated. use 'npm run build' command", err)
		return err
	}
	log.Println("initialized")

	return nil
}
