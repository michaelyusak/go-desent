package service

import (
	"context"
	"michaelyusak/go-desent.git/entity"
)

type Book interface {
	CreateBook(ctx context.Context, book *entity.Book) error
	GetAllBook(ctx context.Context) ([]*entity.Book, error)
	GetBookById(ctx context.Context, id string) (*entity.Book, error)
	UpdateBook(ctx context.Context, book *entity.Book) (*entity.Book, error)
	DeleteBook(ctx context.Context, id string) (*entity.Book, error)
}
