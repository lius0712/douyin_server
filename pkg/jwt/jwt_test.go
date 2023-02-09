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
	signingKey := []byte("DouYinKey")
	jwt := NewJWt(signingKey)
	token, err := jwt.CreateToken(MyCustomClaims{
		AuthorId: 20,
	})
	fmt.Println(token, err)
	claims, err := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MjAsIkF1dGhvcklkIjowfQ.2S5kc71ENwm7dC5OnPiGrX6hN5mDJ8ontMObiFADM98")
	fmt.Println(claims, err)
}
