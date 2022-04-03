package main

import (
	"net"

	"github.com/perfectogo/catalog-service/config"
	"github.com/perfectogo/catalog-service/genproto/catalog"
	"github.com/perfectogo/catalog-service/pkg/db"
	"github.com/perfectogo/catalog-service/pkg/logger"
	"github.com/perfectogo/catalog-service/service"
	"github.com/perfectogo/catalog-service/storage"
	"google.golang.org/grpc"
)

func main() {
	// Loading configures.
	cfg := config.Load()

	// Loading logger
	log := logger.New(cfg.LogLevel, "catalog-service")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	// Creating DB connection.
	sqlxDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	//defer sqlxDB.Close()

	//Getting storage instance
	storage := storage.NewStoragePg(sqlxDB)

	// Getting an instance of listener
	catalogService := service.NewCatalogService(storage, log)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	newService := grpc.NewServer()
	catalog.RegisterCatalogServiceServer(newService, catalogService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := newService.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
