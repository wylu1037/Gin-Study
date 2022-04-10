package jwt

import (
	"fmt"
	"testing"
)

// 测试创建token
func TestCreateToken(t *testing.T) {
	token, _ := CreateToken("JayChou", 1)
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IkpheUNob3UiLCJ1c2VySWQiOjEsImV4cCI6MTY0OTQzMDI1MH0.2EUI-heKf8WuZU4qkaVbLDzM6BTpZBcCXO7QKIY_nPA"
	claims, _ := ParseToken(token)
	fmt.Println(claims)
}
