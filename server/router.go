package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/marchar/fabric-entry/server/handlers"

	"github.com/gin-gonic/gin"
)

func newRouter() http.Handler {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/-/health", handlers.Health)
	return r
}
