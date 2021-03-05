package service

import (
	"context"
	"github.com/aeramu/clean-architecture/entity"
	"github.com/aeramu/clean-architecture/service/api"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
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
	if err := req.Validate(); err != nil {
		return nil, err
	}
	err := s.adapter.BookRepository.InsertBook(ctx, entity.Book{
		ID:         xid.New().String(),
		Title:      req.Title,
		Author:     req.Author,
		CoverImage: req.CoverImage,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"payload": req,
		}).Errorln("[CreateBook] Error when create book to db")
		return nil, err
	}

	return &api.CreateBookRes{
		Message: "Success",
	}, nil
}

func (s *service) GetBook(ctx context.Context, req api.GetBookReq) (*api.GetBookRes, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	book, err := s.adapter.BookRepository.FindBookByID(ctx, req.BookID)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
			"payload": req,
		}).Errorln("[GetBook] Error when get book from db")
		return nil, err
	}

	return &api.GetBookRes{
		Book: book,
	}, nil
}
