package jwt

import (
	"github.com/golang-jwt/jwt"
	"gvb_blog/global"
	"gvb_blog/models"
	"time"
)

var jwtKey = []byte("asdfghjkl")

type Claims struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// ReleaseToken 生成token
func ReleaseToken(user models.UserModel) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId:   user.ID,
		UserName: user.UserName,
		Role:     user.Role.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    global.Config.Jwt.Issuer,
			Subject:   global.Config.Jwt.Subject,
		},
	}
	// 生成内容
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 加上前缀
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
