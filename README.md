# Nicepay Go Library

Nicepay :heart: Go !

Go is a highly modern programming language that combines aspects of dynamic and static typing, making it particularly well-suited for web development and various other applications. It has a concise syntax and offers the advantages of both dynamic and static typing. Furthermore, Go has a small memory footprint, which is a notable benefit. The purpose of this module is to assist you in utilizing Nicepay product's REST APIs effectively within the Go programming language.

## 1. Installation

### 1.1 Using Go Module

Run this command on your project to initialize Go mod (if you haven't):

```go
go mod init
```

then reference Nicepay-go in your project file with `import`:

```go
import (
    "github.com/nicepay-dev/golang-nicepay/snap"
)
```

### 1.2 Using go get

Also, the alternative way you can use `go get` the package into your project

```go
go get -u github.com/nicepay-dev/golang-nicepay
```

## 2. Usage

There is a type named `Client` (`snap.Client`) that should be instantiated through
function `New` which holds any possible setting to the library. Any activity (charge, approve, etc) is done in the client level.

### 2.1 Choose Product/Method

We have [3 different products](https://beta-docs.nicepay.com/) that you can use:

- [Snap](#22A-snap) - Customizable payment popup will appear on **your web/app** (no redirection). [doc ref](https://snap-docs.Nicepay.com/)
- [Snap Redirect](#22B-snap-redirect) - Customer need to be redirected to payment url **hosted by Nicepay**. [doc ref](https://snap-docs.Nicepay.com/)

To learn more and understand each of the product's quick overview you can visit https://docs.Nicepay.com.

### 2.2 Client Initialization and Configuration

Get your client key and server key from [Nicepay Dashboard](https://dashboard.Nicepay.com)

Create API client object, You can also check the [project's implementation](example/simple) for more examples. Please proceed there for more detail on how to run the example.

#### 2.2.1 Using global config

Set a config with globally, (except for iris api)

> **WARNING:** Credentials used here are for testing purposes only.

```go
snap := snap.Snap{}
privateKey := key.(*rsa.PrivateKey)

config := map[string]interface{}{
"isProduction": false,
"privateKey":   privateKey,
"clientSecret": "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
"clientId":     "IONPAYTEST",
"isCloudServer": false,
"merchantKey":   "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
}
snap.ApiConfig.SetConfiguration(config)
```

### 2.3 Snap

Snap is Nicepay existing tool to help merchant charge customers using a mobile-friendly, in-page,
no-redirect checkout facilities. [Using snap is simple](https://docs.Nicepay.com/en/snap/overview).

Available methods for Snap

```go
// CreateTransaction : Do `/transactions` API request to SNAP API to get Snap token and redirect url with `snap.Request`
func RequestSnapTransaction(parameter map[string]interface{}) ([]byte, error)

// CreateTransactionToken : Do `/transactions` API request to SNAP API to get Snap token with `snap.Request`
func RequestSnapAccessToken(parameter map[string]interface{}) ([]byte, error)

```

Snap usage example, create transaction with minimum Snap parameters (choose **one** of alternatives below):

#### 2.3.1 Using global Config & static function

Sample usage if you prefer Nicepay global configuration & using static function. Useful if you only use 1 merchant account API key, and keep the code short.

```go
// 1. Set you ServerKey with globally
snap := snap.Snap{}
privateKey := key.(*rsa.PrivateKey)

config := map[string]interface{}{
"isProduction": false,
"privateKey":   privateKey,
"clientSecret": "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
"clientId":     "IONPAYTEST",
"isCloudServer": false,
"merchantKey":   "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
}
snap.ApiConfig.SetConfiguration(config)

// 2. Initiate Snap request
parameter := map[string]interface{}{
"grantType":      "client_credentials",
"additionalInfo": map[string]interface{}{},
}
reqToken := snap.RequestSnapAccessToken(parameter)

```

#### 2.3.2 Verify Signature Notif

```go
signatureString := "YOUR_SIGNATURE_IN_STRING"// Ex : "VoxMPjbcV9pro4YyHGQgoRj4rDVJgYk2Ecxn+95B90w47Wnabtco35BfhGpR7a5RukUNnAdeOEBNczSFk4B9uYyu3jc+ceX+Dvz5OYSgSnw5CiMHtGiVnTAqCM/yHZ2MRpIEqekBc4BWMLVtexSWp0YEJjLyo9dZPrSkSbyLVuD7jkUbvmEpVdvK0uK15xb8jueCcDA6LYVXHkq/OMggS1/5mrLNriBhCGLuR7M7hBUJbhpOXSJJEy7XyfItTBA+3MRC2FLcvUpMDrn/wz1uH1+b9A6FP7mG0bRSBOm2BTLyf+xJR5+cdd88RhF70tNQdQxhqr4okVo3IFqlCz2FFg=="
dataString := "YOUR_DATA_TO_SIGN" // Ex :"TNICEVA023|2024-08-19T17:12:40+07:00"
publicKeyString := "PUBLIC_KEY_PEM" // Ex : "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApizrKJl/1Legp3Zj8f0oTIjKnUWe2HJCBSoRsVLxtpf0Dr1MI+23y+AMNKKxVXxbvReZq/sD91uN4GFYMUr16LY9oX7nJXh9C1JlI4/Xb/Q9MF30o1XYvogHLATtvTR/KQ8hxrf6Nlj/yuzeqrT+PiQMZt1CaKiE6UMn36kq11DmDq4ocwcNhChKDudNZSZ4YYIFn5IgH05K+VsRjehpa0szbO8qHmvnprXVVcqvk7ZSS+6fYwDynOq0f552aL0LWX0glNhh9F0oJqmTreW4lM0mdhNDq4GhlJZl5IpaUiaGRM2Rz/t6spgwR7nqUhI9aE2kjzaorgP4ZWUGm3wlTwIDAQAB"

isVerified, error := snap.Helper.VerifySHA256RSA(dataString, publicKeyString, signatureString)
```

#### 2.4.1 Request Register API Payment/Direct

```go
    secondVersion := APIVersion2{}
    privateKey := key.(*rsa.PrivateKey)
    config := map[string]interface{}{
    		"isProduction":  false,
    		"privateKey":    privateKey,
    		"clientSecret":  "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
    		"clientId":      "IONPAYTEST",
    		"isCloudServer": false,
    		"merchantKey":   "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
    	}
    secondVersion.ApiConfig.SetConfiguration(config)
    reffNo = "order" + strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	parameterRegisterVA := map[string]interface{}{
		"payMethod":       "01",
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
		"bankCd":          "BBBA",
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
```

#### 2.4.2 Request Register API Checkout/Redirect

```go
 	secondVersion := APIVersion2{}
    privateKey := key.(*rsa.PrivateKey)
    config := map[string]interface{}{
    		"isProduction":  false,
    		"privateKey":    privateKey,
    		"clientSecret":  "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
    		"clientId":      "IONPAYTEST",
    		"isCloudServer": false,
    		"merchantKey":   "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
    	}
    secondVersion.ApiConfig.SetConfiguration(config)

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

	fmt.Printf("Success registration with TxId = %s | responseCode = %s", tXid, responseCode)
```

#### 2.4.3 Request Payment API Payment/Direct

```go
	secondVersion := APIVersion2{}
    privateKey := key.(*rsa.PrivateKey)
    config := map[string]interface{}{
    		"isProduction":  false,
    		"privateKey":    privateKey,
    		"clientSecret":  "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
    		"clientId":      "IONPAYTEST",
    		"isCloudServer": false,
    		"merchantKey":   "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
    	}
    secondVersion.ApiConfig.SetConfiguration(config)
	
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
```


#### 2.5.1 Request Register API Payout

```go
	secondVersion := APIVersion2{}
    privateKey := key.(*rsa.PrivateKey)
    config := map[string]interface{}{
    		"isProduction":  false,
    		"privateKey":    privateKey,
    		"clientSecret":  "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
    		"clientId":      "IONPAYTEST",
    		"isCloudServer": false,
    		"merchantKey":   "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
    	}
    secondVersion.ApiConfig.SetConfiguration(config)
	
	parameterRegistPayout := map[string]interface{}{
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

	responseTransaction, err := secondVersion.RequestPayoutV2(parameterRegistPayout)

	responseCode, _ := responseTransaction["resultCd"].(string)
	tXid := ""
	if responseTransaction["tXid"] != nil {
		tXid = responseTransaction["tXid"].(string)
	}

	fmt.Printf("Success registration with TxId = %s | responseCode = %s", tXid, responseCode)
```

## 3. Examples

Integration test are available

- [Snap Sample Functional Test](snap/snap_test.go)

## Get help

- [Nicepay Docs](https://docs.nicepay.co.id/)
- [Nicepay Dashboard ](https://bo.nicepay.co.id/)
- [SNAP documentation](http://snap-docs.Nicepay.com)
- Can't find answer you looking for? email to [cs@nicepay.co.id](mailto:cs@nicepay.co.id)
