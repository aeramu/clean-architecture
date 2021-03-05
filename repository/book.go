package repository

import (
	"context"
	"github.com/aeramu/clean-architecture/entity"
	"github.com/aeramu/mongolib"
)

type repository struct {
	db *mongolib.Database
}

func (r *repository) FindBookByID(ctx context.Context, bookID string) (*entity.Book, error) {
	var book Book
	err := r.db.Coll("book").Query().
		Equal("id", bookID).
		FindOne(ctx).Consume(&book)
	if err != nil {
		return nil, err
	}
	return book.Entity(), nil
}

func (r *repository) InsertBook(ctx context.Context, book entity.Book) error {
	panic("implement me")
}

