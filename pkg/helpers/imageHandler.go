package helpers

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateImgUrl(c *gin.Context, filename string) string {
	var imgUrl string
	if strings.Contains(c.Request.Host, "localhost") || strings.Contains(c.Request.Host, "127.0.0.1") {
		imgUrl = fmt.Sprintf("http://%s/static/img/%s", c.Request.Host, filename)
	} else {
		imgUrl = fmt.Sprintf("https://%s/static/img/%s", c.Request.Host, filename)
	}

	return imgUrl
}

func SaveFile(c *gin.Context, file *multipart.FileHeader) (string, error) {

	// generate randomized filename using timestamps that convert to miliseconds
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))

	// get active directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	// set file location
	fileLocation := filepath.Join(dir, "uploads/img", newFileName)

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, fileLocation)

	// generate image url
	imgUrl := GenerateImgUrl(c, newFileName)

	return imgUrl, err
}
