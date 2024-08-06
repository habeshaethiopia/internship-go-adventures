package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
    ID    string `json:"id"`
    Role  string `json:"role"`
    jwt.StandardClaims
}