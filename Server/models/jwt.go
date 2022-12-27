package models

import (
	"fmt"

	"github.com/go-chi/jwtauth"
)

func Createtoken(username *string, password *string) string {
	var authtoken *jwtauth.JWTAuth
	const secretkey = "123abc"
	authtoken = jwtauth.New("HS256", []byte(secretkey), nil)
	_, payloadclaims, _ := authtoken.Encode(map[string]interface{}{"username": username, "password": password})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", payloadclaims)
	return payloadclaims
}
