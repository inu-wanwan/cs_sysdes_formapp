package stateless

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Start(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "start.html", gin.H{ "Target": "/stateless/start" })
}