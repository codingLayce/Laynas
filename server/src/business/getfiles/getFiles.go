package getfiles

import (
	"github.com/gin-gonic/gin"
	"laynas/server/src/model"
	"net/http"
)

type Response struct {
	HttpCode int
	Response interface{}
}

func GetFiles(_ *gin.Context) (Response, error) {
	root := model.BuildFileTree()

	return Response{
		HttpCode: http.StatusOK,
		Response: root,
	}, nil
}
