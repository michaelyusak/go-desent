package repository

import (
	"context"
	"michaelyusak/go-desent.git/entity"
)

type Books interface {
	CreateOne(ctx context.Context, book *entity.Book)
	GetAll(ctx context.Context) []*entity.Book
	GetById(ctx context.Context, id string) *entity.Book
	UpdateOne(ctx context.Context, book *entity.Book) (*entity.Book, error)
	DeleteOne(ctx context.Context, id string) *entity.Book
}
