package model

type Districts struct {
	ID     uint64 `json:"id"`
	CityID uint64 `json:"city_id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status int8   `json:"status"`

	City *Cities `json:"city,omitempty"`
}
