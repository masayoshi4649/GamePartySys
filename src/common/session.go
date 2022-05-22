package common

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginFuncion ... ログイン後のセッション作成
func LoginFuncion(ctx *gin.Context, userid string) {
	session := sessions.Default(ctx)
	session.Set("userid", userid)
	session.Save()
}

// LoginCheck ... ログイン済か否か
func LoginCheck(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	userid := session.Get("userid")

	return (userid != nil)
}

// LogoutFunction ... logout Clear Session
func LogoutFunction(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	println("Session clear")
}
