package model

import (
	"apiserver/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) Create() error {
	logger.Info("create user ", u)
	return DB.Self.Create(&u).Error
}

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
