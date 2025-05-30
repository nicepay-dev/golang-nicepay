package snap

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type APIVersion2 struct {
	ApiConfig  Config
	HttpClient HttpRequest
	Helper     Helper
}

// Do transaction API request to API Direct/Redirect V2 (Register, Inquiry, Cancel)
// @param  {interface} parameter - interface of request body API Register/Checkout/Inquiry/Cancel Nicepay.
// @param  {string} parameter - end point url want to used.
// @return {interface} - response body from request API Nicepay.
func (av *APIVersion2) RequestRegisterAPIVersion2(parameter map[string]interface{}, endPoint string) (map[string]interface{}, error) {
	var response map[string]interface{}

	timeStamp := av.Helper.GetTimestampFormat()
	amt, ok := parameter["amt"].(string)
	if !ok {
		return nil, fmt.Errorf("amt is missing or not a string")
	}
	var referenceNo string
	if parameter["referenceNo"] != nil {
		referenceNo = parameter["referenceNo"].(string)
	}
	var tXid string
	if parameter["tXid"] != nil {
		tXid = parameter["tXid"].(string)
	}

	var stringToSign string
	if endPoint == "/direct/v2/cancel" {
		stringToSign = fmt.Sprintf("%s%s%s%s%s", timeStamp, av.ApiConfig.ClientID, tXid, amt, av.ApiConfig.MerchantKey)
	} else {
		stringToSign = fmt.Sprintf("%s%s%s%s%s", timeStamp, av.ApiConfig.ClientID, referenceNo, amt, av.ApiConfig.MerchantKey)
	}		
	
	merchantToken := av.Helper.SHA256Encrypt(stringToSign)

	requestBody := make(map[string]interface{})
	for k, v := range parameter {
		requestBody[k] = v
	}

	requestBody["merchantToken"] = merchantToken
	requestBody["timeStamp"] = timeStamp
	requestBody["iMid"] = av.ApiConfig.ClientID

	headers := map[string]string{
		"Content-type": "Application/JSON",
	}

	url := fmt.Sprintf("%s%s", av.ApiConfig.GetSnapAPIBaseURL(), endPoint)
	body, err := av.HttpClient.Request(headers, url, requestBody, `POST`)

	if err != nil {
		fmt.Printf("Error when requesting %s : %s", url, err.Error())
	}

	if len(body) == 0 {
		return response, fmt.Errorf("empty response body")
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Failed to parse JSON response:", err)
		return response, err
	}

	return response, nil
}

func (av *APIVersion2) RequestPaymentAPIVersion2(parameter map[string]interface{}) ([]byte, error) {
	headers := map[string]string{
		"Content-type": "application/x-www-form-urlencoded",
	}

	timeStamp := av.Helper.GetTimestampFormat()

	amt, ok := parameter["amt"]
	if !ok {
		return nil, fmt.Errorf("amt is missing")
	}
	referenceNo, ok := parameter["referenceNo"]
	if !ok {
		return nil, fmt.Errorf("referenceNo is missing")
	}

	stringToSign := fmt.Sprintf("%s%s%s%s%s", timeStamp, av.ApiConfig.ClientID, referenceNo, amt, av.ApiConfig.MerchantKey)
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

	url := fmt.Sprintf("%s/direct/v2/payment", av.ApiConfig.GetSnapAPIBaseURL())

	body, err := av.HttpClient.RequestPayment(headers, url, []byte(formData.Encode()), "POST")
	if err != nil {
		return nil, fmt.Errorf("error when requesting %s: %w", url, err)
	}

	return body, nil
}


// Do transaction API request to API Payout V2
// @param  {interface} parameter - interface of request body API Register/Checkout/Inquiry/Cancel Nicepay.
// @param  {string} parameter - end point url want to used.
// @return {interface} - response body from request API Nicepay.
func (av *APIVersion2) RequestPayoutV2(parameter map[string]interface{}, endPoint string) (map[string]interface{}, error) {
	var response map[string]interface{}

	timeStamp := av.Helper.GetTimestampFormat()
	var amt string
	if parameter["amt"] != nil {
		amt = parameter["amt"].(string)
	}
	var accountNo string
	if parameter["accountNo"] != nil {
		accountNo = parameter["accountNo"].(string)
	}
	var tXid string
	if parameter["tXid"] != nil {
		tXid = parameter["tXid"].(string)
	}

	var stringToSign string
	if endPoint == "/api/direct/v2/requestPayout" {
		stringToSign = fmt.Sprintf("%s%s%s%s%s", timeStamp, av.ApiConfig.ClientID, amt, accountNo, av.ApiConfig.MerchantKey)
	} else if endPoint == "/api/direct/v2/balanceInquiry" {
		stringToSign = fmt.Sprintf("%s%s%s", timeStamp, av.ApiConfig.ClientID, av.ApiConfig.MerchantKey)
	} else if endPoint == "/api/direct/v2/inquiryPayout" {
		stringToSign = fmt.Sprintf("%s%s%s%s%s", timeStamp, av.ApiConfig.ClientID, tXid, accountNo, av.ApiConfig.MerchantKey)
	} else {
		stringToSign = fmt.Sprintf("%s%s%s%s", timeStamp, av.ApiConfig.ClientID, tXid, av.ApiConfig.MerchantKey)
	}
	merchantToken := av.Helper.SHA256Encrypt(stringToSign)

	requestBody := make(map[string]interface{})
	for k, v := range parameter {
		requestBody[k] = v
	}
	requestBody["merchantToken"] = merchantToken

	fmt.Println(merchantToken)
	requestBody["timeStamp"] = timeStamp
	requestBody["iMid"] = av.ApiConfig.ClientID

	headers := map[string]string{
		"Content-type": "Application/JSON",
	}

	url := fmt.Sprintf("%s%s", av.ApiConfig.GetSnapAPIBaseURL(), endPoint)
	body, err := av.HttpClient.Request(headers, url, requestBody, `POST`)

	if err != nil {
		fmt.Printf("Error when requesting %s : %s", url, err.Error())
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Failed o parse JSON response:", err)
		return response, err
	}

	return response, nil
}
