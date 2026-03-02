package service

import (
	"context"
	"fmt"
	"michaelyusak/go-desent.git/apperror"
	"michaelyusak/go-desent.git/entity"
	"michaelyusak/go-desent.git/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type book struct {
	booksRepo repository.Books
}

func NewBook(
	booksRepo repository.Books,
) *book {
	return &book{
		booksRepo: booksRepo,
	}
}

func (s *book) CreateBook(ctx context.Context, book *entity.Book) error {
	book.Id = uuid.NewString()

	s.booksRepo.CreateOne(ctx, book)

	return nil
}

func (s *book) GetAllBook(ctx context.Context, filter entity.GetBookFilter) ([]*entity.Book, error) {
	filter.Offset = (filter.Page - 1) * filter.Limit

	books := s.booksRepo.GetAll(ctx, filter)

	logrus.WithField("books", fmt.Sprintf("%+v", books)).WithField("count", len(books)).WithField("filter", fmt.Sprintf("%+v", filter)).Info("[service][book][GetAllBook] got books")

	return books, nil
}

func (s *book) GetBookById(ctx context.Context, id string) (*entity.Book, error) {
	book := s.booksRepo.GetById(ctx, id)
	if book == nil {
		logrus.Warn("[service][book][GetBookById] book not found")

		return nil, &apperror.AppError{
			Code:    http.StatusNotFound,
			Message: "not found",
		}
	}

	return book, nil
}

func (s *book) UpdateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	updated, err := s.booksRepo.UpdateOne(ctx, book)
	if err != nil {
		if err.Error() == "not found" {
			return nil, &apperror.AppError{
				Code:    http.StatusNotFound,
				Message: "not found",
			}
		}

		return nil, err
	}

	return updated, nil
}

func (s *book) DeleteBook(ctx context.Context, id string) (*entity.Book, error) {
	return s.booksRepo.DeleteOne(ctx, id), nil
}
