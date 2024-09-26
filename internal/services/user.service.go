package services

import (
	"context"

	dbrepo "github.com/fzzp/ebook/dbrepo/sqlc"
	"github.com/fzzp/ebook/internal/formdto"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, input formdto.CreateUserRequest) error
}

var _ UserService = (*userService)(nil)

type userService struct {
	store dbrepo.Store
}

func NewUserService(store dbrepo.Store) *userService {
	srv := &userService{
		store: store,
	}

	return srv
}

func (s *userService) CreateUser(ctx context.Context, input formdto.CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return err
	}
	arg := dbrepo.CreateUserParams{
		Email:        input.Email,
		Username:     input.Username,
		PasswordHash: string(hashedPassword),
	}
	res, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
