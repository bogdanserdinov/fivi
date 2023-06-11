//go:build mage

package main

import (
	"context"
	"database/sql"
	comments_cmd "fivi/services/comments/v1/cmd/app"
	followers_cmd "fivi/services/followers/v1/cmd/app"
	gateway_cmd "fivi/services/gateway/v1/cmd/gateway/app"
	likes_cmd "fivi/services/likes/v1/cmd/app"
	posts_cmd "fivi/services/posts/v1/cmd/app"
	profile_cmd "fivi/services/profile/v1/cmd/app"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
	_ "github.com/lib/pq"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"golang.org/x/sync/errgroup"
)

type service struct {
	name          string
	version       string
	containerName string
}

func (s *service) nameWithVersion() string {
	if s.version == "" {
		return s.name
	}
	return fmt.Sprintf("%v/%v", s.name, s.version)
}

var (
	// TODO(evg): extract service names automatically?
	services = []*service{
		{name: "lightning", version: "v1", containerName: "lightning-service"},
		{name: "profile", version: "v1", containerName: "profile-service"},
		//{name: "did", version: "v1", containerName: "did-service"},
		{name: "gateway", version: "v1", containerName: "gateway-service"},
		{name: "comments", version: "v1", containerName: "comments-service"},
		{name: "posts", version: "v1", containerName: "posts-service"},
		{name: "likes", version: "v1", containerName: "likes-service"},
		{name: "followers", version: "v1", containerName: "followers-service"},
		{name: "lightning_watcher", version: "v1", containerName: "lightning-watcher-service"},
	}

	itestsTags    = []string{"mock_coingecko", "mock_postmark", "mock_s3", "mock_did", "mock_firebase"}
	itestsTagsOpt = fmt.Sprintf("-tags=%v", strings.Join(itestsTags, ","))
)

func Up() error {
	return sh.Run("docker-compose", "up", "-d")
}

func UpAndScale(numOfInstances int) error {
	args := []string{"up", "-d"}

	for _, service := range services {
		if service.name != "lightning_watcher" && service.name != "scheduler" {
			args = append(args, "--scale", fmt.Sprintf("%v=%v", service.containerName, numOfInstances))
		}
	}

	return sh.Run("docker-compose", args...)
}

func Down() error {
	return sh.Run("docker-compose", "down", "-v", "--rmi=local", "--remove-orphans")
}

func getContainerNames() []string {
	var containers []string

	for _, service := range services {
		if service.name == "lightning_watcher" {
			containers = append(containers, "lightning-watcher-service1", "lightning-watcher-service2")
		} else {
			containers = append(containers, service.containerName)
		}
	}

	return containers
}

// Use: mage stopApplicationContainers
func StopApplicationContainers() error {
	containers := getContainerNames()
	args := []string{"stop"}
	args = append(args, containers...)

	if err := sh.Run("docker-compose", args...); err != nil {
		return err
	}

	return nil
}

// Use: mage restartService
func RestartService(serviceName string) error {
	if err := sh.Run("docker-compose", "stop", serviceName); err != nil {
		return err
	}
	if err := sh.Run("docker-compose", "rm", "-y", serviceName); err != nil {
		return err
	}
	if err := sh.Run("docker-compose", "up", "-d", serviceName); err != nil {
		return err
	}

	return nil
}

func Services(mocks bool) error {
	// Clean up the build directory
	sh.Run("rm", "-rf", "build")

	if err := sh.Run("mkdir", "-p", "build"); err != nil {
		return errors.Wrap(err, "failed to create build directory")
	}

	for _, s := range services {
		servicePath := fmt.Sprintf("./services/%v/cmd/%v", s.nameWithVersion(), s.name)
		if exists, _ := dirExists(servicePath); !exists {
			log.Printf("Skipping %v, not found", servicePath)
			continue
		}

		log.Printf("Building service %v", servicePath)
		var err error
		if mocks {
			err = sh.Run("go", "build", "-o", "build", itestsTagsOpt, servicePath)
		} else {
			err = sh.Run("go", "build", "-o", "build", servicePath)
		}
		if err != nil && mg.ExitStatus(err) != 1 {
			return errors.Wrap(err, "failed to build service")
		}
	}

	return nil
}

