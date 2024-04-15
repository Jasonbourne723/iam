package dto

type UserDto struct {
	Id     int64
	Name   string
	Avatar string
	Remark string
}

type AddUserDto struct {
	Name   string
	Avatar string
	Remark string
}

type UpdateUserDto struct {
	Id     int64
	Name   string
	Avatar string
	Remark string
}
