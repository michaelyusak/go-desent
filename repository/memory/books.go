package memory

import (
	"context"
	"fmt"
	"michaelyusak/go-desent.git/entity"
	"sync"
)

type books struct {
	storage    []*entity.Book
	hotStorage map[string]*entity.Book

	mu sync.Mutex
}

func NewBooks() *books {
	sample := entity.Book{
		Id: "521caa08-97e4-49a1-a175-174db38d5528",
	}

	return &books{
		storage: []*entity.Book{&sample},
		hotStorage: map[string]*entity.Book{
			"521caa08-97e4-49a1-a175-174db38d5528": &sample,
		},
	}
}

func (r *books) CreateOne(ctx context.Context, book *entity.Book) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage = append(r.storage, book)
	r.hotStorage[book.Id] = book
}

func (r *books) GetAll(ctx context.Context, filter entity.GetBookFilter) []*entity.Book {
	r.mu.Lock()
	booksCopy := make([]*entity.Book, len(r.storage))
	copy(booksCopy, r.storage)
	r.mu.Unlock()

	paged := filter.Page != 0 && filter.Limit != 0
	skipped := 0

	res := []*entity.Book{}

	for _, book := range booksCopy {
		if filter.Author != "" && filter.Author != book.Author {
			continue
		}

		if paged && skipped < filter.Offset {
			skipped++
			continue
		}

		res = append(res, book)

		if paged && len(res) == filter.Limit {
			break
		}
	}

	return res
}

func (r *books) GetById(ctx context.Context, id string) *entity.Book {
	r.mu.Lock()
	defer r.mu.Unlock()

	book, ok := r.hotStorage[id]
	if !ok {
		return nil
	}

	return book
}

func (r *books) UpdateOne(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	existing, ok := r.hotStorage[book.Id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}

	existing.Mu.Lock()
	defer existing.Mu.Unlock()

	if book.Author != "" {
		existing.Author = book.Author
	}

	if book.Title != "" {
		existing.Title = book.Title
	}

	if book.Year != 0 {
		existing.Year = book.Year
	}

	return existing, nil
}

func (r *books) DeleteOne(ctx context.Context, id string) *entity.Book {
	r.mu.Lock()
	defer r.mu.Unlock()

	new := []*entity.Book{}

	var deleted *entity.Book

	for i, book := range r.storage {
		if book.Id == id {
			new = append(new, r.storage[:i]...)
			new = append(new, r.storage[i+1:]...)

			deleted = book

			break
		}
	}

	r.storage = new
	delete(r.hotStorage, id)

	return deleted
}
