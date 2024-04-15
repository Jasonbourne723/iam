package dto

type Result struct {
	Code    int32
	Message string
}

type PageResult[T any] struct {
	TotalPage int64
	TotalRows int64
	PageIndex int64
	PageSize  int64
	Data      []T
}
