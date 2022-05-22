package routeauth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/masayoshi4649/GamePartySys/src/common"
	"github.com/masayoshi4649/GamePartySys/src/discord"
)

// Oauth ... https://example.com/auth/oauth?code=*
func Oauth(ctx *gin.Context) {
	// コード取得
	code := ctx.Query("code")
	fmt.Println("code=>", code)

	if code != "" {
		discordid := discord.GetToken(code)

		if discordid != "" {
			common.LoginFuncion(ctx, discordid)
			ctx.Redirect(303, "/")
		} else {
			// IDが取れない場合はログアウト
			Logout(ctx)
		}
	} else {
		// コードなしはログアウト
		Logout(ctx)
	}
}
