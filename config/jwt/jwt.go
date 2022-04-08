package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 指定加密秘钥
var jwtSecret = []byte("HUST_1037")

type Claims struct {
	UserName string `json:"userName"`
	UserId   int64  `json:"userId"`
	jwt.StandardClaims
}

// CreateToken 生成token
func CreateToken(userName string, userId int64) (string, error) {
	// 设置有效期
	expiredAt := time.Now().Add(2 * time.Hour)
	claims := Claims{
		UserName: userName,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名
	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 获取过期时间
