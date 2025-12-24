package model

type Provinces struct {
	ID        uint64 `json:"id"`
	CountryID uint64 `json:"country_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Rank      uint   `json:"rank"`
	Status    uint8  `json:"status"`

	Country *Countries `json:"country,omitempty"`
}
