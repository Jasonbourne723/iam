package dto

type UserDto struct {
	Id     int64
	Name   string
	Avatar string
	Remark string
	Roles  []RoleDto
}

type AddUserDto struct {
	Name   string
	Avatar string
	Remark string
	Roles  []RoleDto
}

type UpdateUserDto struct {
	Id     int64
	Name   string
	Avatar string
	Remark string
	Roles  []RoleDto
}


