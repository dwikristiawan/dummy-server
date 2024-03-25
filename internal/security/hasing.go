package security

import (
	"golang.org/x/crypto/bcrypt"
)

func StrHashing(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CompareHashingData(hashDataStr string, DataStr string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashDataStr), []byte(DataStr))
}
