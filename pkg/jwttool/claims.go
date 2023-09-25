package jwttool

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaims struct {
	UID uint `json:"uid,omitempty"`
	jwt.RegisteredClaims
}

func BuildUserClaims(uid uint, et time.Duration) UserClaims {
	return UserClaims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(et)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
