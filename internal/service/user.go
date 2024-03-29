package service

import (
	"github.com/jasonbourne723/iam/internal/dto"
	"github.com/jasonbourne723/iam/internal/global"
	"github.com/jasonbourne723/iam/internal/model"
)

var UserSrv = UserService{}

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

func (u *UserService) Delete(userId int64) {
	global.App.DB.Delete(&model.User{}, userId)
}

func (u *UserService) Update(updateUserDto *dto.UpdateUserDto) error {

	var user model.User
	if err := global.App.DB.First(&user, updateUserDto.Id).Error; err != nil {
		return err
	}

	user.Avatar = updateUserDto.Avatar
	user.Name = updateUserDto.Name
	user.Remark = updateUserDto.Remark

	return global.App.DB.Save(&user).Error
}

func (u *UserService) Get(userId int64) (*dto.UserDto, error) {
	var user model.User
	if err := global.App.DB.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &dto.UserDto{
		Id:     user.UserId,
		Name:   user.Name,
		Avatar: user.Avatar,
		Remark: user.Remark,
	}, nil
}
