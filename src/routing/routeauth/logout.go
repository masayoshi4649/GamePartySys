package routeauth

import (
	"github.com/gin-gonic/gin"
	"github.com/masayoshi4649/GamePartySys/src/common"
)

// Logout ... https://example.com/auth/logout/
func Logout(ctx *gin.Context) {
	common.LogoutFunction(ctx)
	ctx.Redirect(303, "/")
}