// Restarts and rebuilds the services by just one command
// Use: mage restart
func Restart() error {
	if err := sh.Run("mage", "-v", "down"); err != nil {
		return errors.Wrap(err, "failed to down services")
	}

	if err := sh.Run("mage", "-v", "proto:buf"); err != nil {
		return errors.Wrap(err, "failed to generate protobuf")
	}

	if err := sh.Run("mage", "-v", "docker:buildImage"); err != nil {
		return errors.Wrap(err, "failed to build docker image")
	}

	// Uncomment this if you want to rebuild the lightning image
	// if err := sh.Run("mage", "-v", "docker:buildLightningImage"); err != nil {
	// 	return errors.Wrap(err, "failed to build docker image")
	// }

	if err := sh.Run("mage", "-v", "up"); err != nil {
		return errors.Wrap(err, "failed to up services")
	}

	return nil
}

func Tidy() error {
	return sh.Run("go", "mod", "tidy", "-compat=1.17")
}

func Vendor() error {
	return sh.Run("go", "mod", "vendor")
}

func IntegrationTests(runRegexp string) error {
	return sh.Run(
		"go",
		"test",
		"-v",
		"-p=1",
		"-count=1",
		itestsTagsOpt,
		fmt.Sprint("-run=", runRegexp),
		"./test/integration-tests/...",
	)
}

func IntegrationTestsForService(serviceName string) error {
	return sh.Run(
		"go",
		"test",
		"-v",
		"-p=1",
		"-count=1",
		itestsTagsOpt,
		fmt.Sprintf("./test/integration-tests/%s/...", serviceName),
	)
}

func UnitTests() error {
	//var docker Docker

	//err := docker.CopyLndFiles()
	//if err != nil {
	//	return err
	//}
	//
	//err = PrepareLndNode()
	//if err != nil {
	//	return err
	//}

	return sh.Run("go", "test", "-p=1", "-count=1", itestsTagsOpt, "./services/...", "./lib/...", "./test/unit_tests/...")
}

func E2ETests() error {
	return sh.Run("go", "test", "-p=1", "-count=1", "-tags=", "./test/e2e/...")
}

type Docker mg.Namespace

func (Docker) BuildImageNoMocks() error {
	return buildImage(false, false)
}

func (Docker) BuildImage() error {
	return buildImage(true, false)
}

func (Docker) BuildImageM1() error {
	return buildImage(true, true)
}

func buildImage(mocks, m1 bool) error {
	dockerfileName := "Dockerfile"
	if m1 {
		dockerfileName = "Dockerfile.m1"
	}
	return sh.Run(
		"docker",
		"build",
		"--build-arg", "CACHE_DATE=$(date +%Y-%m-%d:%H:%M:%S)",
		"--build-arg", fmt.Sprintf("MOCKS=%v", mocks),
		"-t",
		"fivi",
		"-f",
		dockerfileName,
		".",
	)
}

func (Docker) BuildLightningImage() error {
	return sh.Run("docker", "build", "-f", "Dockerfile.lightning", "-t", "fivi-lightning", ".")
}

func (Docker) ITests(runRegexp string) error {
	return sh.Run("docker", "exec", "build", "mage", "-v", "integrationTests", runRegexp)
}

func (Docker) ITestsForService(serviceName string) error {
	return sh.Run("docker", "exec", "build", "mage", "-v", "integrationTestsForService", serviceName)
}

// func (Docker) UnitTests() error {
// 	return sh.Run("docker", "exec", "build", "mage", "-v", "UnitTests")
// }

func (Docker) E2ETests() error {
	return sh.Run("docker", "exec", "build", "mage", "-v", "e2etests")
}

func (Docker) Migrate() error {
	return sh.Run("docker", "exec", "build", "mage", "-v", "db:migrate")
}

// Use: mage -v docker:protoBufBuild v1.28.0 v2.11.0
func (Docker) ProtoBufBuild(protoBufVersion string, grpcGatewayVersion string) error {
	fmt.Println("✅ Building protobuf compiler docker image")

	err := sh.Run("docker", "build",
		"--platform", "linux/amd64",
		"-f", "Dockerfile.proto",
		"-t", "protobuf-builder",
		"--build-arg", fmt.Sprintf("PROTOBUF_VERSION=%s", protoBufVersion),
		"--build-arg", fmt.Sprintf("GRPC_GATEWAY_VERSION=%s", grpcGatewayVersion),
		".")

	return err
}

