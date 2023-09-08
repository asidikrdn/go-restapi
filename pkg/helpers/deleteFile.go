package helpers

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func DeleteFile(imgUrl string) bool {
	fileName := strings.Split(imgUrl, "static/img/")[1]

	err := os.Remove(fmt.Sprintf("uploads/img/%s", fileName))
	if err != nil {
		return false
	}

	log.Println("File " + fileName + " deleted")
	return true
}
