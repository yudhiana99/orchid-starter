package applications

import (
	"context"
	"sync"
	"time"

	"orchid-starter/internal/common"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/mataharibiz/ward/logging"

	sentryIris "github.com/getsentry/sentry-go/iris"
)

type irisAppUtil struct {
	app *iris.Application
}

var irisAppInstance *irisAppUtil
var onceIrisApp sync.Once

// GetIrisApplication get iris Application instance
func GetIrisApplication() *iris.Application {
	onceIrisApp.Do(func() {
		log := logging.NewLogger()
		log.Info("Initialize iris application instance...")

		app := iris.New()

		idleConnsClosed := make(chan struct{})
		iris.RegisterOnInterrupt(func() {
			log.Info("⚠️ Shutting down server...")

			// Default 10 seconds timeout for graceful shutdown
			timeout := common.GetIntEnv("GRACEFUL_SHUTDOWN_TIMEOUT", 10)
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
			defer cancel()

			// close all hosts
			if err := app.Shutdown(ctx); err != nil {
				log.Error("Server shutdown error", "error", err)
			} else {
				log.Info("✅ Graceful shutdown completed")
			}

			close(idleConnsClosed)
		})

		// recover from any http-relative panics
		app.Use(recover.New())

		// Log everything to terminal
		app.Use(logger.New())

		app.Use(sentryIris.New(sentryIris.Options{
			Repanic: true,
		}))

		irisAppInstance = &irisAppUtil{
			app: app,
		}
	})

	return irisAppInstance.app
}
