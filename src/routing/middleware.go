package routing

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/masayoshi4649/GamePartySys/src/common"
	"github.com/masayoshi4649/GamePartySys/src/routing/routeauth"
)

// LoginCheck ... ログイン確認
func LoginCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ルーティング内の処理の前に実行される
		login := common.LoginCheck(ctx)

		fmt.Println("login=>", login)

		switch login {
		case true:
			ctx.Next()

		case false:
			fmt.Println("checkIgnorePath=>", checkIgnorePath(ctx.Request.URL.Path))

			if checkIgnorePath(ctx.Request.URL.Path) == false {
				routeauth.Login(ctx)
				ctx.Abort()
			}
			ctx.Next()
		}
	}
}

// checkIgnorePath ... 認証を抜けるパス一覧
// セッションチェックを無視する場合はtrue
func checkIgnorePath(path string) bool {
	if path == "/favicon.ico" {
		return true
	}

	if strings.HasPrefix(path, "/auth/css/") == true {
		return true
	}
	if strings.HasPrefix(path, "/auth/js/") == true {
		return true
	}

	if path == "/auth/oauth" {
		return true
	}

	return false
}
