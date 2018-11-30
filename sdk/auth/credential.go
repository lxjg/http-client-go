package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credential struct {
	AccessKeyID     string
	AccessKeySecret string
}

func GenerateJWT(credential *Credential) (string, error) {
	var sm jwt.SigningMethod = jwt.SigningMethodHS256
	claims := jwt.MapClaims{
		"id":  credential.AccessKeyID,
		"exp": time.Now().Add(time.Second * 30).Unix(),
	}
	return jwt.NewWithClaims(sm, claims).SignedString([]byte(credential.AccessKeySecret))
}
