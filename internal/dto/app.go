package dto

type AppDto struct {
	AppId  int64
	Name   string
	Avatar string
	Secret string
}

type AddAppDto struct {
	Name   string
	Avatar string
}

type UpdateAppDto struct {
	AppId  int64
	Name   string
	Avatar string
}
