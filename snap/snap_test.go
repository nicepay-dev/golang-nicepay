package snap

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var privKeyStr string = `-----BEGIN PRIVATE KEY-----

-----END PRIVATE KEY-----`

var snap = Snap{}

var txId string
var reffNo string
var amount = "10000.00"
var vaNo string

func SetApiConfig() {

	block, _ := pem.Decode([]byte(privKeyStr))

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return
	}

	config := map[string]interface{}{
		"isProduction": false,
		"privateKey":   privateKey,
		"clientSecret": "", // credentials
		"clientId":     "", // clientId / merchantID
	}
	snap.ApiConfig.SetConfiguration(config)

}

func GetAccessToken() (*ResponseAccessToken, error) {

	parameter := map[string]interface{}{
		"grantType":      "client_credentials",
		"additionalInfo": map[string]interface{}{},
	}

	return snap.RequestSnapAccessToken(parameter)

}

func TestRequestSnapAccessToken(t *testing.T) {
	// Create an instance of the Snap struct

	//snap := Snap{}

	// Set API Config
	SetApiConfig()

	// Get Access token
	response, err := GetAccessToken()

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

	// Set API Config
	SetApiConfig()

	// Get Access Token
	response, err := GetAccessToken()

	if err != nil {
		fmt.Printf("Error request snap access token : %s", err.Error())
	}
	accessToken := response.AccessToken

	reffNo = "order" + strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	parameterRegisterVA := map[string]interface{}{
		"body": map[string]interface{}{
			"partnerServiceId":   "",
			"customerNo":         "",
			"virtualAccountNo":   "",
			"virtualAccountName": "John Test",
			"trxId":              reffNo,
			"totalAmount": map[string]string{
				"value":    amount,
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
			"CHANNEL_ID":    "", // merchantId
		},
	}
	endPoint := "/api/v1.0/transfer-va/create-va"
	httpMethod := "POST"
	responseTransaction, err := snap.RequestSnapTransaction(parameterRegisterVA, endPoint, accessToken, httpMethod)
	if err != nil {
		fmt.Printf("Error requesting %s : %s", endPoint, err.Error())
	}

	responseCode, _ := responseTransaction["responseCode"].(string)
	virtualAccountNo, ok := responseTransaction["virtualAccountData"].(map[string]interface{})["virtualAccountNo"].(string)

	vaNo = virtualAccountNo
	txId = responseTransaction["virtualAccountData"].(map[string]interface{})["additionalInfo"].(map[string]interface{})["tXidVA"].(string)

	if !ok {
		fmt.Println("Error Virtual Account not found")
	}

	if responseCode != "2002700" {
		t.Errorf("Unexpected response. Got %v, want %v", response, "2002700")
	}

	fmt.Printf("Success VA registration with TxId = %s | VaNo = %s | reffNo = %s\n", txId, vaNo, reffNo)

}

func TestCancelVASnapTransaction(t *testing.T) {

	// Create VA
	TestRequestSnapTransaction(t)

	// START TEST CANCEL VA
	externalId2 := fmt.Sprintf("%06d", rand.Intn(1000000))
	// Set API Config
	SetApiConfig()

	// Get Access Token
	response, err := GetAccessToken()

	if err != nil {
		fmt.Printf("Error request snap access token : %s", err.Error())
	}
	accessToken := response.AccessToken

	parameterDelete := map[string]interface{}{
		"body": map[string]interface{}{
			"partnerServiceId": "",
			"customerNo":       "",
			"virtualAccountNo": vaNo,
			"trxId":            reffNo,
			"additionalInfo": map[string]interface{}{
				"totalAmount": map[string]string{
					"value":    amount,
					"currency": "IDR",
				},
				"tXidVA":        txId,
				"cancelMessage": "Cancel Virtual Account",
			},
		},
		"headers": map[string]string{
			"X_EXTERNAL_ID": externalId2,
			"CHANNEL_ID":    "", //merchantId
		},
	}

	endPoint := "/api/v1.0/transfer-va/delete-va"
	httpMethod := "DELETE"

	responseTransaction, err := snap.RequestSnapTransaction(parameterDelete, endPoint, accessToken, httpMethod)

	if err != nil {
		fmt.Printf("Error requesting %s : %s", endPoint, err.Error())
	}

	responseCode, _ := responseTransaction["responseCode"].(string)

	if responseCode != "2003100" {
		t.Errorf("Unexpected response. Got %v, want %v", responseCode, "2003100")
	}
	fmt.Printf("Success cancel VA with TxId = %s | VaNo = %s | reffNo = %s\n", txId, vaNo, reffNo)

}

func TestVerifySHA256RSA(t *testing.T) {

	// let snap = new Snap();

	signatureString := "VoxMPjbcV9pro4YyHGQgoRj4rDVJgYk2Ecxn+95B90w47Wnabtco35BfhGpR7a5RukUNnAdeOEBNczSFk4B9uYyu3jc+ceX+Dvz5OYSgSnw5CiMHtGiVnTAqCM/yHZ2MRpIEqekBc4BWMLVtexSWp0YEJjLyo9dZPrSkSbyLVuD7jkUbvmEpVdvK0uK15xb8jueCcDA6LYVXHkq/OMggS1/5mrLNriBhCGLuR7M7hBUJbhpOXSJJEy7XyfItTBA+3MRC2FLcvUpMDrn/wz1uH1+b9A6FP7mG0bRSBOm2BTLyf+xJR5+cdd88RhF70tNQdQxhqr4okVo3IFqlCz2FFg=="
	dataString := "TNICEVA023|2024-08-19T17:12:40+07:00"
	publicKeyString := "" // String public key

	isVerified, error := snap.Helper.VerifySHA256RSA(dataString, publicKeyString, signatureString)

	fmt.Printf("Is the signature valid? %t \n", isVerified)
	if error != nil {
		t.Errorf("error verifying data : %s", error.Error())
	}

	if isVerified != true {
		t.Errorf("Data is not valid")
	} else {
		fmt.Println("Data is valid")
	}

}
