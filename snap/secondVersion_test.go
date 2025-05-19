package snap

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

var secondVersion = APIVersion2{}

func setConf() {

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
		"isCloudServer": false,
		"merchantKey":   "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
	}
	secondVersion.ApiConfig.SetConfiguration(config)

}

func TestRequestV2PaymentRegist(t *testing.T) {
	// Set API Config
	setConf()
	parameterRegisterVA := map[string]interface{}{
		"payMethod":       "02",
		"currency":        "IDR",
		"amt":             "10000",
		"referenceNo":     "ord12320250514090551",
		"goodsNm":         "Test Transaction Nicepay",
		"billingNm":       "John doe",
		"billingPhone":    "081623621623",
		"billingEmail":    "jhondoe@gmail.com",
		"billingAddr":     "Jalan Cempaka Putih",
		"billingCity":     "Jakarta",
		"billingState":    "DKI Jakarta",
		"billingPostCd":   "10540",
		"billingCountry":  "Indonesia",
		"description":     "test cc",
		"deliveryNm":      "John Doe",
		"deliveryPhone":   "0851731575341",
		"deliveryAddr":    "Jalan Cempaka Putih",
		"deliveryCity":    "Jakarta",
		"deliveryState":   "DKI Jakarta",
		"deliveryPostCd":  "10540",
		"deliveryCountry": "Indonesia",
		"dbProcessUrl":    "https://httpdump.app/dumps/fa101255-f007-43f6-9ce2-b581c2b645a3",
		"userIP":          "127.0.0.1",
		"cartData":        "{\"count\":3,\"item\":[{\"goods_id\":30,\"goods_name\":\"Beanie\",\"goods_type\":\"Accessories\",\"goods_amt\":1000,\"goods_sellers_id\":\"NICEPAY-NamaMerchant\",\"goods_sellers_name\":\"NICEPAYSHOP\",\"goods_quantity\":1,\"goods_url\":\"http://www.nicestore.com/product/beanie/\"},{\"goods_id\":31,\"goods_name\":\"Belt\",\"goods_type\":\"Accessories\",\"goods_amt\":5000,\"goods_sellers_id\":\"NICEPAY-NamaMerchant\",\"goods_sellers_name\":\"NICEPAYSHOP\",\"goods_quantity\":1,\"goods_url\":\"http://www.nicestore.store/product/belt/\"},{\"img_url\":\"http://www.jamgora.com/media/avatar/noimage.png\",\"goods_name\":\"Shipping Fee\",\"goods_id\":\"Shipping for Ref. No. 278\",\"goods_detail\":\"Flat rate\",\"goods_type\":\"Shipping with Flat rate\",\"goods_amt\":\"4000\",\"goods_sellers_id\":\"NICEPAY-NamaMerchant\",\"goods_sellers_name\":\"NICEPAYSHOP\",\"goods_quantity\":\"1\",\"goods_url\":\"https://wwww.nicestore.store\"}]}",
		"sellers":         "[{\"sellersId\":\"NICEPAY-NamaMerchant\",\"sellersNm\":\"NICEPAYSHOP\",\"sellersUrl\":\"http://nicestore.store/product/beanie/\",\"sellersEmail\":\"Nicepay@nicepay.co.id\",\"sellersAddress\":{\"sellerNm\":\"NICEPAYSHOP\",\"sellerLastNm\":\"NICEPAYSHOP\",\"sellerAddr\":\"Jln. Kasablanka Kav 88\",\"sellerCity\":\"Jakarta\",\"sellerPostCd\":\"14350\",\"sellerPhone\":\"082111111111\",\"sellerCountry\":\"ID\"}}]",
		"bankCd":          "BRIN",
		"userAgent":       "Mozilla",
		"mitraCd":         "ALMA",
		"instmntMon":      "1",
		"instmntType":     "1",
		"shopId":          "",
	}
	endPoint := "/direct/v2/registration"

	responseTransaction, err := secondVersion.RequestRegisterAPIVersion2(parameterRegisterVA, endPoint)

	if err != nil {
		fmt.Printf("Error requesting %s : %s", endPoint, err.Error())
	}
	responseCode, _ := responseTransaction["resultCd"].(string)
	tXid := responseTransaction["tXid"].(string)

	fmt.Printf("Success registration with TxId = %s | responseCode = %s", tXid, responseCode)
}

func TestRequestV2PaymentCancel(t *testing.T) {
	// Set API Config
	setConf()
	parameterCancelVA := map[string]interface{}{
		"tXid":           "TNICEVA02302202505161604437522",
		"payMethod":      "02",
		"amt":            "10000",
		"cancelType":     "1",
		"cancelMsg":      "Testing Cancel Of Virtual Account",
		"cancelUserId":   "",
		"cancelUserIp":   "127.0.0.1",
		"cancelServerIp": "127.0.0.1",
		"cancelUserInfo": "",
		"cancelRetryCnt": "",
		"worker":         "",
	}
	endPoint := "/direct/v2/cancel"

	responseTransaction, err := secondVersion.RequestRegisterAPIVersion2(parameterCancelVA, endPoint)

	if err != nil {
		fmt.Printf("Error requesting %s : %s", endPoint, err.Error())
	}
	responseCode, _ := responseTransaction["resultCd"].(string)

	fmt.Printf("Success cancel with responseCode = %s", responseCode)
}