//// Use: mage -v docker:copyLndFiles
//func (Docker) CopyLndFiles() error {
//	return framework.CopyLndFiles()
//}
//
//// Use: mage -v prepareLndNode
//func PrepareLndNode() error {
//	paramsToConnect, err := app_config.GetLNDParamsForUnit()
//	if err != nil {
//		return err
//	}
//
//	_, _, err = framework.Prepare(paramsToConnect)
//
//	return err
//}

type Proto mg.Namespace

// Generate the protobuf files for all services,
// Use: mage proto:buf.
func (Proto) Buf() error {
	// Clean up the gen directory
	if err := sh.Run("rm", "-rf", "gen"); err != nil {
		return err
	}

	// Generate the protobuf files for all registered services.
	for _, service := range services {
		path := fmt.Sprintf("proto/%v", service.nameWithVersion())
		// Skip if the service doesn't have a proto file.
		if !doesDirectoryHasFiles(path, ".proto") {
			fmt.Printf("-- Skipping %v, no protobuf files found\n", path)
			continue
		}

		// Generate the protobuf files for the service.
		if err := sh.Run("buf", "generate",
			"--path", fmt.Sprintf("proto/%v", service.nameWithVersion()),
		); err != nil {
			fmt.Printf("❌ Failed to generate protobuf for %v\n", service.name)
		} else {
			fmt.Printf("✅ Generated protobuf for %v\n", service.name)
		}

		protofiles, _ := getProtoFiles(path)
		for _, protofile := range protofiles {
			// Generate documentation for the service.
			err := sh.Run(
				"protoc",
				"-I.",
				"-Ithird_party/proto",
				"-I./proto",
				"--openapiv2_out=./docs",
				"--openapiv2_opt=logtostderr=true",
				fmt.Sprintf("proto/%v/%s", service.nameWithVersion(), protofile),
			)
			if err != nil && mg.ExitStatus(err) != 1 {
				fmt.Printf("❌ Failed to generate documentation for %v\n", service.name)
			} else {
				fmt.Printf("✅ Generated documentation for %v\n", service.name)
			}
		}
	}

	return nil
}

// Deprecated: Use mage proto:buf instead.
func (Proto) Generate() error {
	for _, s := range services {
		err := sh.Run(
			"protoc",
			"-I.",
			"-I./third_party/proto",
			"--go_out=./gen/go",
			"--go-grpc_out=require_unimplemented_servers=false:./gen/go",
			"--go_opt=paths=source_relative",
			"--go-grpc_opt=paths=source_relative",
			fmt.Sprintf("proto/%v/v1/service.proto", s.nameWithVersion()),
		)
		if err != nil && mg.ExitStatus(err) != 1 {
			return err
		}
	}

	for _, s := range services {
		err := sh.Run(
			"protoc",
			"-I.",
			"-Ithird_party/proto",
			"--grpc-gateway_out=./gen/go",
			"--grpc-gateway_opt=logtostderr=true",
			"--grpc-gateway_opt=paths=source_relative",
			"--grpc-gateway_opt=generate_unbound_methods=true",
			fmt.Sprintf("proto/%v/v1/service.proto", s.nameWithVersion()),
		)
		if err != nil && mg.ExitStatus(err) != 1 {
			return err
		}
	}

	for _, s := range services {
		err := sh.Run(
			"protoc",
			"-I.",
			"-Ithird_party/proto",
			"--openapiv2_out=./docs",
			"--openapiv2_opt=logtostderr=true",
			fmt.Sprintf("proto/%v/v1/service.proto", s.nameWithVersion()),
		)
		if err != nil && mg.ExitStatus(err) != 1 {
			return err
		}
	}

	return nil
}

type DB mg.Namespace

