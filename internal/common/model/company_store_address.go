package model

import "time"

type CompanyStoreAddress struct {
	ID             int64  `json:"id"`
	CompanyStoreID int64  `json:"company_store_id"`
	CountryID      int    `json:"country_id"`
	ProvinceID     int    `json:"province_id"`
	CityID         int    `json:"city_id"`
	DistrictID     int    `json:"district_id"`
	SubdistrictID  int    `json:"subdistrict_id"`
	Zipcode        string `json:"zipcode"`
	Address        string `json:"address"`
	AddressText    string `json:"address_text"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Country     *Countries    `json:"country,omitempty"`
	Province    *Provinces    `json:"province,omitempty"`
	City        *Cities       `json:"city,omitempty"`
	District    *Districts    `json:"district,omitempty"`
	Subdistrict *Subdistricts `json:"subdistrict,omitempty"`
}
