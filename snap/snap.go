package snap

import (
	"encoding/json"
	"fmt"
)

type Snap struct {
	ApiConfig  Config
	HttpClient HttpRequest
	Helper     Helper
}

func (s *Snap) RequestSnapAccessToken(parameter map[string]interface{}) (*ResponseAccessToken, error) {
	var response ResponseAccessToken
	formattedDate := s.Helper.GetFormattedDate()
	url := fmt.Sprintf("%s/v1.0/access-token/b2b", s.ApiConfig.GetSnapAPIBaseURL())
	stringToSign := fmt.Sprintf("%s|%s", s.ApiConfig.ClientID, formattedDate)

	bytes := []byte(stringToSign)
	signature, erro := s.Helper.GetSignatureAccessToken(s.ApiConfig.PrivateKey, bytes)

	if erro != nil {
		return &ResponseAccessToken{}, erro
	}

	headers := map[string]string{
		"Content-type": "Application/JSON",
		"X-TIMESTAMP":  formattedDate,
		"X-CLIENT-KEY": s.ApiConfig.ClientID,
		"X-SIGNATURE":  signature,
	}

	body, err := s.HttpClient.Request(headers, url, parameter, `POST`)

	if err != nil {
		fmt.Printf("Error when requesting %s : %s", url, err.Error())
	}

	if len(body) == 0 {
		return &ResponseAccessToken{}, fmt.Errorf("empty response body")
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Failed to parse JSON response:", err)
		return &ResponseAccessToken{}, err
	}

	return &response, nil
}

func (s *Snap) RequestSnapTransaction(parameter map[string]interface{}, endPoint string, accessToken string, httpMethod string) (map[string]interface{}, error) {
	var response map[string]interface{}
	formattedDate := s.Helper.GetFormattedDate()
	url := s.ApiConfig.GetSnapAPIBaseURL() + endPoint
	hexPayload, err := s.Helper.GetEncodePayload(parameter["body"])

	if err != nil {
		fmt.Println("Error when encoding payload")
	}
	stringToSign := fmt.Sprintf("%s:%s:%s:%s:%s", httpMethod, endPoint, accessToken, hexPayload, formattedDate)
	signature := s.Helper.GetRegistSignature(stringToSign, s.ApiConfig.ClientSecret)

	headers := map[string]string{
		"Content-type":  "Application/JSON",
		"X-TIMESTAMP":   formattedDate,
		"X-SIGNATURE":   signature,
		"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		"X-PARTNER-ID":  s.ApiConfig.ClientID,
		"X-EXTERNAL-ID": parameter["headers"].(map[string]string)["X_EXTERNAL_ID"],
		"CHANNEL-ID":    parameter["headers"].(map[string]string)["CHANNEL_ID"],
	}

	body, err := s.HttpClient.Request(headers, url, parameter["body"], httpMethod)
	if err != nil {
		fmt.Printf("Error when requesting %s : %s", url, err.Error())
	}

	if len(body) == 0 {
		return nil, fmt.Errorf("empty response body")
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
