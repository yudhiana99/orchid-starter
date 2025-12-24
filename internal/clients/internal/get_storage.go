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

func (ic *InternalClient) GetStorageById(ctx context.Context, id uint64, input ParamClientGetStorage) (modelCommon.StorageData, error) {
	templateUrl := "{{app_api_url}}/storage/storages/{{storage_id}}"

	url, errRender := common.Render(templateUrl, map[string]any{
		"app_api_url": os.Getenv("APP_API_URL"),
		"storage_id":  id,
	})
	if errRender != nil {
		return modelCommon.StorageData{}, errRender
	}

	var urlParams string
	if len(input.URLParams) > 0 {
		urlParams = strings.Join(input.URLParams, "&")
	}
	if urlParams != "" {
		url = url + "?" + urlParams
	}

	restyClient := GetRestyClient().
		SetDebug(ic.Debug).
		SetTimeout(time.Minute).
		SetHeader("Accept", "application/json").
		SetAuthScheme("Bearer").
		SetAuthToken(common.GetAppTokenFromContext(ctx)).
		SetHeader("X-App-Origin", common.GetAppOriginFromContext(ctx))

	if common.GetAppTokenFromContext(ctx) == "" {
		restyClient.SetHeader("X-App-Origin", ic.AppOriginSystem)
		restyClient.SetAuthToken(ic.TokenSystem)
	}

	var result struct {
		ID               uint64 `json:"id"` // used to unmarshal storage data from storage service
		FileName         string `json:"file_name,omitempty"`
		Type             string `json:"type"`
		Path             string `json:"path"`
		Filename         string `json:"filename"`
		Mime             string `json:"mime"`
		OriginalFilename string `json:"original_filename,omitempty"`
	}

	response, err := restyClient.R().SetContext(ctx).SetResult(&result).Get(url)
	if err != nil {
		return modelCommon.StorageData{}, sange.NewError(sange.CallServicesFailed, err, "failed to get storage by id", "internalClient", "GetStorageById")
	}

	if response.IsError() {
		return modelCommon.StorageData{}, sange.NewError(sange.CallServicesFailed, fmt.Errorf("failed to get storage by id: %s", response.String()), "failed to get storage by id", "internalClient", "GetStorageById")
	}
	return modelCommon.StorageData{
		StorageID:        result.ID,
		Type:             result.Type,
		Path:             result.Path,
		Filename:         result.Filename,
		Mime:             result.Mime,
		OriginalFilename: result.OriginalFilename,
	}, nil
}
