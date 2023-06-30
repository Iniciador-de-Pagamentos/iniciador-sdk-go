package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"iniciador-sdk/iniciador/auth"
	"iniciador-sdk/iniciador/payments"
	"iniciador-sdk/tests/sdk/helpers"
)

func TestSendGetStatus(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method and path
		expectedURL := "/payments"
		if r.URL.Path != expectedURL {
			t.Errorf("expected path to be %s, but got %s", expectedURL, r.URL.Path)
		}

		// Verify the authorization header
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer testAccessToken"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected authorization header to be %s, but got %s", expectedAuthHeader, authHeader)
		}

		// Verify the request body
		var payment payments.PaymentInitiationPayload
		err := json.NewDecoder(r.Body).Decode(&payment)
		if err != nil {
			t.Errorf("failed to decode the request body: %v", err)
		}

		expectedPayment := payments.PaymentInitiationPayload{
			ID:     "testID",
			Amount: 100.0,
		}
		if !helpers.IsEqual(payment, expectedPayment) {
			t.Errorf("expected request body: %+v, actual request body: %+v", expectedPayment, payment)
		}

		// Create a simulated response
		status := payments.PaymentInitiationStatus(payments.Started)
		paymentInitiationPayload := payments.PaymentInitiationPayload{
			ID:     "testID",
			Status: &status,
		}
		responseBody, err := json.Marshal(paymentInitiationPayload)
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

	// Set up the payment payload
	payment := &payments.PaymentInitiationPayload{
		ID:     "testID",
		Amount: 100.0,
	}

	// Execute the Send function
	sentPayment, err := payments.Send("testAccessToken", payment, authClient)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the sent payment
	status := payments.PaymentInitiationStatus(payments.Started)
	expectedSentPayment := &payments.PaymentInitiationPayload{
		ID:     "testID",
		Status: &status,
	}
	if !helpers.IsEqual(sentPayment, expectedSentPayment) {
		t.Errorf("expected sent payment: %+v, actual sent payment: %+v", expectedSentPayment, sentPayment)
	}

	// Create a test server for the Get and Status functions
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method
		if r.Method != http.MethodGet {
			t.Errorf("expected a GET method, but got %s", r.Method)
		}

		// Verify the authorization header
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer testAccessToken"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected authorization header to be %s, but got %s", expectedAuthHeader, authHeader)
		}

		// Verify the request URL
		expectedURL := fmt.Sprintf("/payments/%s", payment.ID)
		if r.URL.Path != expectedURL {
			t.Errorf("expected path to be %s, but got %s", expectedURL, r.URL.Path)
		}

		// Create a simulated response
		status := payments.PaymentInitiationStatus(payments.Started)
		paymentPayload := payments.PaymentInitiationPayload{
			ID:     payment.ID,
			Status: &status,
		}
		responseBody, err := json.Marshal(paymentPayload)
		if err != nil {
			t.Errorf("failed to encode the response: %v", err)
		}

		// Send the simulated response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseBody)
	}))
	defer server.Close()

	// Override the environment URL with the test server's URL
	authClient.Environment = server.URL

	// Execute the Get function
	retrievedPayment, err := payments.Get("testAccessToken", authClient)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the retrieved payment
	expectedRetrievedPayment := &payments.PaymentInitiationPayload{
		ID:     payment.ID,
		Status: &status,
	}
	if !helpers.IsEqual(retrievedPayment, expectedRetrievedPayment) {
		t.Errorf("expected retrieved payment: %+v, actual retrieved payment: %+v", expectedRetrievedPayment, retrievedPayment)
	}

	// Create a test server for the Status function
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method
		if r.Method != http.MethodGet {
			t.Errorf("expected a GET method, but got %s", r.Method)
		}

		// Verify the authorization header
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer testAccessToken"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected authorization header to be %s, but got %s", expectedAuthHeader, authHeader)
		}

		// Verify the request URL
		expectedURL := fmt.Sprintf("/payments/%s/status", payment.ID)
		if r.URL.Path != expectedURL {
			t.Errorf("expected path to be %s, but got %s", expectedURL, r.URL.Path)
		}

		// Create a simulated response
		statusPayload := payments.PaymentStatusPayload{
			ID:     payment.ID,
			Status: string(payments.Started),
		}
		responseBody, err := json.Marshal(statusPayload)
		if err != nil {
			t.Errorf("failed to encode the response: %v", err)
		}

		// Send the simulated response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseBody)
	}))
	defer server.Close()

	// Override the environment URL with the test server's URL
	authClient.Environment = server.URL

	// Execute the Status function
	paymentStatus, err := payments.Status("testAccessToken", authClient)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify the payment status
	expectedPaymentStatus := &payments.PaymentStatusPayload{
		ID:     payment.ID,
		Status: string(payments.Started),
	}
	if !helpers.IsEqual(paymentStatus, expectedPaymentStatus) {
		t.Errorf("expected payment status: %+v, actual payment status: %+v", expectedPaymentStatus, paymentStatus)
	}
}
