package modelPorts

import "time"

type CompanyDocumentVerification struct {
	Status uint8 `json:"status"`
}

type Institution struct {
	Code string `json:"code"`
}

type CatalogRelation struct {
	Name string `json:"name"`
}

type Province struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type District struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type Subdistrict struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type StoreSuggest struct {
	Input []string `json:"input"`
}

type CompanyData struct {
	ID                          uint64                      `json:"id"`
	Name                        string                      `json:"name"`
	StoreName                   string                      `json:"store_name"`
	StoreSlug                   *string                     `json:"store_slug"`
	DescriptionEn               string                      `json:"description_en"`
	Description                 string                      `json:"description"`
	IsSellerPkp                 uint8                       `json:"is_seller_pkp"`
	Type                        string                      `json:"type"`
	Image                       ImageData                   `json:"image"`
	CompanyDocumentVerification CompanyDocumentVerification `json:"company_document_verification"`
	Institution                 Institution                 `json:"institution"`
	CatalogRelation             CatalogRelation             `json:"catalog_relation"`
	Province                    Province                    `json:"province"`
	City                        City                        `json:"city"`
	District                    District                    `json:"district"`
	Subdistrict                 Subdistrict                 `json:"subdistrict"`
	StoreSuggest                StoreSuggest                `json:"store_suggest"`
	CreatedAt                   time.Time                   `json:"created_at"`
	UpdatedAt                   time.Time                   `json:"updated_at"`
	Timestamp                   time.Time                   `json:"timestamp"`
}

type CompanyLocationData struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Province    Province    `json:"province"`
	City        City        `json:"city"`
	District    District    `json:"district"`
	Subdistrict Subdistrict `json:"subdistrict"`
}

type ImageData struct {
	StorageID        uint64 `json:"storage_id"`
	Type             string `json:"type"`
	Path             string `json:"path"`
	Filename         string `json:"filename"`
	Mime             string `json:"mime"`
	OriginalFilename string `json:"original_filename"`
}
