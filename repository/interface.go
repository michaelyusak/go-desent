package repository

import (
	"context"
	"michaelyusak/go-desent.git/entity"
)

type Books interface {
	CreateOne(ctx context.Context, book *entity.Book)
	GetAll(ctx context.Context, filter entity.GetBookFilter) []*entity.Book
	GetById(ctx context.Context, id string) *entity.Book
	UpdateOne(ctx context.Context, book *entity.Book) (*entity.Book, error)
	DeleteOne(ctx context.Context, id string) *entity.Book
}

type Users interface {
	GetByUsername(ctx context.Context, username string) *entity.User
}

type Tokens interface {
	IsExist(ctx context.Context, token string) bool
	InsertToken(ctx context.Context, token string)
}
