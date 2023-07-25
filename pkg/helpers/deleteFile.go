package helpers

import (
	"fmt"
	"os"
	"strings"
)

func DeleteFile(imgUrl string) bool {
	fileName := strings.Split(imgUrl, "static/img/")[1]

	err := os.Remove(fmt.Sprintf("uploads/img/%s", fileName))
	if err != nil {
		return false
	}

	fmt.Println("File deleted")
	return true
}
