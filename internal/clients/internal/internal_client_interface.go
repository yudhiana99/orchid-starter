package internalClient

import (
	"context"

	modelCommon "orchid-starter/internal/common/model"
)

type ParamClientGetStorage struct {
	URLParams []string
}

type InternalClientInterface interface {
	GetStorageById(ctx context.Context, id uint64, input ParamClientGetStorage) (modelCommon.StorageData, error)
	GetCompanyGQLDetail(ctx context.Context, companyID uint64, selected string) (result modelCommon.CompanyDetailGQLResponse, err error)
	GetProductDetail(ctx context.Context, productID uint64) (result modelCommon.ProductDetailResponse, err error)
	GetSkuDetail(ctx context.Context, skuID uint64, urlParams []string) (result modelCommon.SkuDetailResponse, err error)
	AuthenticateToken(ctx context.Context) (err error)
	AuthenticateThirdPartyToken(ctx context.Context) (err error)
}
