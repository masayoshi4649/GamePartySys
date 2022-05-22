package routeroot

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Index ... GET root
func Index(ctx *gin.Context) {
	opt := ctx.Query("opt")
	fmt.Println("opt=>", opt)

	ctx.HTML(200, "root/index.html", nil)
}
