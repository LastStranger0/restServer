package tokens

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	timeTokenAccess = 12 * time.Hour
)

type myClaims struct {
	jwt.StandardClaims
	UserID string
}

func GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &myClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeTokenAccess).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	return token.SignedString([]byte("secretKey"))
}
