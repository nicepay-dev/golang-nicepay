package snap

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var privKeyStr string = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQBzElwG7JVPdSvPu4xDIKuh0Dw3KkporGM0m9TuPMYSHm8/S4Ko
dPSv2+VJ0pCoxngIRdRdxyR1ps4T+up+GWRZHy0qRXvtMEiAJuYAmKvxkqaiYb7r
TFl2qQRA8gfBU97FfQM1YerclPp6B9JD5bRL5Dm7e+bXIeLdtPYGG3B99hru/6vn
Iri5yPa2EVmdCIkUD/wP6mY8W9nRIyp7X0ChjebnMsGtcRmMfGsbLKWyPV1PeeZw
dxMNnr27XdBcxsYeDcrfe+lytDgyJzj+pRNPbWMv/KD00yg3+4ZQQdK/rxy+Uw3/
WN7Fw8mSbpob6vv3Oz0bHTCVFhWS8mp5dDvPAgMBAAECggEAToW7uJnedV8ea12u
o+v6UqwXOwmnxu/DrpWb4ookGx8biNSNL0jH4+0o9Iw0XIc6R2LnPKr0zTfrLiUt
uKi5Gju1BUBvBXbKMnDYyJVl163b+bi7oDL0ZY2GMo82DY2e2aKp+tZ7ftRGa9lE
eUKZGqR9ZNtytWERP3sJ2zcEN187ZvASFABrDNAFwUKQEmNjPeXEyFLJbY2VVXHB
RvlC8Gf1gu4BjaKKwTrfgo+gYBkBoEeV5wvtYngKrWZmgNBKKOqcMqKGcER3giCv
NQ7zpUc4wsvvkrjFPJoLTXPi/AYZ1unbxEjUehEkcWTAAWekPI2waVRKnIduve1b
m3bKAQKBgQC2lo68fR97Al31SzicxCfA2GdRtrnAdBfyc+nw57p59jNKs2piSBAg
YyvcUBZJG1FxTY62OQR2Xn7h98hXgfJ3fXQTk9uhhjbiKFlmMBkGlq1k5DEP9PLv
HZ5HswNEPFNSOupjkrOMsoQGDInQr1VW+6AR3YNfoAxrTQEWPx7DAQKBgQChVngh
c1RAkN6eGn3faUjL1+8yhPj5Q8NRUoBSw/ujZ0Her5S2i8h6sx6r2n8hMHEIUZJ+
ZvvizLv2Ijl/wF9XTehJxnG62P4uhP3WcAOE6kimt/tiwRPUHfgPnT5HiASI/fJY
OLMIKCgdCa3kuQwPoH3ZTT9NfIK79aEKoWqOzwKBgAptmONdBhJBdVpQHICfl2Gl
OmlpVTyPpNp9Ekxm/7h9fjpy+s14LiubXmLr1AoC3GjrNA5mPUIBbZ+8Rh3xVwbK
DHodxLp57uKFyW1Tq+o7atXLTp4JsGJFv8d6iuI3y85zfPWI6GZNv8qUpr5bdTVN
k7vRefJZMrxiHoDFxB0BAoGAPtvvxiinBNjsw3DS5f6hTDp/iZFhZ8zNBpw8PwL4
wftzII4MROtFWvj61D43FflHsNQHXZRGQ2E9QnKnMG0FOIC0JjpZCVGOBxXtyGSw
GlMlpz87hIhxb02V3o+HOlt2WOGIUHMW3fC3YEjrJZgraNNA9S8xoMEINq9G5Vtq
puUCgYEAooHr6y4XkbSaOYc+0b1WGRR/1nAv2VaYB3qkAte8X+oyzdBhjRan4u4c
IJEWjZv/7W/Nsf7PaDSB94e5TIKDvxEnXhsIXIqWDwUI7Op4tcFuZcB46+CVAiWJ
V6J483Lm/wHi/e27LKQF5hc30/ZED5HRcoy/Gr7yNB2lVXjDZpI=
-----END RSA PRIVATE KEY-----`

var snap = Snap{}

func TestRequestSnapAccessToken(t *testing.T) {
	// Create an instance of the Snap struct

	//snap := Snap{}
	block, _ := pem.Decode([]byte(privKeyStr))

	if block == nil || block.Type != "RSA PRIVATE KEY" {
		fmt.Println("Failed to parse RSA private key")
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return
	}

	config := map[string]interface{}{
		"isProduction": false,
		"privateKey":   privateKey,
		"clientSecret": "1af9014925cab04606b2e77a7536cb0d5c51353924a966e503953e010234108a",
		"clientId":     "TNICEVA023",
	}
	snap.ApiConfig.SetConfiguration(config)

	parameter := map[string]interface{}{
		"grantType":      "client_credentials",
		"additionalInfo": map[string]interface{}{},
	}

	response, err := snap.RequestSnapAccessToken(parameter)

	if err != nil {
		t.Errorf("RequestSnapAccessToken failed: %v", err)
	}
	if response.ResponseCode != "2007300" {
		t.Errorf("Unexpected response. Got %v, want %v", response, "2007300")
	}

}

func TestRequestSnapTransaction(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	externalId2 := fmt.Sprintf("%06d", rand.Intn(1000000))

	privateKey, err := snap.Helper.ParsePrivateKey(privKeyStr)

	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return
	}

	config := map[string]interface{}{
		"isProduction": false,
		"privateKey":   privateKey,
		"clientSecret": "1af9014925cab04606b2e77a7536cb0d5c51353924a966e503953e010234108a",
		"clientId":     "TNICEVA023",
	}
	snap.ApiConfig.SetConfiguration(config)

	parameterAccessToken := map[string]interface{}{
		"grantType":      "client_credentials",
		"additionalInfo": map[string]interface{}{},
	}

	response, err := snap.RequestSnapAccessToken(parameterAccessToken)

	accessToken := response.AccessToken

	parameterRegisterVA := map[string]interface{}{
		"body": map[string]interface{}{
			"partnerServiceId":   "",
			"customerNo":         "",
			"virtualAccountNo":   "",
			"virtualAccountName": "John Test",
			"trxId":              "2020102900000000000001",
			"totalAmount": map[string]string{
				"value":    "10000.00",
				"currency": "IDR",
			},
			"additionalInfo": map[string]string{
				"bankCd":       "BMRI",
				"goodsNm":      "Testing",
				"dbProcessUrl": " https://google.com",
				"vacctValidDt": "",
				"vacctValidTm": "",
				"msId":         "",
				"msFee":        "",
				"msFeeType":    "",
				"mbFee":        "",
				"mbFeeType":    "",
			},
		},
		"headers": map[string]string{
			"X_EXTERNAL_ID": externalId2,
			"CHANNEL_ID":    "TNICEVA02301",
		},
	}
	endPoint := "/api/v1.0/transfer-va/create-va"
	responseTransaction, err := snap.RequestSnapTransaction(parameterRegisterVA, endPoint, accessToken)
	responseCode, _ := responseTransaction["responseCode"].(string)
	virtualAccountNo, ok := responseTransaction["virtualAccountData"].(map[string]interface{})["virtualAccountNo"].(string)
	if !ok {
		fmt.Println("Error Virtual Account not found")
	}

	fmt.Println(virtualAccountNo)

	if responseCode != "2002700" {
		t.Errorf("Unexpected response. Got %v, want %v", response, "2002700")
	}

}
