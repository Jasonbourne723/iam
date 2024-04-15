package service

import (
	"math"

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

func (u *UserService) List(pageIndex int, pageSize int) (*dto.PageResult[dto.UserDto], error) {

	users := make([]model.User, 0, pageSize)
	if err := global.App.DB.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&users).Error; err != nil {
		return nil, err
	}
	var totalRows int64
	if err := global.App.DB.Model(&model.User{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	result := dto.PageResult[dto.UserDto]{
		PageIndex: int64(pageIndex),
		PageSize:  int64(pageSize),
		TotalRows: totalRows,
		TotalPage: int64(math.Ceil(float64(totalRows) / float64(pageSize))),
		Data:      mapToUserDto(users),
	}

	return &result, nil
}

func mapToUserDto(users []model.User) []dto.UserDto {
	userDtos := make([]dto.UserDto, len(users))
	for i, v := range users {
		userDtos[i] = dto.UserDto{
			Id:     v.UserId,
			Avatar: v.Avatar,
			Name:   v.Name,
			Remark: v.Remark,
		}
	}
	return userDtos
}