func TestRequestV2PaymentInquiry(t *testing.T) {
	// Set API Config
	setConf()
	parameterInquiryVA := map[string]interface{}{
		"tXid":        "IONPAYTEST01202505161557537162",
		"referenceNo": "ord12320250514090551",
		"amt":         "10000",
	}
	endPoint := "/direct/v2/inquiry"

	responseTransaction, err := secondVersion.RequestRegisterAPIVersion2(parameterInquiryVA, endPoint)

	if err != nil {
		fmt.Printf("Error requesting %s : %s", endPoint, err.Error())
	}

	responseCode, _ := responseTransaction["resultCd"].(string)
	tXid := responseTransaction["tXid"].(string)
	status := responseTransaction["status"].(string)

	fmt.Printf("Success inquiry with TxId = %s | responseCode = %s | status= %s", tXid, responseCode, status)
}

func TestRequestV2CheckoutRegist(t *testing.T) {
	// Set API Config
	setConf()
	parameterRegisterVA := map[string]interface{}{
		"payMethod":       "00",
		"currency":        "IDR",
		"amt":             "10000",
		"referenceNo":     "ord12320250514090551",
		"goodsNm":         "Test Transaction Nicepay",
		"billingNm":       "John doe",
		"billingPhone":    "08122163478127",
		"billingEmail":    "johndoe@gmail.com",
		"billingAddr":     "Jalan Cempaka Putih",
		"billingCity":     "Jakarta",
		"billingState":    "DKI Jakarta",
		"billingPostCd":   "10540",
		"billingCountry":  "Indonesia",
		"description":     "test cc",
		"deliveryNm":      "John Doe",
		"deliveryPhone":   "0851731575341",
		"deliveryAddr":    "Jalan Cempaka Putih",
		"deliveryCity":    "Jakarta",
		"deliveryState":   "DKI Jakarta",
		"deliveryPostCd":  "10540",
		"deliveryCountry": "Indonesia",
		"callBackUrl":     "https://httpdump.app/dumps/fa101255-f007-43f6-9ce2-b581c2b645a3",
		"dbProcessUrl":    "https://httpdump.app/dumps/fa101255-f007-43f6-9ce2-b581c2b645a3",
		"userIP":          "127.0.0.1",
		"cartData":        "{\"count\":3,\"item\":[{\"goods_id\":30,\"goods_name\":\"Beanie\",\"goods_type\":\"Accessories\",\"goods_amt\":1000,\"goods_sellers_id\":\"NICEPAY-NamaMerchant\",\"goods_sellers_name\":\"NICEPAYSHOP\",\"goods_quantity\":1,\"goods_url\":\"http://www.nicestore.com/product/beanie/\"},{\"goods_id\":31,\"goods_name\":\"Belt\",\"goods_type\":\"Accessories\",\"goods_amt\":5000,\"goods_sellers_id\":\"NICEPAY-NamaMerchant\",\"goods_sellers_name\":\"NICEPAYSHOP\",\"goods_quantity\":1,\"goods_url\":\"http://www.nicestore.store/product/belt/\"},{\"img_url\":\"http://www.jamgora.com/media/avatar/noimage.png\",\"goods_name\":\"Shipping Fee\",\"goods_id\":\"Shipping for Ref. No. 278\",\"goods_detail\":\"Flat rate\",\"goods_type\":\"Shipping with Flat rate\",\"goods_amt\":\"4000\",\"goods_sellers_id\":\"NICEPAY-NamaMerchant\",\"goods_sellers_name\":\"NICEPAYSHOP\",\"goods_quantity\":\"1\",\"goods_url\":\"https://wwww.nicestore.store\"}]}",
		"sellers":         "[{\"sellersId\":\"NICEPAY-NamaMerchant\",\"sellersNm\":\"NICEPAYSHOP\",\"sellersUrl\":\"http://nicestore.store/product/beanie/\",\"sellersEmail\":\"Nicepay@nicepay.co.id\",\"sellersAddress\":{\"sellerNm\":\"NICEPAYSHOP\",\"sellerLastNm\":\"NICEPAYSHOP\",\"sellerAddr\":\"Jln. Kasablanka Kav 88\",\"sellerCity\":\"Jakarta\",\"sellerPostCd\":\"14350\",\"sellerPhone\":\"082111111111\",\"sellerCountry\":\"ID\"}}]",
		"bankCd":          "BBBA",
		"userAgent":       "Mozilla",
		"mitraCd":         "ALMA",
		"instmntMon":      "1",
		"instmntType":     "1",
		"shopId":          "",
	}
	endPoint := "/redirect/v2/registration"

	responseTransaction, err := secondVersion.RequestRegisterAPIVersion2(parameterRegisterVA, endPoint)

	if err != nil {
		fmt.Printf("Error requesting %s : %s", endPoint, err.Error())
	}

	responseCode, _ := responseTransaction["resultCd"].(string)
	tXid := responseTransaction["tXid"].(string)

	fmt.Printf("Success registration with TxId = %s | responseCode = %s | reffNo = %s\n", tXid, responseCode, reffNo)
}

