package dto

type PaginationRequestDto struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}
