package downloadfile

import (
	"github.com/gin-gonic/gin"
	"laynas/server/src/model"
	"os"
)

type Response struct {
	HttpCode int
	Path     string
	IsDir    bool
}

func DownloadFiles(ctx *gin.Context) (Response, error) {
	path := model.GetFilesPath(ctx.Query("file"))

	file, err := os.Open(path)
	if err != nil {
		return Response{}, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return Response{}, err
	}

	return Response{
		HttpCode: 200,
		Path:     path,
		IsDir:    fileInfo.IsDir(),
	}, nil
}
