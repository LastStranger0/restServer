package tokens

import (
	"github.com/dgrijalva/jwt-go"
	"log"
)

func ParseTokens(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
	})

	if err != nil {
		log.Println(err)
		return "", err
	}
	claims, _ := token.Claims.(*myClaims)

	return claims.UserID, err
}
