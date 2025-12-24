package model

type Company struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	StoreName   string `json:"store_name"`
	IsSellerPKP uint8  `json:"is_seller_pkp"`

	CompanyDocumentVerification *CompanyDocumentVerifications `json:"company_document_verification"`
	CompanyStores               []CompanyStore                `json:"company_stores"`
}
