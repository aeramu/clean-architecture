package service

import (
	"context"
	"github.com/aeramu/clean-architecture/entity"
)

type Adapter struct {
	BookRepository BookRepository
}

type BookRepository interface {
	FindBookByID(ctx context.Context, bookID string) (*entity.Book, error)
	InsertBook(ctx context.Context, book entity.Book) error
}
