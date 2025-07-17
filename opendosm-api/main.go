package main

import (
	"context"
	"embed"
	"fmt"
	"os"

	"github.com/cschleiden/go-workflows/backend"
	mysqlWorkflow "github.com/cschleiden/go-workflows/backend/mysql"
	"github.com/cschleiden/go-workflows/worker"
	"github.com/muazwzxv/opendosm-api/database/repository"
	"github.com/muazwzxv/opendosm-api/service/item"
	"github.com/muazwzxv/opendosm-api/service/premise"

	"github.com/muazwzxv/opendosm-api/http/route"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/fsutil"

	// Import the appropriate GORM dialect for the database you're using.
	// _ "goyave.dev/goyave/v5/database/dialect/mysql"
	_ "goyave.dev/goyave/v5/database/dialect/postgres"
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

	// repository setup
	itemLookupRepository := repository.NewItemLookup(server.DB(), server.Logger)
	premiseLookupRepository := repository.NewPremiseLookup(server.DB(), server.Logger)

	// service setup
	itemService := item.NewItemService(
		itemLookupRepository,
		server.Logger.With("service", "item_service"),
	)

	premiseService := premise.NewPremiseService(
		premiseLookupRepository,
		server.Logger.With("service", "premise_service"),
	)

	// register service to instance container
	server.RegisterService(itemService)
	server.RegisterService(premiseService)
}

func runWorker(ctx context.Context, mb backend.Backend) {
	w := worker.New(mb, nil)

	//w.RegisterActivity(Activity1)
	//w.RegisterActivity(Activity2)

	if err := w.Start(ctx); err != nil {
		panic("could not start worker")
	}
}