func TestRequestV2Payment(t *testing.T) {
	// Set API Config
	setConf()
	parameterPaymentCC := map[string]interface{}{
		"tXid":            "IONPAYTEST01202505141739386661",
		"amt":             "10000",
		"referenceNo":     "ord12320250514090551",
		"callBackUrl":     "https://httpdump.app/dumps/fa101255-f007-43f6-9ce2-b581c2b645a3",
		"cardNo":          "4561347179821189",
		"cardExpYymm":     "2607",
		"cardCvv":         "123",
		"recurringToken":  "",
		"cardHolderNm":    "John Doe",
		"preauthToken":    "",
		"cardHolderEmail": "johndoe@gmail.com",
	}

	responseTransaction, err := secondVersion.RequestPaymentAPIVersion2(parameterPaymentCC)

	if err != nil {
		fmt.Printf("Error requesting : %s", err.Error())
	}
	htmlString := string(responseTransaction)
	fmt.Println(htmlString)
}

func TestRequestV2Payout(t *testing.T) {
	// Set API Config
	setConf()
	parameterRegisterPayout := map[string]interface{}{
		"msId":         "",
		"accountNo":    "5345000060",
		"benefNm":      "PT IONPAY NETWORKS",
		"benefStatus":  "1",
		"benefType":    "1",
		"bankCd":       "BDIN",
		"amt":          "10000",
		"referenceNo":  "ORD12345",
		"reservedDt":   "",
		"reservedTm":   "",
		"benefPhone":   "082111111111",
		"description":  "This is test request",
		"payoutMethod": "",
	}
	endPoint := "/api/direct/v2/requestPayout"

	responseTransaction, err := secondVersion.RequestPayoutV2(parameterRegisterPayout, endPoint)

	if err != nil {
		fmt.Printf("Error requesting %s : %s", endPoint, err.Error())
	}

	responseCode, _ := responseTransaction["resultCd"].(string)
	tXid := ""
	if responseTransaction["tXid"] != nil {
		tXid = responseTransaction["tXid"].(string)
	}

	fmt.Printf("Success registration with TxId = %s | responseCode = %s", tXid, responseCode)
}

func TestRequestV2PayoutInquiryBalance(t *testing.T) {
	// Set API Config
	setConf()
	parameterBalance := map[string]interface{}{}
	endPoint := "/api/direct/v2/balanceInquiry"

	responseTransaction, err := secondVersion.RequestPayoutV2(parameterBalance, endPoint)

	if err != nil {
		fmt.Printf("Error requesting %s : %s", endPoint, err.Error())
	}

	responseCode, _ := responseTransaction["resultCd"].(string)
	balance := ""
	if responseTransaction["balance"] != nil {
		balance = responseTransaction["balance"].(string)
	}

	fmt.Printf("Success inquiry balance with balance = %s | responseCode = %s", balance, responseCode)
}

func TestRequestV2PayoutApprove(t *testing.T) {
	// Set API Config
	setConf()
	parameterApprovePayout := map[string]interface{}{
		"tXid": "IONPAYTEST07202505161536386084",
	}
	endPoint := "/api/direct/v2/approvePayout"

	responseTransaction, err := secondVersion.RequestPayoutV2(parameterApprovePayout, endPoint)

	if err != nil {
		fmt.Printf("Error requesting %s : %s", endPoint, err.Error())
	}

	responseCode, _ := responseTransaction["resultCd"].(string)
	tXid := ""
	if responseTransaction["tXid"] != nil {
		tXid = responseTransaction["tXid"].(string)
	}

	fmt.Printf("Success approve with tXid = %s | responseCode = %s", tXid, responseCode)
}

func TestRequestV2PayoutInquiry(t *testing.T) {
	// Set API Config
	setConf()

	parameterInquiryPayout := map[string]interface{}{
		"tXid":      "IONPAYTEST07202505161536386084",
		"accountNo": "5345000060",
	}
	endPoint := "/api/direct/v2/inquiryPayout"

	responseTransaction, err := secondVersion.RequestPayoutV2(parameterInquiryPayout, endPoint)

	if err != nil {
		fmt.Printf("Error requesting %s : %s", endPoint, err.Error())
	}

	responseCode, _ := responseTransaction["resultCd"].(string)
	tXid := ""
	if responseTransaction["tXid"] != nil {
		tXid = responseTransaction["tXid"].(string)
	}

	fmt.Printf("Success inquiry with tXid = %s | responseCode = %s", tXid, responseCode)
}
