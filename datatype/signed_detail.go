package datatype

import (
	"github.com/golang-jwt/jwt/v5"
)

type SignedDetails struct {
	ID     int64  `json:"id"`
	UserId string `json:"userid"`
	jwt.RegisteredClaims
}
