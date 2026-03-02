package memory

import (
	"context"
	"michaelyusak/go-desent.git/entity"
)

type users struct {
	storage []entity.User
}

func NewUsers() *users {
	return &users{
		storage: []entity.User{
			{
				Username: "admin",
				Password: "password",
			},
		},
	}
}

func (r *users) GetByUsername(ctx context.Context, username string) *entity.User {
	for _, user := range r.storage {
		if user.Username == username {
			return &user
		}
	}

	return nil
}
