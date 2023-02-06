package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRepo interface {
	Save(context.Context, *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (c *UserUseCase) Register(ctx context.Context, u *User) (*User, error) {
	c.log.WithContext(ctx).Infof("Registering user: %v", u)
	return c.repo.Save(ctx, u)
}
