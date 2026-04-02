package snap

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

func (r *HttpRequest) Request(headers map[string]string, requestURL string, requestBody interface{}, httpMethod string) ([]byte, error) {
	if r.HttpClient == nil {
		r.HttpClient = &http.Client{}
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, &NicepayError{
			Message: fmt.Sprintf("fail to parse 'body parameters' as JSON: %v", err),
		}
	}

	req, err := http.NewRequest(httpMethod, requestURL, bytes.NewBuffer(bodyBytes))
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

func (r *HttpRequest) RequestPayment(headers map[string]string, requestURL string, requestBody []byte, httpMethod string) ([]byte, error) {
	if r.HttpClient == nil {
		r.HttpClient = &http.Client{}
	}

	req, err := http.NewRequest(httpMethod, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBody, nil
}
