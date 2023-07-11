package utils

import "golang.org/x/crypto/bcrypt"

func Hash(value string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func CompareHash(hashedPassword, unHashedPassword string) bool {
	hashed := []byte(hashedPassword)
	unHashed := []byte(unHashedPassword)
	err := bcrypt.CompareHashAndPassword(hashed, unHashed)
	if err != nil {
		return false
	}
	return true
}
