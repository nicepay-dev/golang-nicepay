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
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC2QB0PXu3lJszdsTgGIS7PIImQABpqWxnhl9Fde+CTmkDqqKea9LAWcTgpOtr2Gtw5r6ovOwZdl5J8KGsY+XE3QolUeNQTiSy2pjcYl/wL1aSm9n2XkIUqeEtY9Ft2PBFD9DGScTkN3mJyhvt4kkLaDjQiLDXBnBvl+wBJzzNOKpnEdhFxsBnpQo1955QWltBBUtedGB70IyPNZEIB7gVYXZ4/o2CxM7KGjC1Jlqb0uWGSzDSE44pVVSkCtokfw3LDtV0e7D8SlHvVmxXaef/ft6Nhv3xdOkaMiiRDFNC6bdOYiy1tljDrm2MZVV/nV5QCgUgAulTJcIdNQSBaPhZHAgMBAAECggEAWuCyfPEpsCv+WQefN5NBW6BOaCNdCK6/w5moGTUFwaRX/Ys29FJSIga36ftCpxiyuwMo2h9VJ8NGlKm06cGsnlEL2LbdjZZH2RYeACH9WUthrK3Z54N1m71bWRKULut58ogoVe0mdY9wSMqdR7yrID+X6HhiH9Z/pNjaBnQPEcjddaYtas77gaP5WqMwe/I8MlUPY5VEWpFVDx4LG8Uetlh4zdYEEPc70crFovDoKshrRnimABwIGoWAuX9PHYY5Lo29PzN3Tk25APcaaoSLBK2vCZALiQ3tAZizFo+qQLjC5UKzVzLtN/cJ86sktKHnK6TRsYcx/izPhdmj48hlkQKBgQDnJ7rLu4vWjHAPxsl4PNyV/JS56yb8fmFVvUmH+ackUOguQHkwY/hXwEP6OJ7+lo4oL7a+CBvo7RNHQvaq4yK3PzNMMU/YAdQLM/tW8VrFHMQZhc+u8bam3NWce92wWHW3y8zjJKmLnAhsXfHffyXXQt+Dbk9TGXwilE17MCllFwKBgQDJ1sMfAWiA6iI73FzF7w2IlMDsMW335ajNhdTrULmDUr0eQgjIZm4tV9GkPMnefjY03S0H1sLHoJObw1xMKvfYVeD79DqNZbcgk+YKwFEZK0hFAteJclFn2A0DlkvMC+KPqs+jE4QvQCnLLt0mxKp0m6jq1y7F//IXvGa0PKf2UQKBgQCEcig+ufw694byEzW3FjBSJEJXcNyKyiMdTHMIXUyeq1kNv1VxG6bdKMYKZkz7lOppLkWoBt9vDAAC0eSiL7jhhG3xF0QngYysyqEVxP78eCoIcbp5A/hjDZ+7pOF2PIlewYBpGcWnv8S3yvBe3eyhtah6F0eOVsjgy1bF4eemCwKBgQCAG/U678zhzjouXn7wDvw7DZeqEvGmn7lVwbVUKqelB9YLp4QlloYl95CTuxWyR8+mHCBh0llNFcm62vPxUHCBenjT0r97Ue07G0Su3ERdQlCbpOMjVVEAJWyVM0cm2wBRiexCqLeEuigM09EAs2ExpD9B15TTjdGeaTGTAtDlEQKBgQCurHorA6cAttCiiE8Z1CROjSOhv185gvL58tDS/pcMieik0zKyRV8gfnb5Kv8Z6njfE/6zlP5jp2gRAUgS+jCMpDDgd8cRwQoZqdIDxfsaes9SrBCl1XcDS4YO5NXLK3+BhEHlG+/eStPE4emKYpi7tIDS4TdnTIqka0k/l47+gQ==
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
		"isProduction":  false,
		"privateKey":    privateKey,
		"clientSecret":  "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==", // credentials
		"clientId":      "IONPAYTEST",                                                                               // clientId / merchantID
		"isCloudServer": true,
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

func TestRequestSnapTransactionCloud(t *testing.T) {

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

	fmt.Println(response)
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
			"CHANNEL_ID":    "IONPAYTEST", // merchantId
		},
	}
	endPoint := "/api/v1.0/transfer-va/create-va"
	httpMethod := "POST"
	responseTransaction, err := snap.RequestSnapTransaction(parameterRegisterVA, endPoint, accessToken, httpMethod)
	fmt.Println("Finish")
	fmt.Println(responseTransaction)
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
