package participants

import (
	"fmt"
	"iniciador-sdk/iniciador/auth"
	"iniciador-sdk/iniciador/utils"
	"net/http"
	"net/url"
)

type ParticipantsPayload struct {
	ID     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Cursor struct {
	AfterCursor  string `json:"afterCursor"`
	BeforeCursor string `json:"beforeCursor"`
}

type ParticipantFilterOutput struct {
	Data   []ParticipantsPayload `json:"data"`
	Cursor Cursor                `json:"cursor"`
}

type ParticipantsFilter struct {
	ID                string
	Name              string
	Slug              string
	Status            string
	FirstParticipants string
	Limit             string
	AfterCursor       string
	BeforeCursor      string
}

func Participants(accessToken string, filters *ParticipantsFilter, authClient *auth.AuthClient) (*ParticipantFilterOutput, error) {
	environment := authClient.GetEnvironment()
	filterParams := make(url.Values)

	if filters != nil {
		if filters.ID != "" {
			filterParams.Set("id", filters.ID)
		}
		if filters.Name != "" {
			filterParams.Set("name", filters.Name)
		}
		if filters.Slug != "" {
			filterParams.Set("slug", filters.Slug)
		}
		if filters.Status != "" {
			filterParams.Set("status", filters.Status)
		}
		if filters.FirstParticipants != "" {
			filterParams.Set("firstParticipants", filters.FirstParticipants)
		}
		if filters.Limit != "" {
			filterParams.Set("limit", filters.Limit)
		}
		if filters.AfterCursor != "" {
			filterParams.Set("afterCursor", filters.AfterCursor)
		}
		if filters.BeforeCursor != "" {
			filterParams.Set("beforeCursor", filters.BeforeCursor)
		}
	}

	var queryString string
	if len(filterParams) > 0 {
		queryString = filterParams.Encode()
	}

	url := fmt.Sprintf("%s/participants?%s", environment, queryString)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var output ParticipantFilterOutput
	err = utils.HandleResponse(resp, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
