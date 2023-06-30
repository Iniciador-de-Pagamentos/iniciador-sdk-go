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
	ClientID     string
	ClientSecret string
	Environment  string
}

func NewAuthClient(clientID, clientSecret, environment string) *AuthClient {
	return &AuthClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Environment:  utils.SetEnvironment(environment),
	}
}

func (c *AuthClient) GetEnvironment() string {
	return c.Environment
}

func (c *AuthClient) Auth() (*AuthOutput, error) {
	requestBody := map[string]interface{}{
		"clientId":     c.ClientID,
		"clientSecret": c.ClientSecret,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(
		c.Environment+"/auth",
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
		"clientId":     c.ClientID,
		"clientSecret": c.ClientSecret,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(
		c.Environment+"/auth/interface",
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
