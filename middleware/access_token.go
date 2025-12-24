package middleware

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/mataharibiz/sange/v2/middleware/auth"
)

func AccessTokenAuthentication() iris.Handler {
	return auth.New(auth.Config{
		AuthURL: os.Getenv("AUTH_API_URL") + "/auth/authenticate",
	})
}
