package savefiles

import (
	"github.com/gin-gonic/gin"
	"laynas/server/src/business"
	"laynas/server/src/model"
	"net/http"
)

type Response struct {
	HttpCode int
	Response interface{}
}

func SavesFiles(ctx *gin.Context) (Response, error) {
	if err := model.SaveFiles(ctx); err != nil {
		return Response{}, business.ErrSavingFiles
	}

	return Response{
		HttpCode: http.StatusOK,
		Response: gin.H{"msg": "OK"},
	}, nil
}
