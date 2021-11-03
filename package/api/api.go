package api

import (
	"askUs/v1/package/advice"
	admdb "askUs/v1/package/advice/database/psql"
	adviceR "askUs/v1/package/advice/router"
	"askUs/v1/package/asset"
	assetR "askUs/v1/package/asset/router"
	"askUs/v1/package/asset/store"
	"askUs/v1/package/auth"
	amdb "askUs/v1/package/auth/databases/psql"
	authR "askUs/v1/package/auth/router"
	"askUs/v1/package/user"
	umdb "askUs/v1/package/user/databases/psql"
	userR "askUs/v1/package/user/router"
	jAuth "askUs/v1/util/auth"
	"askUs/v1/util/config"
	mid "askUs/v1/util/middleware"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type api struct {
	Client      *gorm.DB
	Version     string
	MiddleWares []echo.MiddlewareFunc
	Jwt         *jAuth.Auth
	Config      *config.Env
}

func Init(c *gorm.DB, jwt *jAuth.Auth, env *config.Env, m ...echo.MiddlewareFunc) *api {
	return &api{
		Client:      c,
		Version:     os.Getenv("VERSION"),
		MiddleWares: m,
		Jwt:         jwt,
		Config:      env,
	}
}
func (ap *api) Route(e *echo.Echo) {
	e.Use(mid.ConnectionMDB(ap.Client), mid.Logger(), mid.ClientCheck())

	v1 := e.Group("/api/" + ap.Version)

	// v1.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusAccepted, "Works well\n")
	// })
	userService := user.Init(&umdb.UserDb{})
	adviceService := advice.Init(&admdb.AdviceData{})
	authService := auth.Init(amdb.AuthDb{}, ap.Jwt, userService, ap.Config)
	assetService := asset.Init(store.Init())
	// ideaService := idea.Init(ideamdb.IdeaDB{})
	authR.Route(authService, v1, mid.ConnectionMDB(ap.Client))
	userR.Route(v1, userService, ap.MiddleWares...)
	adviceR.Route(v1, adviceService, ap.MiddleWares...)
	assetR.Route(v1, assetService, ap.MiddleWares...)
}
