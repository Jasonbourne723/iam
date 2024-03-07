package service

import (
	"github.com/jasonbourne723/iam/internal/dto"
	"github.com/jasonbourne723/iam/internal/global"
	"github.com/jasonbourne723/iam/internal/model"
)

type UserService struct {
}

func (u *UserService) Add(userDto *dto.AddUserDto) error {

	user := model.User{
		Name:   userDto.Name,
		Avatar: userDto.Avatar,
		Remark: userDto.Remark,
		UserId: 1,
	}

	return global.App.DB.Create(&user).Error
}