type MigrationConfig struct {
	DatabaseUrl     string `env:"DATABASE_URL,notEmpty" envDefault:"postgresql://postgres:123456@localhost:6432/db_name_db?sslmode=disable"`
	SystemDBName    string `env:"SYSTEM_DB_NAME" envDefault:"postgres"`
	MigrationsTable string `env:"DATABASE_MIGRATIONS_TABLE" envDefault:"migrations"`
	MigrationsPath  string `env:"MIGRATIONS_PATH" envDefault:"./services/%s/repository/sql/migrations"`
}

// Migrate the database to the latest version.
// Use: mage db:migrate.
func (DB) Migrate(serviceName string) error {
	var (
		n            int
		migrationNum int
		err          error
	)

	cfg := new(MigrationConfig)
	if err := env.Parse(cfg); err != nil {
		return errors.Wrap(err, "can't get MigrationConfig from envvars")
	}

	if serviceName != "all" {
		service, err := findServiceByName(serviceName)
		if err != nil {
			return err
		}
		log.Println(service.name)

		if n, err = runMigrations(cfg, service); err != nil {
			return err
		}

		fmt.Printf("Applied %d migrations!\n", n)

		return nil
	}

	for _, s := range services {
		n, err = runMigrations(cfg, s)
		dirNotExistsErr := "dir does not exist"
		if err != nil && !strings.Contains(err.Error(), dirNotExistsErr) {
			return err
		}

		migrationNum += n
	}

	fmt.Printf("Applied %d migrations!\n", migrationNum)

	return nil
}

func findServiceByName(serviceName string) (*service, error) {
	for _, s := range services {
		if serviceName == s.name {
			return s, nil
		}
	}

	return nil, errors.Errorf("can't find service by %v name", serviceName)
}

func runMigrations(cfg *MigrationConfig, s *service) (int, error) {
	dirPath := fmt.Sprintf(cfg.MigrationsPath, s.nameWithVersion())
	if exist, _ := dirExists(dirPath); !exist {
		return 0, errors.New(fmt.Sprintf("%v dir does not exist", dirPath))
	}

	regexpStr := regexp.MustCompile(`/([a-z_]+)_db`)
	databaseUrl := regexpStr.ReplaceAllString(cfg.DatabaseUrl, fmt.Sprintf("/%s_db", s.name))
	log.Println("databaseUrl", databaseUrl)
	dbUrl, err := url.Parse(databaseUrl)
	if err != nil {
		return 0, errors.Wrap(err, "can't parse DB conn string")
	}
	dbUser := dbUrl.User.Username()

	defaultDatabaseUrl := regexpStr.ReplaceAllString(cfg.DatabaseUrl, fmt.Sprintf("/%s", cfg.SystemDBName))
	db, err := sql.Open("postgres", defaultDatabaseUrl)
	if err != nil {
		return 0, errors.Wrapf(err, "can't open connection to database %s_db", s.name)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		return 0, errors.Wrapf(err, "can't ping database %s_db", s.name)
	}

	{
		sqlQuery := fmt.Sprintf("CREATE DATABASE %v_db OWNER %v;", s.name, dbUser)
		dbAlreadyExistsErr := fmt.Sprintf("database \"%v_db\" already exists", s.name)
		_, err = db.Exec(sqlQuery)
		if err != nil && !strings.Contains(err.Error(), dbAlreadyExistsErr) {
			return 0, errors.Wrapf(err, "can't create %v_db database", s.name)
		}
	}

	db, err = sql.Open("postgres", databaseUrl)
	if err != nil {
		return 0, errors.Wrapf(err, "can't open connection to database %s_db", s.name)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		return 0, errors.Wrapf(err, "can't ping database %s_db", s.name)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: dirPath,
	}

	m := migrate.MigrationSet{
		TableName: cfg.MigrationsTable,
	}

	n, err := m.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return 0, errors.Wrap(err, "can't exec migrations")
	}

	return n, nil
}

// check if the directory exists
// returns true if it does, false if it doesn't
// or an error if something went wrong
func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// Check if the directory not empty
// returns true if it is, false if it is empty or an error if something went wrong
func doesDirectoryHasFiles(path string, ext string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}

	if len(files) == 0 {
		return false
	}

	// check if the directory contains any files with given extension
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%s is a directory\n", file.Name())
			continue
		}
		if strings.HasSuffix(file.Name(), ext) {
			return true
		}
	}

	return false
}

