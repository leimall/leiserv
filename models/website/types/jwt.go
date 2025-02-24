package website

import (
	"github.com/gofrs/uuid/v5"
	jwt "github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	BaseClient
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClient struct {
	UUID       uuid.UUID
	ID         uint
	UserId     string
	Email      string
	Username   string
	Permission string
}
