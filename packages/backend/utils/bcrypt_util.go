package utils

import "golang.org/x/crypto/bcrypt"

func Encrpt(stringToEncrypt string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(stringToEncrypt), bcrypt.DefaultCost)
}

func ValidateEncrypt(stringToValidate string, stringEncrypted string) error {
	return bcrypt.CompareHashAndPassword([]byte(stringEncrypted), []byte(stringToValidate))
}
