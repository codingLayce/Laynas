package model

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

const storageRootPath = "C:/Laynas_storage/"

func GetFilesPath(path string) string {
	return storageRootPath + path
}

func SaveFiles(ctx *gin.Context) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println("Error while retrieving MultipartForm : " + err.Error())
		return err
	}

	files := form.File

	errs, _ := errgroup.WithContext(context.Background())

	for subPath, filesHeader := range files {
		for _, file := range filesHeader {
			filePath := buildFilepath(subPath, file.Filename)

			fileToSave := file

			errs.Go(func() error {
				return saveFile(ctx, filePath, fileToSave)
			})
		}
	}

	return errs.Wait()
}

func saveFile(ctx *gin.Context, filePath string, fileToSave *multipart.FileHeader) error {
	log.Printf("[%s] Writing...\n", filePath)

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		log.Printf("[%s] Error while creating filepath : %s\n", filePath, err.Error())

		return err
	}

	err = ctx.SaveUploadedFile(fileToSave, filePath)
	if err != nil {
		log.Printf("[%s] Error saving the file : %s\n", filePath, err.Error())

		return err
	}

	log.Printf("[%s] Done !\n", filePath)

	return nil
}

func buildFilepath(subPath, filename string) string {
	return storageRootPath + subPath + "/" + filename
}
