package tools

import (
	"blog_server/resp"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	ID       int
	UserName string
	*jwt.StandardClaims
}

var jwtSecret = []byte("myBlog")

func GenerateToken(userId int, userName string) (string, error) {
	nowTime := time.Now()
	claims := &UserClaims{
		ID:       userId,
		UserName: userName,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: nowTime.Add(time.Hour).Unix(),
			Issuer:    userName,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerificationToken(token string) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			switch v.Errors {
			case jwt.ValidationErrorMalformed:
				return nil, resp.NOT_TOKEN
			case jwt.ValidationErrorExpired:
				return nil, resp.TOKEN_EXPIRED
			case jwt.ValidationErrorNotValidYet:
				return nil, resp.TOKEN_NOT_ACTIVE
			default:
				return nil, resp.TOKEN_INVAILD
			}
		}
	}
	if tokenClaims != nil {
		if v, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return v, nil
		} else {
			return nil, resp.TOKEN_INVAILD
		}
	}
	return nil, resp.TOKEN_INVAILD
}
