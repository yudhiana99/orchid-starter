package model

type CompanyDocumentVerifications struct {
	ID        int64 `json:"id"`
	CompanyID int64 `json:"company_id"`
	Status    uint8 `json:"status"`
}
