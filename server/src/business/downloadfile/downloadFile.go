package downloadfile

import (
	"github.com/gin-gonic/gin"
	"laynas/server/src/model"
)

type Response struct {
	HttpCode int
	Path     string
}

func DownloadFiles(ctx *gin.Context) (Response, error) {
	path := model.GetFilesPath(ctx.Query("file"))

	return Response{
		HttpCode: 200,
		Path:     path,
	}, nil
}
