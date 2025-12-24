package middleware

import (
	"orchid-starter/constants"
	"orchid-starter/http"
	"orchid-starter/internal/common"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/versioning"
)

func SetAPIVersion(ctx iris.Context) {
	versioning.SetVersion(ctx, func() string {
		if apiVersion := ctx.GetHeader(constants.HeadersVersionKey); apiVersion != "" {
			return apiVersion
		}
		return LatestAPIVersion()
	}())
	versioning.NotFoundHandler = http.NotFoundApiVersion
	ctx.Next()
}

func LatestAPIVersion() string {
	return common.GetEnvWithDefault("LATEST_API_VERSION", "2.0.0")
}
