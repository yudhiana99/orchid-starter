package internalClient

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"orchid-starter/internal/common"

	modelCommon "orchid-starter/internal/common/model"

	"github.com/mataharibiz/sange/v2"
)

func (ic *InternalClient) GetSkuDetail(ctx context.Context, skuID uint64, urlParams []string) (result modelCommon.SkuDetailResponse, err error) {

	templateUrl := "{{app_api_url}}/products/sku/{{sku_id}}"

	url, errRender := common.Render(templateUrl, map[string]any{
		"app_api_url": os.Getenv("APP_API_URL"),
		"sku_id":      skuID,
	})
	if errRender != nil {
		return result, sange.NewError(sange.OperationFailed, errRender, "failed to render template", "internalClient", "GetSkuDetail")
	}

	var urlParam string
	if len(urlParams) > 0 {
		urlParam = strings.Join(urlParams, "&")
	}

	if urlParam != "" {
		url = url + "?" + urlParam
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
		return result, sange.NewError(sange.CallServicesFailed, err, "failed to get SKU detail", "internalClient", "GetSkuDetail")
	}

	if response.IsError() {
		return result, sange.NewError(sange.CallServicesFailed, fmt.Errorf("failed to get SKU detail: %s", response.String()), "failed to get SKU detail", "internalClient", "GetSkuDetail")
	}
	return
}
