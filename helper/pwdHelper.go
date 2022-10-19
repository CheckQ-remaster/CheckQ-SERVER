package helper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PwdHelper interface {
	HashPassword(string) (string, error)
	CheckPasswordHash(string, string) bool
}

// 비밀번호를 해싱
func HashPassword(pw string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(bytes), err
}

// 비밀번호를 해싱해 해싱값과 같은지 확인
func CheckPasswordHash(hashVal, userPw string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))

	fmt.Println(err)

	if err != nil {
		return false
	} else {
		return true
	}
}
