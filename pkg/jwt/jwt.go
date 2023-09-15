package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// GetJwtToken
// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func GetJwtToken(secretKey string, iat, seconds int64, id int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["id"] = id
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetId(secretKey string, tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	seconds, ok := claims["exp"].(float64)
	if !ok {
		return 0, errors.New("token error")
	}
	if int64(seconds) < time.Now().Unix() {
		return 0, errors.New("token error")
	}
	id, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("token error")
	}
	return int64(id), nil
}
