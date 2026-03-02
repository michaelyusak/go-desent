package memory

import (
	"context"
	"sync"
)

type tokens struct {
	storage map[string]bool

	mu sync.Mutex
}

func NewTokens() *tokens {
	return &tokens{
		storage: map[string]bool{
			"94d4c7ba-8957-46e0-ac23-7f5a7c8a464d": true,
		},
	}
}

func (r *tokens) IsExist(ctx context.Context, token string) bool {
	_, ok := r.storage[token]
	return ok
}

func (r *tokens) InsertToken(ctx context.Context, token string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage[token] = true
}
