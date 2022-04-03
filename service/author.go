package service

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/perfectogo/catalog-service/genproto/catalog"
	"github.com/perfectogo/catalog-service/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//
func (s *catalogService) CreateAuthor(ctx context.Context, req *catalog.Author) (*catalog.Author, error) {
	uuId, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	req.AuthorId = uuId.String()

	resp, err := s.storage.Author().InsertAuthor(*req)
	if err != nil {
		s.logger.Error("failed to create author", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to create author")
	}

	return &resp, nil
}

//
func (s *catalogService) GetAuthor(ctx context.Context, req *catalog.ByIdReq) (*catalog.Author, error) {
	author, err := s.storage.Author().SelectAuthor(req.Id)
	if err != nil {
		s.logger.Error("failed to get author", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to get author")
	}
	return &author, nil
}

//
func (s *catalogService) GetAuthors(ctx context.Context, req *catalog.ListReq) (*catalog.AuthorListResp, error) {

	return &catalog.AuthorListResp{}, nil

}

//
func (s *catalogService) UpdateAuthor(ctx context.Context, req *catalog.Author) (*catalog.Author, error) {
	return &catalog.Author{}, nil
}

//
func (s *catalogService) DeleteAuthor(ctx context.Context, req *catalog.ByIdReq) (*catalog.Empty, error) {
	return &catalog.Empty{}, nil
}
