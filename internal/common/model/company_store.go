package model

import "time"

type CompanyStore struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	CompanyID     int64  `json:"company_id"`
	LogoStorageID int64  `json:"logo_storage_id"`

	StoreAddress []CompanyStoreAddress `json:"store_address"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
