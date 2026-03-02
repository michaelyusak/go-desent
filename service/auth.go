package service

import (
	"context"
	"michaelyusak/go-desent.git/apperror"
	"michaelyusak/go-desent.git/entity"
	"michaelyusak/go-desent.git/repository"
	"net/http"

	"github.com/google/uuid"
)

type auth struct {
	usersRepo  repository.Users
	tokensRepo repository.Tokens
}

func NewAuth(
	usersRepo repository.Users,
	tokensRepo repository.Tokens,
) *auth {
	return &auth{
		usersRepo:  usersRepo,
		tokensRepo: tokensRepo,
	}
}

func (s *auth) Validate(ctx context.Context, user entity.User) (string, error) {
	registered := s.usersRepo.GetByUsername(ctx, user.Username)
	if registered == nil || registered.Password != user.Password {
		return "", &apperror.AppError{
			Code:    http.StatusUnauthorized,
			Message: "invalid credentials",
		}
	}

	token := uuid.NewString()

	s.tokensRepo.InsertToken(ctx, token)

	return token, nil
}

func (s *auth) ValidateToken(ctx context.Context, token string) bool {
	return s.tokensRepo.IsExist(ctx, token)
}
