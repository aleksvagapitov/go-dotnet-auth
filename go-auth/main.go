package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var secretkey string = "JWTAuthenticationHIGHsecuredPasswordVVVp1OH7Xzyr"

func main() {
	userId := uuid.New()
	token, err := GenerateJWT(userId, "Administrator")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(token)
}

func GenerateJWT(userId uuid.UUID, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["iss"] = "http://go-auth"
	claims["aud"] = "http://dotnet-auth"
	claims["usedid"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
