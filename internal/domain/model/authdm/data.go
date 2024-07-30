// Package authdm contains the value object used in the authentication.
package authdm

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	jwt.RegisteredClaims
}
