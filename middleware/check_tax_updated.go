package middleware

import (
	"orchid-starter/constants"
	"orchid-starter/internal/clients"
	"orchid-starter/internal/common"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/mataharibiz/sange/v2"
)

func CheckTaxUpdated(irisCtx iris.Context) {
	client := clients.NewClient()
	internalClient := client.InternalClient
	ctx := common.SetRequestContext(irisCtx.Request().Context(), irisCtx)
	appOrigin := common.GetAppOriginFromContext(ctx)
	if constants.IsSeller(appOrigin) {
		companyId := common.GetCompanyIDFromContext(ctx)
		selected := `items { updateProductTaxStatus type }`

		result, errGet := internalClient.GetCompanyGQLDetail(ctx, companyId, selected)
		if errGet != nil {
			sange.NewResponse(irisCtx, iris.StatusInternalServerError, errGet.Error())
			return
		}

		whiteListCompanyType := map[string]bool{
			strconv.FormatInt(constants.GetConstant("TYPE_VENDOR"), 10): true,
			strconv.FormatInt(constants.GetConstant("TYPE_BOTH"), 10):   true,
		}

		whiteListProductTaxStatus := map[string]iris.Map{
			strconv.FormatInt(constants.GetConstant("PRODUCT_TAX_NEED_UPDATE"), 10): {
				"error_code":    "PRODUCT_TAX_NEED_UPDATE",
				"error_message": "need update product tax",
			},
			strconv.FormatInt(constants.GetConstant("PRODUCT_TAX_INPROGRESS"), 10): {
				"error_code":    "UPDATE_PRODUCT_TAX_INPROGRESS",
				"error_message": "update product tax still in progress",
			},
		}

		if whiteListCompanyType[strconv.Itoa(int(result.Data.CompanyDetail.Items.Type))] {
			if message, ok := whiteListProductTaxStatus[strconv.Itoa(int(result.Data.CompanyDetail.Items.UpdateProductTaxStatus))]; ok {
				sange.NewResponse(irisCtx, iris.StatusBadRequest, message)
				return
			}
		}
	}
	irisCtx.Next()
}
