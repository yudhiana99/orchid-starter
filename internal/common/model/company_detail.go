package model

import (
	"time"

	"github.com/mataharibiz/sange/v2/gql"
)

type CompanyDetailGQLResponse struct {
	Errors []*gql.ErrorGQL `json:"errors,omitempty"`
	Data   struct {
		CompanyDetail struct {
			Items struct {
				ID                     uint64    `json:"id"`
				Name                   string    `json:"name"`
				StoreName              string    `json:"storeName"`
				DescriptionEn          string    `json:"descriptionEn"`
				DescriptionID          string    `json:"descriptionId"`
				IsSellerPkp            uint8     `json:"isSellerPkp"`
				Type                   uint8     `json:"type"`
				CreatedAt              time.Time `json:"createdAt"`
				UpdatedAt              time.Time `json:"updatedAt"`
				UpdateProductTaxStatus uint8     `json:"updateProductTaxStatus"`
				ImageStorageId         uint64    `json:"imageStorageId"`
				SellerBusinessScale    *uint8    `json:"seller_business_scale"`
				StorageData            struct {
					FileName         string `json:"fileName"`
					FileType         string `json:"fileType"`
					Mime             string `json:"mime"`
					OriginalFilename string `json:"originalFilename"`
					Path             string `json:"path"`
				} `json:"storageData"`
				CompanyStores []struct {
					ID   uint64 `json:"id"`
					Name string `json:"name"`
					Slug string `json:"slug"`
				} `json:"companyStores"`
				CompanyDocumentVerifications struct {
					Status uint8 `json:"status"`
				} `json:"companyDocumentVerifications"`
				Institutions struct {
					Code string `json:"code"`
				} `json:"institutions"`
				Province struct {
					ID   uint64 `json:"id"`
					Name string `json:"name"`
				} `json:"province"`
				City struct {
					ID   uint64 `json:"id"`
					Name string `json:"name"`
				} `json:"city"`
				District struct {
					ID   uint64 `json:"id"`
					Name string `json:"name"`
				} `json:"district"`
				SubDistrict struct {
					ID   uint64 `json:"id"`
					Name string `json:"name"`
				} `json:"subDistrict"`
			} `json:"items"`
		} `json:"companyDetail"`
	} `json:"data"`
}
