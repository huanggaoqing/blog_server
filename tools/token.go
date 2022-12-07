package tools

import (
	"blog_server/constant"
	"blog_server/resp"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	ID       int
	UserName string
	*jwt.StandardClaims
}

var instance *TokenCore

type TokenCore struct {
	jwtSecret []byte
}

// GenerateToken 生成token
func (t *TokenCore) GenerateToken(userId int, userName string) (string, error) {
	nowTime := time.Now()
	claims := &UserClaims{
		ID:       userId,
		UserName: userName,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: nowTime.Add(time.Hour).Unix(),
			Issuer:    userName,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(t.jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

// VerificationToken 验证token
func (t *TokenCore) VerificationToken(token string) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return t.jwtSecret, nil
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

func NewToken() *TokenCore {
	if instance == nil {
		instance = &TokenCore{
			jwtSecret: []byte(constant.JwtSecret),
		}
	}
	return instance
}
