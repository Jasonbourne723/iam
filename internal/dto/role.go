package dto

type RoleDto struct {
	Id   int64
	Name string
	Apps []AppDto
}

type AddRoleDto struct {
	Name string
	Apps []AppDto
}

type UpdateRoleDto struct {
	Id   int64
	Name string
	Apps []AppDto
}
