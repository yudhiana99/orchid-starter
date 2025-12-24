package model

type Subdistricts struct {
	ID         uint64 `json:"id"`
	DistrictID uint64 `json:"district_id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Zipcode    string `json:"zipcode"`
	Status     int8   `json:"status"`

	District *Districts `json:"district,omitempty"`
}
