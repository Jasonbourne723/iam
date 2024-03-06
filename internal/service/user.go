package service

import "github.com/jasonbourne723/iam/internal/dto"

type UserService struct {
}

func (u *UserService) Add(user *dto.AddUserDto) (int64, error) {
	return 1, nil
}
