package tool

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT详细用法请看：
// https://dinglz.cn/p/jwt%E7%9A%84%E5%BA%94%E7%94%A8/

type JWTUserModel struct {
	ID        uint
	Telephone string
	jwt.RegisteredClaims
}

// 生成token
func GenerateJWToken(jwtPassword string, id uint, telephone string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTUserModel{
		id,
		telephone,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})
	t, _ := token.SignedString([]byte(jwtPassword))
	return t
}

// 解析token
func ParseJWToken(jwtpassword string, token string) (*JWTUserModel, error) {
	tokenAfter, e := jwt.ParseWithClaims(token, &JWTUserModel{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtpassword), nil
	})
	if e != nil {
		return &JWTUserModel{}, e
	}
	cliamAfter := tokenAfter.Claims.(*JWTUserModel)
	return cliamAfter, nil
}
