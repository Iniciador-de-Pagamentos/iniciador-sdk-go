package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"iniciador-sdk/iniciador/auth"
	"iniciador-sdk/tests/sdk/helpers"
)

func TestAuthClient_Auth(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method and path
		if r.Method != http.MethodPost {
			t.Errorf("expected a POST method, but got %s", r.Method)
		}
		if r.URL.Path != "/auth" {
			t.Errorf("expected path to be /auth, but got %s", r.URL.Path)
		}

		// Decode the request body
		var requestBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			t.Errorf("failed to decode the request body: %v", err)
		}

		// Verify the fields in the request body
		clientID, ok := requestBody["clientId"].(string)
		if !ok || clientID != "testClientID" {
			t.Errorf("invalid value for the clientId field")
		}

		clientSecret, ok := requestBody["clientSecret"].(string)
		if !ok || clientSecret != "testClientSecret" {
			t.Errorf("invalid value for the clientSecret field")
		}

		// Create a simulated response
		authOutput := auth.AuthOutput{
			AccessToken: "testAccessToken",
		}
		responseBody, err := json.Marshal(authOutput)
		if err != nil {
			t.Errorf("failed to encode the response: %v", err)
		}

		// Send the simulated response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseBody)
	}))
	defer server.Close()

	// Configure the authentication client for the test
	authClient := auth.NewAuthClient("testClientID", "testClientSecret", "dev")

	// Override the environment URL with the test server's URL
	authClient.Environment = server.URL

	// Execute the method to be tested
	authOutput, err := authClient.Auth()

	// Verify the results
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedOutput := &auth.AuthOutput{
		AccessToken: "testAccessToken",
	}
	if !helpers.IsEqual(authOutput, expectedOutput) {
		t.Errorf("expected result: %+v, actual result: %+v", expectedOutput, authOutput)
	}
}

func TestAuthClient_AuthInterface(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method and path
		if r.Method != http.MethodPost {
			t.Errorf("expected a POST method, but got %s", r.Method)
		}
		if r.URL.Path != "/auth/interface" {
			t.Errorf("expected path to be /auth/interface, but got %s", r.URL.Path)
		}

		// Decode the request body
		var requestBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			t.Errorf("failed to decode the request body: %v", err)
		}

		// Verify the fields in the request body
		clientID, ok := requestBody["clientId"].(string)
		if !ok || clientID != "testClientID" {
			t.Errorf("invalid value for the clientId field")
		}

		clientSecret, ok := requestBody["clientSecret"].(string)
		if !ok || clientSecret != "testClientSecret" {
			t.Errorf("invalid value for the clientSecret field")
		}

		// Create a simulated response
		authInterfaceOutput := auth.AuthInterfaceOutput{
			AccessToken:  "testAccessToken",
			InterfaceURL: "testInterfaceURL",
			PaymentID:    "testPaymentID",
		}
		responseBody, err := json.Marshal(authInterfaceOutput)
		if err != nil {
			t.Errorf("failed to encode the response: %v", err)
		}

		// Send the simulated response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseBody)
	}))
	defer server.Close()

	// Configure the authentication client for the test
	authClient := auth.NewAuthClient("testClientID", "testClientSecret", "dev")

	// Override the environment URL with the test server's URL
	authClient.Environment = server.URL

	// Execute the method to be tested
	authInterfaceOutput, err := authClient.AuthInterface()

	// Verify the results
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedOutput := &auth.AuthInterfaceOutput{
		AccessToken:  "testAccessToken",
		InterfaceURL: "testInterfaceURL",
		PaymentID:    "testPaymentID",
	}
	if !helpers.IsEqual(authInterfaceOutput, expectedOutput) {
		t.Errorf("expected result: %+v, actual result: %+v", expectedOutput, authInterfaceOutput)
	}
}
