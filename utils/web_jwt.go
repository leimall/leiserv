package utils

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"leiserv/global"
	website "leiserv/models/website/types"
)

type WebJWT struct {
	SigningKey []byte
}

var (
	ErrTokenWebExpired     = errors.New("token is expired")
	ErrTokenWebNotValidYet = errors.New("token not active yet")
	ErrTokenWebMalformed   = errors.New("that's not even a token")
	ErrTokenWebInvalid     = errors.New("couldn't handle this token")
)

func NewWebJWT() *WebJWT {
	return &WebJWT{
		[]byte(global.MALL_CONFIG.JWT.SigningKey),
	}
}

func (j *WebJWT) CreateWebClaims(baseClaims website.BaseClient) website.JWTClaims {
	bf, _ := ParseDuration(global.MALL_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.MALL_CONFIG.JWT.ExpiresTime)
	claims := website.JWTClaims{
		BaseClient: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    global.MALL_CONFIG.JWT.Issuer,             // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *WebJWT) CreateWebToken(claims website.JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *WebJWT) CreateWebTokenByOldToken(oldToken string, claims website.JWTClaims) (string, error) {
	v, err, _ := global.MALL_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateWebToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *WebJWT) ParseWebToken(tokenString string) (*website.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &website.JWTClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenWebMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenWebExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenWebNotValidYet
			} else {
				return nil, ErrTokenWebInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*website.JWTClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenWebInvalid

	} else {
		return nil, ErrTokenWebInvalid
	}
}
