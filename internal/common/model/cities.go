package model

type Cities struct {
	ID         uint64 `json:"id"`
	ProvinceID uint64 `json:"province_id"`
	Name       string `json:"name"`
	Status     uint8  `json:"status"`

	Province *Provinces `json:"province,omitempty"`
}
