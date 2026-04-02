package test

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	v1 "github.com/nicepay-dev/golang-nicepay/v1"
)

var firstVersion = v1.APIVersion1{}

func setConfig() {

	block, _ := pem.Decode([]byte(privKeyStr))

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return
	}

	config := map[string]interface{}{
		"isProduction":  false,
		"privateKey":    privateKey,
		"clientSecret":  "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==", // credentials>
		"clientId":      "IONPAYTEST",                                                                               // clientId / merchantID
		"isCloudServer": false,
		"merchantKey":   "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
	}
	firstVersion.ApiConfig.SetConfiguration(config)
}

func TestRequestV1PaymentRegist(t *testing.T) {
	// Set API Config
	setConfig()

	params := url.Values{
		"payMethod":       {"02"},
		"currency":        {"IDR"},
		"amt":             {"10000"},
		"referenceNo":     {"ord12320250514090551"},
		"goodsNm":         {"Test Transaction Nicepay"},
		"billingNm":       {"John doe"},
		"billingPhone":    {"081623621623"},
		"billingEmail":    {"jhondoe@gmail.com"},
		"billingAddr":     {"Jalan Cempaka Putih"},
		"billingCity":     {"Jakarta"},
		"billingState":    {"DKI Jakarta"},
		"billingPostCd":   {"10540"},
		"billingCountry":  {"Indonesia"},
		"description":     {"test cc"},
		"deliveryNm":      {"John Doe"},
		"deliveryPhone":   {"0851731575341"},
		"deliveryAddr":    {"Jalan Cempaka Putih"},
		"deliveryCity":    {"Jakarta"},
		"deliveryState":   {"DKI Jakarta"},
		"deliveryPostCd":  {"10540"},
		"deliveryCountry": {"Indonesia"},
		"dbProcessUrl":    {"https://httpdump.app/dumps/fa101255-f007-43f6-9ce2-b581c2b645a3"},
		"userIP":          {"127.0.0.1"},
		"cartData":        {`{"count":3,"item":[{"goods_id":30,"goods_name":"Beanie","goods_type":"Accessories","goods_amt":1000,"goods_sellers_id":"NICEPAY-NamaMerchant","goods_sellers_name":"NICEPAYSHOP","goods_quantity":1,"goods_url":"http://www.nicestore.com/product/beanie/"},{"goods_id":31,"goods_name":"Belt","goods_type":"Accessories","goods_amt":5000,"goods_sellers_id":"NICEPAY-NamaMerchant","goods_sellers_name":"NICEPAYSHOP","goods_quantity":1,"goods_url":"http://www.nicestore.store/product/belt/"},{"img_url":"http://www.jamgora.com/media/avatar/noimage.png","goods_name":"Shipping Fee","goods_id":"Shipping for Ref. No. 278","goods_detail":"Flat rate","goods_type":"Shipping with Flat rate","goods_amt":"4000","goods_sellers_id":"NICEPAY-NamaMerchant","goods_sellers_name":"NICEPAYSHOP","goods_quantity":"1","goods_url":"https://wwww.nicestore.store"}]}`},
		"sellers":         {`[{"sellersId":"NICEPAY-NamaMerchant","sellersNm":"NICEPAYSHOP","sellersUrl":"http://nicestore.store/product/beanie/","sellersEmail":"Nicepay@nicepay.co.id","sellersAddress":{"sellerNm":"NICEPAYSHOP","sellerLastNm":"NICEPAYSHOP","sellerAddr":"Jln. Kasablanka Kav 88","sellerCity":"Jakarta","sellerPostCd":"14350","sellerPhone":"082111111111","sellerCountry":"ID"}}]`},
		"bankCd":          {"BRIN"},
		"userAgent":       {"Mozilla"},
		"mitraCd":         {"ALMA"},
		"instmntMon":      {"1"},
		"instmntType":     {"1"},
		"shopId":          {""},
	}

	endPoint := "https://dev.nicepay.co.id/nicepay/api/onePass.do"

	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", endPoint, body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 but got %d", resp.StatusCode)
	}

}
