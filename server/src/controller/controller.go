package controller

import (
	"github.com/gin-gonic/gin"
	"laynas/server/src/business/downloadfile"
	"laynas/server/src/business/getfiles"
	"laynas/server/src/business/savefiles"
	"net/http"
)

func DownloadFile() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		resp, err := downloadfile.DownloadFiles(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		if resp.HttpCode == http.StatusOK {
			ctx.FileAttachment(resp.Path, "test.Png")
		}
	}
}

func GetFiles() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		resp, err := getfiles.GetFiles(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(resp.HttpCode, resp.Response)
	}
}

func UploadMultipleFiles() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		resp, err := savefiles.SavesFiles(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(resp.HttpCode, resp.Response)
	}
}
