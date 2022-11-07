package headers

import (
	"github.com/gin-gonic/gin"
)

func ReadHeader(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Authorization")
}
