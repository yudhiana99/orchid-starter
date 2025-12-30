package internalClient

import (
	"context"
	"fmt"
	"orchid-starter/internal/common"
	"os"

	"github.com/mataharibiz/sange/v2"
)

func (ic *InternalClient) AuthenticateToken(ctx context.Context) (err error) {
	templateUrl := "{{app_api_url}}/auth/authenticate"

	url, errRender := common.Render(templateUrl, map[string]any{
		"app_api_url": os.Getenv("APP_API_URL"),
	})
	if errRender != nil {
		return errRender
	}

	restyClient := GetRestyClient().
		SetDebug(ic.Debug).
		SetHeader("Content-Type", "application/json").
		SetAuthScheme("Bearer").
		SetAuthToken(common.GetAppTokenFromContext(ctx)).
		SetHeader("X-App-Origin", common.GetAppOriginFromContext(ctx))

	response, err := restyClient.R().SetContext(ctx).Get(url)
	if err != nil {
		return sange.NewError(sange.CallServicesFailed, err, "failed to authenticate token", "internalClient", "AuthenticateToken")
	}

	if response.IsError() {
		return sange.NewError(sange.CallServicesFailed, fmt.Errorf("failed to authenticate token: %s", response.String()), "failed to authenticate token", "internalClient", "AuthenticateToken")
	}

	return nil
}

func (ic *InternalClient) AuthenticateThirdPartyToken(ctx context.Context) (err error) {
	templateUrl := "{{app_api_url}}/auth/authenticate?thirdparty=true"

	url, errRender := common.Render(templateUrl, map[string]any{
		"app_api_url": os.Getenv("APP_API_URL"),
	})
	if errRender != nil {
		return errRender
	}

	restyClient := GetRestyClient().
		SetDebug(ic.Debug).
		SetHeader("Content-Type", "application/json").
		SetAuthScheme("Bearer").
		SetAuthToken(common.GetAppTokenFromContext(ctx)).
		SetHeader("X-Client-Id", common.GetClientIDFromContext(ctx))

	response, err := restyClient.R().SetContext(ctx).Get(url)
	if err != nil {
		return sange.NewError(sange.CallServicesFailed, err, "failed to authenticate token 3rd party", "internalClient", "AuthenticateThirdPartyToken")
	}

	if response.IsError() {
		return sange.NewError(sange.CallServicesFailed, fmt.Errorf("failed to authenticate token 3rd party: %s", response.String()), "failed to authenticate token", "internalClient", "AuthenticateThirdPartyToken")
	}

	return nil
}
