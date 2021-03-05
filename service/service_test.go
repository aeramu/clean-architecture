package service

import (
	"context"
	"errors"
	"github.com/aeramu/clean-architecture/entity"
	"github.com/stretchr/testify/assert"
	"github.com/aeramu/clean-architecture/mocks"
	"github.com/aeramu/clean-architecture/service/api"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	adapter Adapter
	mockBookRepository *mocks.BookRepository
)

func initTest()  {
	mockBookRepository = new(mocks.BookRepository)
	adapter = Adapter {
		BookRepository: mockBookRepository,
	}
}

func Test_service_CreateBook(t *testing.T)  {
	var (
		ctx = context.Background()
	)
	type args struct {
		ctx context.Context
		req api.CreateBookReq
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    *api.CreateBookRes
		wantErr bool
	}{
		{
			name:    "should error",
			prepare: func(){

			},
			args:    args{
				ctx: ctx,
				req: api.CreateBookReq{},
			},
			want:    nil,
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest()
			if tt.prepare != nil {
				tt.prepare()
			}
			s := &service{
				adapter: adapter,
			}
			got, err := s.CreateBook(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_service_GetBook(t *testing.T)  {
	var (
		ctx = context.Background()
		bookID = "some-id"
		book = &entity.Book{
			ID:         bookID,
			Title:      "some title",
			Author:     "some author",
			CoverImage: "some_image.jpg",
		}
	)
	type args struct {
		ctx context.Context
		req api.GetBookReq
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    *api.GetBookRes
		wantErr bool
	}{
		{
			name:    "invalid argument empty book id",
			prepare: func(){

			},
			args:    args{
				ctx: ctx,
				req: api.GetBookReq{
					BookID: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "error database when get book from db",
			prepare: func(){
				mockBookRepository.On("FindBookByID", mock.Anything, bookID).
					Return(nil, errors.New("some error"))
			},
			args:    args{
				ctx: ctx,
				req: api.GetBookReq{
					BookID: bookID,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "success",
			prepare: func(){
				mockBookRepository.On("FindBookByID", mock.Anything, bookID).
					Return(book, nil)
			},
			args:    args{
				ctx: ctx,
				req: api.GetBookReq{
					BookID: bookID,
				},
			},
			want:    &api.GetBookRes{Book: book},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest()
			if tt.prepare != nil {
				tt.prepare()
			}
			s := &service{
				adapter: adapter,
			}
			got, err := s.GetBook(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
