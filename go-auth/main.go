package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	JWT struct {
		Secret   string `yaml:"secret" env:"SECRET"`
		Issuer   string `yaml:"issuer" env:"ISSUER"`
		Audience string `yaml:"audience" env:"AUDIENCE"`
	}
}

func main() {
	var cfg Config
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Panic("Error Reading Config: ", err)
	}

	userId := uuid.New()
	token, err := GenerateJWT(cfg.JWT.Secret, cfg.JWT.Issuer, cfg.JWT.Audience, userId, "Administrator")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(token)
}

func GenerateJWT(secret string, issuer string, audience string, userId uuid.UUID, role string) (string, error) {
	var mySigningKey = []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["iss"] = issuer
	claims["aud"] = audience
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
