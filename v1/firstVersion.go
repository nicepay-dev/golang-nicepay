package v1

import (
	"encoding/json"
	"fmt"
	"net/url"

	utils "github.com/nicepay-dev/golang-nicepay/utils"
)

type APIVersion1 struct {
	ApiConfig  utils.Config
	HttpClient utils.HttpRequest
	Helper     utils.Helper
}

/*
|--------------------------------------------------------------------------
| Helper: Convert map → application/x-www-form-urlencoded
|--------------------------------------------------------------------------
*/
func mapToFormData(data map[string]interface{}) []byte {
	form := url.Values{}
	for k, v := range data {
		form.Set(k, fmt.Sprintf("%v", v))
	}
	return []byte(form.Encode())
}

/*
|--------------------------------------------------------------------------
| Register / Inquiry / Cancel (API V1)
|--------------------------------------------------------------------------
*/
func (av *APIVersion1) RequestRegisterAPIVersion1(
	parameter map[string]interface{},
	endPoint string,
) (map[string]interface{}, error) {

	var response map[string]interface{}

	timeStamp := av.Helper.GetTimestampFormat()

	amt, ok := parameter["amt"].(string)
	if !ok {
		return nil, fmt.Errorf("amt is missing or not a string")
	}

	var referenceNo, tXid string
	if v, ok := parameter["referenceNo"].(string); ok {
		referenceNo = v
	}
	if v, ok := parameter["tXid"].(string); ok {
		tXid = v
	}

	var stringToSign string
	if endPoint == "/nicepay/api/onePassAllCancel.do" {
		stringToSign = fmt.Sprintf(
			"%s%s%s%s%s",
			timeStamp,
			av.ApiConfig.ClientID,
			tXid,
			amt,
			av.ApiConfig.MerchantKey,
		)
	} else {
		stringToSign = fmt.Sprintf(
			"%s%s%s%s%s",
			timeStamp,
			av.ApiConfig.ClientID,
			referenceNo,
			amt,
			av.ApiConfig.MerchantKey,
		)
	}

	merchantToken := av.Helper.SHA256Encrypt(stringToSign)

	requestBody := map[string]interface{}{}
	for k, v := range parameter {
		requestBody[k] = v
	}

	requestBody["merchantToken"] = merchantToken
	requestBody["timeStamp"] = timeStamp
	requestBody["iMid"] = av.ApiConfig.ClientID

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	requestURL := fmt.Sprintf(
		"%s%s",
		av.ApiConfig.GetSnapAPIBaseURL(),
		endPoint,
	)

	body, err := av.HttpClient.RequestPayment(
		headers,
		requestURL,
		mapToFormData(requestBody),
		"POST",
	)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return nil, fmt.Errorf("empty response body")
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

/*
|--------------------------------------------------------------------------
| OnePass Payment (API V1)
|--------------------------------------------------------------------------
*/
func (av *APIVersion1) RequestPaymentAPIVersion1(
	parameter map[string]interface{},
) ([]byte, error) {

	timeStamp := av.Helper.GetTimestampFormat()

	amt, ok := parameter["amt"]
	if !ok {
		return nil, fmt.Errorf("amt is missing")
	}

	referenceNo, ok := parameter["referenceNo"]
	if !ok {
		return nil, fmt.Errorf("referenceNo is missing")
	}

	stringToSign := fmt.Sprintf(
		"%s%s%s%s%s",
		timeStamp,
		av.ApiConfig.ClientID,
		referenceNo,
		amt,
		av.ApiConfig.MerchantKey,
	)

	merchantToken := av.Helper.SHA256Encrypt(stringToSign)

	formData := url.Values{}
	formData.Set("timeStamp", timeStamp)
	formData.Set("merchantToken", merchantToken)

	optionalParams := []string{
		"tXid",
		"callBackUrl",
		"cardNo",
		"cardExpYymm",
		"cardCvv",
		"recurringToken",
		"cardHolderNm",
		"preauthToken",
		"cardHolderEmail",
	}

	for _, param := range optionalParams {
		if value, ok := parameter[param].(string); ok {
			formData.Set(param, value)
		}
	}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	requestURL := fmt.Sprintf(
		"%s/nicepay/api/onePass.do",
		av.ApiConfig.GetSnapAPIBaseURL(),
	)

	return av.HttpClient.RequestPayment(
		headers,
		requestURL,
		[]byte(formData.Encode()),
		"POST",
	)
}
