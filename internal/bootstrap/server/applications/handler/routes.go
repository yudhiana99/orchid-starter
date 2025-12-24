package handler

import (
	"orchid-starter/internal/bootstrap"

	"github.com/kataras/iris/v12"
)

// RouteSetup represents a function that sets up routes
type RouteSetup func(app *iris.Application, container *bootstrap.Container)

// SetupAllRoutes configures all application routes in an organized manner
func SetupAllRoutes(app *iris.Application, container *bootstrap.Container) {
	container.Log.Info("Setting up all application routes...")

	// Define all route setup functions
	routeSetups := []RouteSetup{
		SetupDefaultRoutes, // Root routes
		// Add more route setups here as your application grows
	}

	// Execute all route setups
	for _, setupFunc := range routeSetups {
		setupFunc(app, container)
	}

	container.Log.Info("All application routes configured successfully")
}
