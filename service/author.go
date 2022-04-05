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
func (s *catalogService) GetAuthors(ctx context.Context, req *catalog.ListReq) (*catalog.AuthorListResp, error) {
	resp, count, err := s.storage.Author().SelectAuthors(req.Page, req.Limit)
	if err != nil {
		return &catalog.AuthorListResp{}, err
	}
	return &catalog.AuthorListResp{
		Authors: resp,
		Count:   count,
	}, nil

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
func (s *catalogService) UpdateAuthor(ctx context.Context, req *catalog.Author) (*catalog.Author, error) {
	author, err := s.storage.Author().UpdateAuthor(*req)
	if err != nil {
		s.logger.Error("failed to update author", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to update author")
	}
	return &author, nil
}

//
func (s *catalogService) DeleteAuthor(ctx context.Context, req *catalog.ByIdReq) (*catalog.Empty, error) {
	if err := s.storage.Author().DeleteAuthor(req.Id); err != nil {
		s.logger.Error("failed to delete author", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete author")
	}
	return &catalog.Empty{}, nil
}
