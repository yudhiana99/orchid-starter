package model

import "time"

type ProductDetailResponse struct {
	Status  int  `json:"status"`
	Success bool `json:"success"`
	Data    struct {
		ID       uint64 `json:"id"`
		Category struct {
			ID        uint64 `json:"id"`
			PathIds   string `json:"path_ids"`
			PathLabel string `json:"path_label"`
			Paths     []struct {
				ID   uint64 `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
			} `json:"paths"`
		} `json:"category"`
		Brand struct {
			ID   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"brand"`
		Company struct {
			ID          uint64 `json:"id"`
			Name        string `json:"name"`
			StoreName   string `json:"store_name"`
			StoreSlug   string `json:"store_slug"`
			IsSellerPkp uint8  `json:"is_seller_pkp"`
			Location    struct {
				Country  string `json:"country"`
				Province string `json:"province"`
				City     string `json:"city"`
				District string `json:"district"`
			} `json:"location"`
			ImageStorageID uint64 `json:"image_storage_id"`
			Status         int    `json:"status"`
			IsStoreActive  int    `json:"is_store_active"`
			CompanyStores  []struct {
				ID            uint64 `json:"id"`
				Name          string `json:"name"`
				Slug          string `json:"slug"`
				CompanyID     uint64 `json:"company_id"`
				LogoStorageID int    `json:"logo_storage_id"`
				StorageData   struct {
					ID               uint64 `json:"id"`
					Type             string `json:"type"`
					Path             string `json:"path"`
					Filename         string `json:"filename"`
					OriginalFilename string `json:"original_filename"`
					Mime             string `json:"mime"`
					URL              string `json:"url"`
				} `json:"storage_data"`
				StoreAddress []struct {
					ID             uint64     `json:"id"`
					CompanyStoreID uint64     `json:"company_store_id"`
					CountryID      uint64     `json:"country_id"`
					ProvinceID     uint64     `json:"province_id"`
					CityID         uint64     `json:"city_id"`
					DistrictID     uint64     `json:"district_id"`
					Zipcode        string     `json:"zipcode"`
					Address        string     `json:"address"`
					CreatedAt      time.Time  `json:"created_at"`
					UpdatedAt      time.Time  `json:"updated_at"`
					DeletedAt      *time.Time `json:"deleted_at"`
					Province       struct {
						ID        uint64 `json:"id"`
						CountryID uint64 `json:"country_id"`
						Code      string `json:"code"`
						Name      string `json:"name"`
						Slug      string `json:"slug"`
						Rank      int    `json:"rank"`
						Status    int    `json:"status"`
					} `json:"province"`
					City struct {
						ID         uint64 `json:"id"`
						ProvinceID uint64 `json:"province_id"`
						Name       string `json:"name"`
						Status     int    `json:"status"`
					} `json:"city"`
					District struct {
						ID     uint64 `json:"id"`
						CityID uint64 `json:"city_id"`
						Code   string `json:"code"`
						Name   string `json:"name"`
						Status int    `json:"status"`
					} `json:"district"`
				} `json:"store_address"`
				StoreRatings any       `json:"store_ratings"`
				CreatedAt    time.Time `json:"created_at"`
				UpdatedAt    time.Time `json:"updated_at"`
				DeletedAt    any       `json:"deleted_at"`
			} `json:"company_stores"`
			CompanyDocumentVerification struct {
				Status uint8 `json:"status"`
			} `json:"company_document_verification"`
			SellerBusinessScale int `json:"seller_business_scale"`
			Institution         struct {
				ID   uint64 `json:"id"`
				Code string `json:"code"`
			} `json:"institution"`
			Documents []struct {
				ID           uint64 `json:"id"`
				CompanyID    uint64 `json:"company_id"`
				DocumentType struct {
					ID    uint64 `json:"id"`
					Slug  string `json:"slug"`
					Title string `json:"title"`
				} `json:"document_type"`
				Value       string `json:"value"`
				StorageID   uint64 `json:"storage_id"`
				StorageData struct {
					ID               uint64 `json:"id"`
					Type             string `json:"type"`
					Path             string `json:"path"`
					Filename         string `json:"filename"`
					OriginalFilename string `json:"original_filename"`
					Mime             string `json:"mime"`
					URL              string `json:"url"`
				} `json:"storage_data"`
				ApprovalStatus int       `json:"approval_status"`
				ApprovalNote   string    `json:"approval_note"`
				CreatedAt      time.Time `json:"created_at"`
				UpdatedAt      time.Time `json:"updated_at"`
			} `json:"documents"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"company"`
		Title            string  `json:"title"`
		Slug             string  `json:"slug"`
		ShortDescription string  `json:"short_description"`
		Description      string  `json:"description"`
		ProductWeight    float32 `json:"product_weight"`
		ProductHeight    float32 `json:"product_height"`
		ProductWidth     float32 `json:"product_width"`
		ProductLength    float32 `json:"product_length"`
		ShippingWeight   float32 `json:"shipping_weight"`
		ShippingHeight   float32 `json:"shipping_height"`
		ShippingWidth    float32 `json:"shipping_width"`
		ShippingLength   float32 `json:"shipping_length"`
		OriginData       struct {
			ID                 uint8   `json:"id"`
			Name               string  `json:"name"`
			TkdnPercentage     float64 `json:"tkdn_percentage"`
			TkdnCertificateURL string  `json:"tkdn_certificate_url"`
			BmpPercentage      float64 `json:"bmp_percentage"`
			BmpCertificateURL  string  `json:"bmp_certificate_url"`
		} `json:"origin_data"`
		Unit struct {
			ID     uint64 `json:"id"`
			NameEn string `json:"name_en"`
			NameID string `json:"name_id"`
		} `json:"unit"`
		Msds struct {
			ID   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"msds"`
		ProductType int `json:"product_type"`
		Warranty    struct {
			WarrantyTypeID       int    `json:"warranty_type_id"`
			WarrantyType         string `json:"warranty_type"`
			WarrantyPeriod       int    `json:"warranty_period"`
			WarrantyPeriodUnitID int    `json:"warranty_period_unit_id"`
			WarrantyPeriodUnit   string `json:"warranty_period_unit"`
			Description          string `json:"description"`
			WarrantyFullText     string `json:"warranty_full_text"`
		} `json:"warranty"`
		Etalase struct {
			ID   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"etalase"`
		IsContract    uint8  `json:"is_contract"`
		Status        uint8  `json:"status"`
		Reason        string `json:"reason"`
		Variants      any    `json:"variants"`
		ProductImages []struct {
			ID          uint64 `json:"id"`
			SkuID       uint64 `json:"sku_id"`
			ProductID   uint64 `json:"product_id"`
			StorageID   uint64 `json:"storage_id"`
			FileName    string `json:"file_name"`
			ImageType   int    `json:"image_type"`
			ImgURL      string `json:"img_url"`
			StorageData any    `json:"storage_data"`
		} `json:"product_images"`
		PriceRules []struct {
			ID        uint64 `json:"id"`
			ProductID uint64 `json:"product_id"`
			Qty       int    `json:"qty"`
		} `json:"price_rules"`
		ProductAttributes []struct {
			ID             uint64 `json:"id"`
			AttributeID    uint64 `json:"attribute_id"`
			ProductID      uint64 `json:"product_id"`
			AttributeName  string `json:"attribute_name"`
			AttributeValue string `json:"attribute_value"`
		} `json:"product_attributes"`
		Specifications any `json:"specifications"`
		Warehouses     []struct {
			ID   uint64 `json:"id"`
			Name string `json:"name"`
		} `json:"warehouses"`
		Skus []struct {
			ID    uint64 `json:"id"`
			SkuNo string `json:"sku_no"`
		}
		ContractData  any        `json:"contract_data"`
		UpdatedBy     int        `json:"updated_by"`
		AdminName     string     `json:"admin_name"`
		CheckedBy     uint64     `json:"checked_by"`
		CheckedByName string     `json:"checked_by_name"`
		CheckedAt     *time.Time `json:"checked_at"`
		IsChecked     uint8      `json:"is_checked"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedAt     time.Time  `json:"updated_at"`
		DeletedAt     *time.Time `json:"deleted_at"`
	} `json:"data"`
}
