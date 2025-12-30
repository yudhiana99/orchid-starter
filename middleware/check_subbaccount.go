package middleware

import (
	"fmt"
	"orchid-starter/constants"
	"orchid-starter/internal/common"

	"github.com/kataras/iris/v12"
	"github.com/mataharibiz/sange/v2"
	mbizUtil "github.com/mataharibiz/sange/v2/utils"
	"github.com/mataharibiz/ward/logging"
)

type CheckSubAccountInput struct {
	CompanyID uint64
	UserID    uint64
	AppOrigin string
}

func CheckSubAccount(irisCtx iris.Context) *sange.Error {
	ctx := common.SetRequestContext(irisCtx.Request().Context(), irisCtx)
	var (
		companyID = common.GetCompanyIDFromContext(ctx)
		userID    = common.GetUserIDFromContext(ctx)
		appOrigin = common.GetAppOriginFromContext(ctx)
	)

	checkInput := &CheckSubAccountInput{
		CompanyID: companyID,
		UserID:    userID,
		AppOrigin: appOrigin,
	}

	if exist := checkInput.CheckCacheAccount(); !exist {
		errMsg := fmt.Errorf("company [%d] doesn't have subaccount", companyID)
		return sange.SetError(sange.Forbidden, errMsg, "company subaccount")
	}

	return nil
}

func (ci *CheckSubAccountInput) CheckCacheAccount() bool {
	rdsUtil, errConn := mbizUtil.NewRedisUtil(constants.RedisTypeStatic)
	if errConn != nil {
		logging.NewLogger().Error("failed to connect redis static", "error", errConn)
		return false
	}

	redisKey := fmt.Sprintf("SUBACCOUNT:COMPANYID:%v", ci.CompanyID)
	value, errHget := rdsUtil.HGet(redisKey, "subaccount")
	if errHget != nil || value == "" {
		logging.NewLogger().Error("subaccount on redis is empty", "error", errHget)
		return false
	}

	return true
}
