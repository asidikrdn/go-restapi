package helpers

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// generate random number with custom length
func GenerateRandomOTP(length int) string {
	rand.Seed(time.Now().Unix())

	token := make([]string, length)
	for i := range token {
		token[i] = strconv.Itoa(rand.Intn(10))
	}
	return strings.Join(token, "")
}