// Get proto files from the directory
// returns a slice of proto files or an error if something went wrong
func getProtoFiles(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't read directory")
	}

	var protoFiles []string
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%s is a directory\n", file.Name())
			continue
		}
		if strings.HasSuffix(file.Name(), ".proto") {
			protoFiles = append(protoFiles, file.Name())
		}
	}

	return protoFiles, nil
}

var (
	DBConnString = "postgresql://postgres:123456@localhost:6432/db_name_db?sslmode=disable"
	ImagesDir    = "/Users/anna/Documents/Education/fivi/assets"

	likesCfg = likes_cmd.Config{
		GrpcServerPort: 9094,
		DBConnString:   "postgresql://postgres:123456@localhost:6432/likes_db?sslmode=disable",
	}

	profileCfg = profile_cmd.Config{
		GrpcServerPort:      9090,
		DBConnString:        "postgresql://postgres:123456@localhost:6432/profile_db?sslmode=disable",
		ImagesDir:           ImagesDir,
		PostsServerAddr:     "0.0.0.0:9092",
		FollowersServerAddr: "0.0.0.0:9093",
	}

	commentsCfg = comments_cmd.Config{
		GrpcServerPort:    9091,
		ProfileServerAddr: "0.0.0.0:9090",
		DBConnString:      "postgresql://postgres:123456@localhost:6432/comments_db?sslmode=disable",
	}

	followersCfg = followers_cmd.Config{
		GrpcServerPort:    9093,
		ProfileServerAddr: "0.0.0.0:9090",
		DBConnString:      "postgresql://postgres:123456@localhost:6432/followers_db?sslmode=disable",
	}

	postsCfg = posts_cmd.Config{
		GrpcServerPort:      9092,
		DBConnString:        "postgresql://postgres:123456@localhost:6432/posts_db?sslmode=disable",
		ImagesDir:           ImagesDir,
		ProfileServerAddr:   "0.0.0.0:9090",
		CommentsServerAddr:  "0.0.0.0:9091",
		LikesServerAddr:     "0.0.0.0:9094",
		FollowersServerAddr: "0.0.0.0:9093",
	}

	gatewayCfg = gateway_cmd.Config{
		HttpPublicServerPort:    8088,
		ProfilesGrpcServerAddr:  "0.0.0.0:9090",
		CommentsGrpcServerAddr:  "0.0.0.0:9091",
		LikesGrpcServerAddr:     "0.0.0.0:9094",
		PostsGrpcServerAddr:     "0.0.0.0:9092",
		FollowersGrpcServerAddr: "0.0.0.0:9093",
		StripPrefix:             "/gateway",
		JwtSigningKey:           "Uv38ByGCZU8WP18PmmIdcpVmx00QA3xNe7sEB9Hixkk=",
		JwtTTL:                  time.Hour * 24,
		StaticDir:               "/Users/anna/Documents/Education/fivi/web",
		ImagesDir:               ImagesDir,
	}
)

func Run() error {
	var (
		ctx = context.Background()
	)

	profile := profile_cmd.WithConfig(&profileCfg)
	comments := comments_cmd.WithConfig(&commentsCfg)
	follower := followers_cmd.WithConfig(&followersCfg)
	likes := likes_cmd.WithConfig(&likesCfg)
	posts := posts_cmd.WithConfig(&postsCfg)
	gateway, err := gateway_cmd.WithConfig(&gatewayCfg)
	if err != nil {
		return err
	}

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		profile.Run(ctx)
		return nil
	})
	group.Go(func() error {
		comments.Run(ctx)
		return nil
	})
	group.Go(func() error {
		follower.Run(ctx)
		return nil
	})
	group.Go(func() error {
		likes.Run(ctx)
		return nil
	})
	group.Go(func() error {
		posts.Run(ctx)
		return nil
	})
	group.Go(func() error {
		gateway.Run(ctx)
		return nil
	})

	err = group.Wait()

	profile.Shutdown()
	comments.Shutdown()
	follower.Shutdown()
	likes.Shutdown()
	posts.Shutdown()
	gateway.Shutdown()

	return err
}
