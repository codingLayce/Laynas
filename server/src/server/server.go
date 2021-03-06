package server

import (
	"github.com/gin-gonic/gin"
	"laynas/server/src/controller"
	"laynas/server/src/middlewares"
)

const (
	fileEndpoint         = "/files"
	fileDownloadEndpoint = "/files/download"
)

func Create() *gin.Engine {
	router := gin.Default()

	createRoutes(router)

	return router
}

func createRoutes(engine *gin.Engine) {
	engine.GET(fileEndpoint, middlewares.WithAuthentication(), controller.GetFiles())
	engine.POST(fileEndpoint, middlewares.WithAuthentication(), controller.UploadMultipleFiles())
	engine.GET(fileDownloadEndpoint, middlewares.WithAuthentication(), controller.DownloadFile())
}
