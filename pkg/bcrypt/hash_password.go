package bcrypt

import "golang.org/x/crypto/bcrypt"

// hashing password
func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedByte), nil
}
