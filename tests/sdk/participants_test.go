package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"iniciador-sdk/iniciador/auth"
	"iniciador-sdk/iniciador/participants"
	"iniciador-sdk/tests/sdk/helpers"
)

func TestGetParticipants(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method and path
		if r.Method != http.MethodGet {
			t.Errorf("expected a GET method, but got %s", r.Method)
		}
		expectedURL := "/participants"
		if r.URL.Path != expectedURL {
			t.Errorf("expected path to be %s, but got %s", expectedURL, r.URL.Path)
		}

		// Verify the authorization header
		authHeader := r.Header.Get("Authorization")
		expectedAuthHeader := "Bearer testAccessToken"
		if authHeader != expectedAuthHeader {
			t.Errorf("expected authorization header to be %s, but got %s", expectedAuthHeader, authHeader)
		}

		// Verify the query parameters
		queryParams := r.URL.Query()
		expectedQueryParams := url.Values{
			"id":                []string{"testID"},
			"name":              []string{"testName"},
			"slug":              []string{"testSlug"},
			"status":            []string{"testStatus"},
			"firstParticipants": []string{"testFirstParticipants"},
			"limit":             []string{"testLimit"},
			"afterCursor":       []string{"testAfterCursor"},
			"beforeCursor":      []string{"testBeforeCursor"},
		}
		if !helpers.AreURLQueryParamsEqual(queryParams, expectedQueryParams) {
			t.Errorf("expected query parameters to be %v, but got %v", expectedQueryParams, queryParams)
		}

		// Create a simulated response
		participantsOutput := participants.ParticipantFilterOutput{
			Data: []participants.ParticipantsPayload{
				{
					ID:     "participant1",
					Slug:   "participant1",
					Name:   "Participant 1",
					Avatar: "avatar1",
				},
				{
					ID:     "participant2",
					Slug:   "participant2",
					Name:   "Participant 2",
					Avatar: "avatar2",
				},
			},
			Cursor: participants.Cursor{
				AfterCursor:  "nextCursor",
				BeforeCursor: "prevCursor",
			},
		}
		responseBody, err := json.Marshal(participantsOutput)
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

	// Set up the filters
	filters := &participants.ParticipantsFilter{
		ID:                "testID",
		Name:              "testName",
		Slug:              "testSlug",
		Status:            "testStatus",
		FirstParticipants: "testFirstParticipants",
		Limit:             "testLimit",
		AfterCursor:       "testAfterCursor",
		BeforeCursor:      "testBeforeCursor",
	}

	// Execute the method to be tested
	output, err := participants.GetParticipants("testAccessToken", filters, authClient)

	// Verify the results
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedOutput := &participants.ParticipantFilterOutput{
		Data: []participants.ParticipantsPayload{
			{
				ID:     "participant1",
				Slug:   "participant1",
				Name:   "Participant 1",
				Avatar: "avatar1",
			},
			{
				ID:     "participant2",
				Slug:   "participant2",
				Name:   "Participant 2",
				Avatar: "avatar2",
			},
		},
		Cursor: participants.Cursor{
			AfterCursor:  "nextCursor",
			BeforeCursor: "prevCursor",
		},
	}
	if !helpers.IsEqual(output, expectedOutput) {
		t.Errorf("expected result: %+v, actual result: %+v", expectedOutput, output)
	}
}
