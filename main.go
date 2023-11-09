package main

import (
	"fmt"
	"net/http"

    "github.com/gin-gonic/gin"
    "formapp.go/service"
    "formapp.go/service/stateless"
)

// config
const port = 8000

func main() {
    // initialize Gin engine
    engine := gin.Default()
    engine.LoadHTMLGlob("templates/*.html")

    // routing
    engine.GET("/", service.RootHandler)
    engine.GET("/hello", service.HelloHandler)
    engine.GET("/bye", service.ByeHandler)
    engine.GET("/hello.jp", service.HellojpHandler)
    engine.GET("/name-form", service.NameFormHandler)
    engine.POST("/register-name", service.RegisterNameHandler)

    engine.GET("/stateless/start", stateless.Start)
    engine.POST("/stateless/start", notImplemented)
    engine.POST("/stateless/name", notImplemented)
    engine.POST("/stateless/birthday", notImplemented)
    engine.POST("/stateless/message", notImplemented)

    // start server
    engine.Run(fmt.Sprintf(":%d", port))
}

func notImplemented(ctx *gin.Context) {
    msg := fmt.Sprintf("%s to %s is not implemented yet", ctx.Request.Method, ctx.Request.URL)
    ctx.String(http.StatusNotImplemented, msg)
}