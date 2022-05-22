package routeauth

import (
	"github.com/gin-gonic/gin"
	"github.com/masayoshi4649/GamePartySys/src/common"
)

// Login ... https://example.com/auth/login/
func Login(ctx *gin.Context) {
	conf := common.Config

	labelmap := make(map[string]interface{})
	labelmap["discordURL"] = conf.Discord.LoginURL

	ctx.HTML(200, "auth/login.html", labelmap)
}
