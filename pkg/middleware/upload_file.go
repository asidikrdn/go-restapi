package middleware

import (
	"fmt"
	"go-restapi-boilerplate/dto"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// single file upload middleware with id as parameter and used for generating filename
func UploadSingleFile() gin.HandlerFunc {
	return func(c *gin.Context) {

		//  parsing form with max memory size 8 Mb
		errParsing := c.Request.ParseMultipartForm(8192)
		if errParsing != nil {
			response := dto.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: errParsing.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		// single file
		file, err := c.FormFile("image")

		// if file doesn't exist
		if err != nil {
			// set up context value and send it to next handler
			c.Set("image", "")
			c.Next()
			return
		}

		log.Println(file.Filename)

		// validation format file
		if filepath.Ext(file.Filename) != ".jpg" && filepath.Ext(file.Filename) != ".jpeg" && filepath.Ext(file.Filename) != ".png" {
			response := dto.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: "Invalid file type",
			}
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

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
		if err != nil {
			response := dto.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		// set up context value and send it to next handler
		c.Set("image", fmt.Sprintf("%s/static/img/%s", os.Getenv("BASE_URL"), newFileName))
		c.Next()
	}
}

// multiple file upload middleware with id as parameter and used for generating filename
func UploadMultipleFiles() gin.HandlerFunc {
	return func(c *gin.Context) {

		var arrImages []string

		//  parsing form with max memory size 8 Mb
		errParsing := c.Request.ParseMultipartForm(8192)
		if errParsing != nil {
			response := dto.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: errParsing.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		// parsing multipart form data
		form, _ := c.MultipartForm()
		files := form.File["images"]

		// if file doesn't exist
		if len(form.File) <= 0 {
			// set up context value and send it to next handler
			c.Set("images", []string{})
			c.Next()
			return
		}

		for _, file := range files {
			log.Println(file.Filename)

			// validation format file
			if filepath.Ext(file.Filename) != ".jpg" && filepath.Ext(file.Filename) != ".jpeg" && filepath.Ext(file.Filename) != ".png" {
				response := dto.ErrorResult{
					Status:  http.StatusBadRequest,
					Message: "Invalid file type",
				}
				c.JSON(http.StatusBadRequest, response)
				c.Abort()
				return
			}

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
			if err != nil {
				response := dto.ErrorResult{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				}
				c.JSON(http.StatusBadRequest, response)
				c.Abort()
				return
			}

			arrImages = append(arrImages, fmt.Sprintf("%s/static/img/%s", os.Getenv("BASE_URL"), newFileName))
		}

		// set up context value and send it to next handler
		c.Set("images", arrImages)
		c.Next()
	}
}
