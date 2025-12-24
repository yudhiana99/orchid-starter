package handler

import (
	"orchid-starter/internal/bootstrap"
	api "orchid-starter/modules/default/delivery/api/rest"

	"github.com/kataras/iris/v12"
)

func GQLRoutes(app *iris.Application, container *bootstrap.Container) {
	container.Log.Info("Initialize default handler...")

	// Get DI from container instead of creating new instance
	di := container.GetDI()

	app.PartyFunc("/", func(defaultParty iris.Party) {
		api.NewDefaultHandler(defaultParty, di)
	})
}
