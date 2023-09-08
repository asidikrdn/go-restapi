package middleware

import (
	"bytes"
	"go-restapi/config/postgres"
	"go-restapi/db/models"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logData := createLogData(c)

		c.Next()

		logData.ResposeTime = time.Since(logData.Date).Seconds()
		logData.StatusCode = c.Writer.Status()
		// logData.Message = c.Errors.String()

		db := postgres.DB
		db.Create(&logData)
	}
}

func createLogData(c *gin.Context) models.Log {
	dash := "----------------------------"

	// reqBody, _ := io.ReadAll(c.Request.Body)
	reqBody := new(bytes.Buffer)
	reqBody.ReadFrom(c.Request.Body)
	c.Request.Body = io.NopCloser(reqBody)

	logData := models.Log{
		Date:      time.Now(),
		IPAddress: c.ClientIP(),
		Host:      c.Request.Host,
		Path:      c.Request.URL.Path,
		Method:    c.Request.Method,
	}

	if strings.Contains(reqBody.String(), "password") { // jika dalam req body terdapat key password
		logData.Body = "this data is encrypted, because contains credentials"
	} else if reqBody.String() != "" {
		// split req body, jika json maka tidak akan tersplit, jika form data maka akan tersplit
		body := strings.Split(reqBody.String(), dash)

		var (
			textBody string
			fileBody string
		)

		for _, b := range body {
			if strings.Contains(b, "image") { // jika terdapat gambar pada req body
				fileBody = dash + strings.Split(b, "\r\n\r\n")[0]
			} else {
				if len(b) >= 1 && b[0] == '{' { // jika berbentuk json (ditandai dengan awalnya adalah '{' )
					textBody += b
				} else if b != "" { // jika datanya adalah form-data
					textBody += dash + b
				}
			}
		}

		logData.Body = textBody
		logData.File = fileBody
	}

	return logData
}
