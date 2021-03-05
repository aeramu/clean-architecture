package service

import (
	"context"
	"github.com/aeramu/clean-architecture/service/api"
)

type Service interface {
	CreateBook(ctx context.Context, req api.CreateBookReq) (*api.CreateBookRes, error)
	GetBook(ctx context.Context, req api.GetBookReq) (*api.GetBookRes, error)
}

func NewService(adapter Adapter) Service {
	return &service {
		adapter: adapter,
	}
}

type service struct {
	adapter Adapter
}

func (s *service) CreateBook(ctx context.Context, req api.CreateBookReq) (*api.CreateBookRes, error) {
	panic("implement me")
}

func (s *service) GetBook(ctx context.Context, req api.GetBookReq) (*api.GetBookRes, error) {
	panic("implement me")
}
