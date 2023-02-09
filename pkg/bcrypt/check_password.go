package bcrypt

import "golang.org/x/crypto/bcrypt"

// compare password
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
