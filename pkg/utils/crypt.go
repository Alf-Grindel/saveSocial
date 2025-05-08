package utils

import (
	"github.com/alg-grindel/saveSocial/pkg/constant"
	"github.com/alg-grindel/saveSocial/pkg/errno"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	password = password + constant.Salt
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[ERROR] utils.crypt: generate password failed")
		return "", errno.SystemErr
	}
	return string(b), nil
}

func ComparePassword(hashPassword, password string) (bool, error) {
	password = password + constant.Salt
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		log.Println("[ERROR] utils.crypt: compare password failed")
		return false, errno.SystemErr
	}
	return true, nil
}
