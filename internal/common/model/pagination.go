package model

type PaginationResult struct {
	Page        int   `json:"page"`
	TotalPage   int   `json:"total_page"`
	TotalItems  int64 `json:"total_items"`
	PerPage     int   `json:"per_page"`
	HasNext     bool  `json:"has_next"`
	HasPrevious bool  `json:"has_previous"`
}
