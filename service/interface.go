package service

import (
	"context"
	"michaelyusak/go-desent.git/entity"
)

type Book interface {
	CreateBook(ctx context.Context, book *entity.Book) error
	GetAllBook(ctx context.Context, filter entity.GetBookFilter) ([]*entity.Book, error)
	GetBookById(ctx context.Context, id string) (*entity.Book, error)
	UpdateBook(ctx context.Context, book *entity.Book) (*entity.Book, error)
	DeleteBook(ctx context.Context, id string) (*entity.Book, error)
}

type Auth interface {
	Validate(ctx context.Context, user entity.User) (string, error)
	ValidateToken(ctx context.Context, token string) bool
}
