package middleware

import (
	"orchid-starter/internal/common"

	"github.com/kataras/iris/v12"
	"github.com/mataharibiz/sange/v2"
	sangeValidation "github.com/mataharibiz/sange/v2/middleware/validate"
	"github.com/mataharibiz/ward/logging"
)

func ValidateMultipleCompany(irisCtx iris.Context) *sange.Error {
	logging.NewLogger().Debug("starting executed authenticate multiple company", "module", "middleware", "function", "ValidateMultipleCompany")
	ctx := common.SetRequestContext(irisCtx.Request().Context(), irisCtx)
	companyId, err := common.ConvertUInt64ToInt64(common.GetCompanyIDFromContext(ctx))
	if err != nil {
		return sange.SetError(sange.IncorrectParam, err, "company id is not valid")
	}

	if validation := sangeValidation.MultipleCompanyAuthenticate(irisCtx, companyId); validation != nil {
		return validation
	}
	return nil
}
