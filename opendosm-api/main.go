package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/cschleiden/go-workflows/backend"
	mysqlWorkflow "github.com/cschleiden/go-workflows/backend/mysql"
	"github.com/cschleiden/go-workflows/worker"
	"os"

	"github.com/muazwzxv/opendosm-api/http/route"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/fsutil"
	// Import the appropriate GORM dialect for the database you're using.
	// _ "goyave.dev/goyave/v5/database/dialect/mysql"
	// _ "goyave.dev/goyave/v5/database/dialect/postgres"
	// _ "goyave.dev/goyave/v5/database/dialect/sqlite"
	// _ "goyave.dev/goyave/v5/database/dialect/mssql"
	// _ "goyave.dev/goyave/v5/database/dialect/clickhouse"
)

//go:embed resources
var resources embed.FS

func main() {
	resources := fsutil.NewEmbed(resources)
	langFS, err := resources.Sub("resources/lang")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
		os.Exit(1)
	}

	opts := goyave.Options{
		LangFS: langFS,
	}

	server, err := goyave.New(opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.(*errors.Error).String())
		os.Exit(1)
	}

	server.Logger.Info("Registering hooks")
	server.RegisterSignalHook()

	server.RegisterStartupHook(func(s *goyave.Server) {
		server.Logger.Info("Server is listening", "host", s.Host())
	})

	server.RegisterShutdownHook(func(s *goyave.Server) {
		s.Logger.Info("Server is shutting down")
	})

	registerServices(server)

	server.Logger.Info("Registering routes")
	server.RegisterRoutes(route.Register)

	if err := server.Start(); err != nil {
		server.Logger.Error(err)
		os.Exit(2)
	}

	ctx := context.Background()

	b := mysqlWorkflow.NewMysqlBackend("", 0, "", "", "",
		mysqlWorkflow.WithApplyMigrations(true))
	go runWorker(ctx, b)
}

func registerServices(server *goyave.Server) {
	server.Logger.Info("Registering services")

	// Services represent the Domain/Business layer.
	// This is where the core logic and value of your application resides.
	// This function is where you will register your services in the server's
	// service container to make them accessible to dependents.
	// https://goyave.dev/basics/services.html#service-container

	// TODO register services
}

func runWorker(ctx context.Context, mb backend.Backend) {
	w := worker.New(mb, nil)

	//w.RegisterActivity(Activity1)
	//w.RegisterActivity(Activity2)

	if err := w.Start(ctx); err != nil {
		panic("could not start worker")
	}
}
