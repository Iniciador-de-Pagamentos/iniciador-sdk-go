package main

import (
	"fmt"
	"iniciador-sdk/iniciador/auth"
	"iniciador-sdk/iniciador/participants"
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

	accessToken := authOutput.AccessToken

	filters := &participants.ParticipantsFilter{}

	participants, err := participants.Participants(accessToken, filters, authClient)
	if err != nil {
		fmt.Println("Get Participants failed:", err)
		return
	}

	fmt.Println("Auth:", accessToken)
	fmt.Println("Participants:", participants)
}
