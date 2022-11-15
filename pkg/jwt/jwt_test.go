package jwt

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	signingKey := []byte("test")
	token, err := NewJWt(signingKey).CreateToken(MyCustomClaims{
		Id:       int64(10001),
		AuthorId: 1234,
	})
	fmt.Println(token, err)
}

func TestParseToken(t *testing.T) {
	signingKey := []byte("test111")
	jwt := NewJWt(signingKey)
	token, err := jwt.CreateToken(MyCustomClaims{
		Id:       int64(10001),
		AuthorId: 1234,
	})
	fmt.Println(token, err)
	claims, err := jwt.ParseToken(token)
	fmt.Println(claims, err)
}
