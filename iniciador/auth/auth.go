package auth

import (
	"bytes"
	"encoding/json"
	"net/http"

	"iniciador-sdk/iniciador/utils"
)

type AuthOutput struct {
	AccessToken string `json:"accessToken"`
}

type AuthInterfaceOutput struct {
	AccessToken  string `json:"accessToken"`
	InterfaceURL string `json:"interfaceURL"`
	PaymentID    string `json:"paymentId"`
}

type AuthClient struct {
	clientID     string
	clientSecret string
	environment  string
}

func NewAuthClient(clientID, clientSecret, environment string) *AuthClient {
	return &AuthClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		environment:  utils.SetEnvironment(environment),
	}
}

func (c *AuthClient) Auth() (*AuthOutput, error) {
	requestBody := map[string]interface{}{
		"clientId":     c.clientID,
		"clientSecret": c.clientSecret,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(
		c.environment+"/auth",
		"application/json",
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var authOutput AuthOutput
	err = utils.HandleResponse(response, &authOutput)
	if err != nil {
		return nil, err
	}

	return &authOutput, nil
}

func (c *AuthClient) AuthInterface() (*AuthInterfaceOutput, error) {
	requestBody := map[string]interface{}{
		"clientId":     c.clientID,
		"clientSecret": c.clientSecret,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(
		c.environment+"/auth/interface",
		"application/json",
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var authInterfaceOutput AuthInterfaceOutput
	err = utils.HandleResponse(response, &authInterfaceOutput)
	if err != nil {
		return nil, err
	}

	return &authInterfaceOutput, nil
}
