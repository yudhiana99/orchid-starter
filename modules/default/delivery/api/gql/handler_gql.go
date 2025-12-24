package gqlHandler

import (
	"orchid-starter/http"
	"orchid-starter/internal/bootstrap"
	v2 "orchid-starter/modules/default/delivery/api/rest/v2"
	"orchid-starter/modules/default/repository"
	"orchid-starter/modules/default/usecase"

	"github.com/kataras/iris/v12"
	promHttp "github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewDefaultGQLHandler(app iris.Party, di *bootstrap.DirectInjection) {

	defaultRepository := repository.NewDefaultRepository(di.GetMySQL(), di.GetElasticsearch())

	// Get the comprehensive client for all API operations
	client := di.GetClient()

	// Initialize usecase with client access
	defaultUseCase := usecase.NewDefaultUsecase(di.GetMySQL(), defaultRepository, client)
	defaultV2 := v2.NewDefaultHandler(defaultUseCase)

	app.Get("/metrics", iris.FromStd(promHttp.Handler()))
	app.Get("/", defaultV2.Welcome)
	app.Get("/health-check", http.HealthCheckHandler)
	app.OnErrorCode(iris.StatusNotFound, http.NotFoundHandler)
}
