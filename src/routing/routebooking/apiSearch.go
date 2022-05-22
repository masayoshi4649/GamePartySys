package routebooking

import (
	"github.com/gin-gonic/gin"
	"github.com/masayoshi4649/GamePartySys/src/dbpsql"
)

// APISearch ... GET root
func APISearch(ctx *gin.Context) {
	dbdata := dbpsql.SelectInstanceInfo()
	ctx.JSON(200, dbdata)
}
