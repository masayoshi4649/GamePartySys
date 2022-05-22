package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"

	"github.com/masayoshi4649/GamePartySys/src/common"
	"github.com/masayoshi4649/GamePartySys/src/routing"
)

func main() {
	conf := common.Config
	var r *gin.Engine = gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions(conf.App.Session, store))
	r.Use(routing.LoginCheck())

	routing.Initialize(r)

	routing.Common(r)
	routing.Root(r)
	routing.Auth(r)
	routing.Booking(r)

	r.Run(":" + conf.App.Port)
}
