package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"go-restapi-boilerplate/dto"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// single file upload middleware with id as parameter and used for generating filename
func UploadSingleFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//  parsing form with max memory size 8 Mb
		errParsing := r.ParseMultipartForm(8192)
		if errParsing != nil {
			response := dto.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: errParsing.Error(),
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// single file
		uploadedFile, handler, err := r.FormFile("image")

		// if file doesn't exist
		if err != nil {
			// set up context value and send it to next handler
			imageCtx := context.WithValue(r.Context(), "image", "")
			next.ServeHTTP(w, r.WithContext(imageCtx))
			return
		}
		defer uploadedFile.Close()

		// validation format file
		if filepath.Ext(handler.Filename) != ".jpg" && filepath.Ext(handler.Filename) != ".jpeg" && filepath.Ext(handler.Filename) != ".png" {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: "Invalid file type",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		// generate randomized filename using timestamps that convert to miliseconds
		newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename))

		// get active directory
		dir, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}

		// set file location
		fileLocation := filepath.Join(dir, "uploads/img", newFileName)

		// Create new file and open it
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			response := dto.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		defer targetFile.Close()
		// Copy uploded image to created file
		if _, err := io.Copy(targetFile, uploadedFile); err != nil {
			response := dto.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		var imgUrl string
		if strings.Contains(r.Host, "localhost") || strings.Contains(r.Host, "127.0.0.1") {
			imgUrl = fmt.Sprintf("http://%s/static/img/%s", r.Host, newFileName)
		} else {
			imgUrl = fmt.Sprintf("https://%s/static/img/%s", r.Host, newFileName)
		}

		// set up context value and send it to next handler
		imageCtx := context.WithValue(r.Context(), "image", imgUrl)
		next.ServeHTTP(w, r.WithContext(imageCtx))
	})
}

// multiple file upload middleware with id as parameter and used for generating filename
func UploadMultipleFiles(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var arrImages []string

		//  parsing form with max memory size 8 Mb
		errParsing := r.ParseMultipartForm(8192)
		if errParsing != nil {
			response := dto.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: errParsing.Error(),
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// parsing multipart form data
		files := r.MultipartForm.File["images"]

		// if file doesn't exist
		if len(files) <= 0 {
			// set up context value and send it to next handler
			imageCtx := context.WithValue(r.Context(), "image", []string{})
			next.ServeHTTP(w, r.WithContext(imageCtx))
			return
		}

		for _, file := range files {
			// open the file
			openedFile, err := file.Open()
			if err != nil {
				response := dto.ErrorResult{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}
			defer openedFile.Close()

			// validation format file
			if filepath.Ext(file.Filename) != ".jpg" && filepath.Ext(file.Filename) != ".jpeg" && filepath.Ext(file.Filename) != ".png" {
				response := dto.ErrorResult{
					Status:  http.StatusBadRequest,
					Message: "Invalid file type",
				}
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
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

			// Create new file and open it
			targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				response := dto.ErrorResult{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}
			defer targetFile.Close()
			// Copy uploded image to created file
			if _, err := io.Copy(targetFile, openedFile); err != nil {
				response := dto.ErrorResult{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}

			var imgUrl string
			if strings.Contains(r.Host, "localhost") || strings.Contains(r.Host, "127.0.0.1") {
				imgUrl = fmt.Sprintf("http://%s/static/img/%s", r.Host, newFileName)
			} else {
				imgUrl = fmt.Sprintf("https://%s/static/img/%s", r.Host, newFileName)
			}

			arrImages = append(arrImages, imgUrl)
		}

		// set up context value and send it to next handler
		imageCtx := context.WithValue(r.Context(), "image", arrImages)
		next.ServeHTTP(w, r.WithContext(imageCtx))
	})
}
