package controllers

import (
	"github.com/fzzp/ebook/internal/services"
	"github.com/go-playground/form/v4"
)

type Repository struct {
	// 用户模块
	users services.UserService

	// form表单解码器
	formDec *form.Decoder
}

func NewRepository(userSvc services.UserService) *Repository {
	return &Repository{
		formDec: form.NewDecoder(),
		users:   userSvc,
	}
}
