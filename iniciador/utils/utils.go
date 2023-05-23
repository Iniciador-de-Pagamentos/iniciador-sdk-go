package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SetEnvironment(environment string) string {
	switch environment {
	case "dev":
		return "https://consumer.dev.inic.dev/v1"
	case "sandbox":
		return "https://consumer.sandbox.inic.dev/v1"
	case "staging":
		return "https://consumer.staging.inic.dev/v1"
	case "prod":
		return "https://consumer.u4c-iniciador.com.br/v1"
	default:
		panic(fmt.Errorf("Something went wrong, verify environment value."))
	}
}

type Error struct {
	ErrorCode  string   `json:"errorCode"`
	Message    []string `json:"message"`
	Method     string   `json:"method"`
	Path       string   `json:"path"`
	StatusCode int      `json:"statusCode"`
	Timestamp  string   `json:"timestamp"`
}

func HandleResponse(response *http.Response, output interface{}) error {
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		err = json.Unmarshal(bodyBytes, output)
		if err != nil {
			return fmt.Errorf("failed to decode response body: %v", err)
		}
		return nil
	}

	var errResponse Error
	err = json.Unmarshal(bodyBytes, &errResponse)
	if err != nil {
		return fmt.Errorf("failed to decode error response: %v", err)
	}

	if len(errResponse.Message) > 0 {
		return fmt.Errorf("request failed with status code %d: %s", errResponse.StatusCode, strings.Join(errResponse.Message, ", "))
	}

	return fmt.Errorf("request failed with status code %d", response.StatusCode)
}

func MarshalWithoutEmptyFields(payload interface{}) ([]byte, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	nonEmptyFields := make(map[string]interface{})
	for k, v := range m {
		if v != nil {
			switch value := v.(type) {
			case string:
				if value != "" {
					nonEmptyFields[k] = v
				}
			default:
				nonEmptyFields[k] = v
			}
		}
	}

	return json.Marshal(nonEmptyFields)
}
