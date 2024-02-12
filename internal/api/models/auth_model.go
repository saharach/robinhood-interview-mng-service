package models

import "github.com/dgrijalva/jwt-go"

// Claims struct for JWT claims
type Claims struct {
	UserId   int    `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
