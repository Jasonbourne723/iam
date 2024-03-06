package model

type User struct {
	UserId int64
	Name   string
	Avatar string
	Remark string
}

type App struct {
	AppId  int64
	Name   string
	Avatar string
	Secret string
}

type Role struct {
	RoleId int64
	Name   string
}

type UserRole struct {
	Id     int64
	RoleId int64
	UserId int64
}

type RoleApp struct {
	Id     int64
	RoleId int64
	AppId  int64
}
