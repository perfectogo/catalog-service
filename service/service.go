package service

import (
	"github.com/perfectogo/catalog-service/pkg/logger"
	"github.com/perfectogo/catalog-service/storage"
)

type catalogService struct {
	storage storage.InterfaceStorage
	logger  logger.Logger
}

func NewCatalogService(storage storage.InterfaceStorage, log logger.Logger) *catalogService {
	return &catalogService{
		storage: storage,
		logger:  log,
	}
}
