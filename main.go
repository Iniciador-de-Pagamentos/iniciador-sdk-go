package main

import (
	"fmt"
	"iniciador-sdk/iniciador/auth"
)

func main() {
	clientID := "c82700f8-f0bf-4cce-9068-a2fd6991ee9b"
	clientSecret := "sB#C8ybhJEN63RjBz6Kpd8NUywHkKzXN$d&Zr3j4"
	environment := "dev"

	authClient := auth.NewAuthClient(clientID, clientSecret, environment)

	authOutput, err := authClient.Auth()
	if err != nil {
		fmt.Println("Authentication failed:", err)
		return
	}

	authInterfaceOutput, err := authClient.AuthInterface()
	if err != nil {
		fmt.Println("Authentication failed:", err)
		return
	}

	accessToken := authOutput.AccessToken

	fmt.Println("Auth:", accessToken)
	fmt.Println("Auth Interface:", authInterfaceOutput)
}
