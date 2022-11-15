package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/lius0712/douyin_server/pkg/errno"
)

type JWT struct {
	SigningKey []byte
}

type MyCustomClaims struct {
	Id       int64
	AuthorId int64
	jwt.StandardClaims
}

func NewJWt(SigningKey []byte) *JWT {
	return &JWT{
		SigningKey: SigningKey,
	}
}

func (j *JWT) CreateToken(claims MyCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, err := token.SignedString(j.SigningKey)
	return res, err
}

func (j *JWT) ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				//That's not even a token,令牌格式错误
				return nil, errno.ErrorMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				//Token is either expired
				return nil, errno.ErrorExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				//Token is not active yet
				return nil, errno.ErrorNotValidYet
			} else {
				//Couldn't handle this token
				return nil, errno.ErrorTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errno.ErrorClaimsInvalid
}
