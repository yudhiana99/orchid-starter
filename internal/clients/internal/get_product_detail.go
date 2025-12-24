package internalClient

import (
	"context"
	"fmt"
	"os"
	"time"

	"orchid-starter/internal/common"

	modelCommon "orchid-starter/internal/common/model"

	"github.com/mataharibiz/sange/v2"
)

func (ic *InternalClient) GetProductDetail(ctx context.Context, productID uint64) (result modelCommon.ProductDetailResponse, err error) {

	templateUrl := "{{app_api_url}}/products/product/{{product_id}}"

	url, errRender := common.Render(templateUrl, map[string]any{
		"app_api_url": os.Getenv("APP_API_URL"),
		"product_id":  productID,
	})
	if errRender != nil {
		return result, sange.NewError(sange.OperationFailed, errRender, "failed to render template", "internalClient", "GetProductDetail")
	}

	token := common.GetAppTokenFromContext(ctx)
	appOrigin := common.GetAppOriginFromContext(ctx)

	restyClient := GetRestyClient().
		SetDebug(ic.Debug).
		SetTimeout(time.Minute).
		SetHeader("Accept", "application/json").
		SetAuthScheme("Bearer").
		SetAuthToken(token).
		SetHeader("X-App-Origin", appOrigin)

	if token == "" {
		restyClient.SetHeader("X-App-Origin", ic.AppOriginSystem)
		restyClient.SetAuthToken(ic.TokenSystem)
	}

	response, err := restyClient.R().SetContext(ctx).SetResult(&result).Get(url)
	if err != nil {
		return result, sange.NewError(sange.CallServicesFailed, err, "failed to get product detail", "internalClient", "GetProductDetail")
	}

	if response.IsError() {
		return result, sange.NewError(sange.CallServicesFailed, fmt.Errorf("failed to get product detail: %s", response.String()), "failed to get product detail", "internalClient", "GetProductDetail")
	}
	return
}
