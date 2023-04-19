package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

// Claims godoc
type Claims struct {
	UserID  int64
	GroupID int64
	jwt.StandardClaims
}

// GenerateToken returns the token used for auth
func GenerateToken(userID, groupID int64) (token string, err error) {
	nowTime := time.Now().UTC()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		userID,
		groupID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "derems",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}

// ParseToken godoc
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
