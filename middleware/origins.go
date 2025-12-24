package middleware

import (
	"fmt"

	"orchid-starter/constants"

	"github.com/kataras/iris/v12"
	"github.com/mataharibiz/sange/v2"
)

type OriginID uint8

const (
	AppSystemID OriginID = 1 // system
	AppAdminID  OriginID = 2 // admin alun-alun (tarrasque)
	AppWorkerID OriginID = 3 // worker
	AppBuyerID  OriginID = 4 // vladmir, vembrace
	AppSellerID OriginID = 5 // voodoo, trident
)

type Origins struct {
	OriginsID []OriginID
}

func AppAllowedOrigin() *Origins {
	return &Origins{}
}

func (base *Origins) SetOriginID(id ...OriginID) *Origins {
	base.OriginsID = id
	return base
}

func (base *Origins) Filter(ctx iris.Context) *sange.Error {
	appOrigin := ctx.GetHeader(constants.HeaderAppOrigin)
	isAllowed := make(map[string]bool, len(base.OriginsID))
	for _, originID := range base.OriginsID {
		switch originID {
		case AppBuyerID:
			isAllowed[constants.AppVladmir] = true
			isAllowed[constants.AppVembrace] = true
		case AppSellerID:
			isAllowed[constants.AppVoodoo] = true
			isAllowed[constants.AppTrident] = true
		case AppAdminID:
			isAllowed[constants.AppTarrasque] = true
		case AppSystemID:
			isAllowed[constants.AppSystem] = true
		case AppWorkerID:
			isAllowed[constants.AppWorker] = true
		}
	}

	if !isAllowed[appOrigin] {
		msgError := fmt.Errorf("action not allowed for origin %s", appOrigin)
		return sange.SetError(sange.Forbidden, msgError)
	}

	return nil
}
