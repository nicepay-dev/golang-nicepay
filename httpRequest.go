package library

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NicepayError struct {
	Message string
}

func (e *NicepayError) Error() string {
	return e.Message
}

type HttpRequest struct {
	HttpClient *http.Client
}

func (r *HttpRequest) Request(headers map[string]string, requestURL string, requestBody interface{}) ([]byte, error) {
	if r.HttpClient == nil {
		r.HttpClient = &http.Client{}
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, &NicepayError{
			Message: fmt.Sprintf("fail to parse 'body parameters' as JSON: %v", err),
		}
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header[key] = []string{value}
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return respBody, nil
}
