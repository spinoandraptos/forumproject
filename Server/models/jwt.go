package models

import (
	"github.com/go-chi/jwtauth"
)

func Createtoken(username string, password string) string {
	var authtoken *jwtauth.JWTAuth
	const secretkey = "123abc"
	authtoken = jwtauth.New("HS256", []byte(secretkey), nil)
	_, payloadclaims, _ := authtoken.Encode(map[string]interface{}{"username": username, "password": password})
	return payloadclaims
}
