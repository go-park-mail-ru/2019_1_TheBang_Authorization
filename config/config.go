package config

import (
	"os"
)

var Logger = NewGlobalLogger()

var (
	CookieName = "bang_token"
	SECRET     = getSecret()
	PORT       = getPort()
)

func getSecret() []byte {
	secret := []byte(os.Getenv("SECRET"))
	if string(secret) == "" {
		Logger.Warn("There is no SECRET!")
		secret = []byte("secret")
	}

	return secret
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		Logger.Warn("There is no PORT!")
		port = "50051"
	}
	return port
}
