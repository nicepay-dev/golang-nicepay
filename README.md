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
    "github.com/nicepay-dev/golang-nicepay"
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

```go
snap := Snap{}
privateKey := key.(*rsa.PrivateKey)

config := map[string]interface{}{
"isProduction": false,
"privateKey":   privateKey,
"clientSecret": "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
"clientId":     "IONPAYTEST",
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
snap := Snap{}
privateKey := key.(*rsa.PrivateKey)

config := map[string]interface{}{
"isProduction": false,
"privateKey":   privateKey,
"clientSecret": "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A==",
"clientId":     "IONPAYTEST",
}
snap.ApiConfig.SetConfiguration(config)

// 2. Initiate Snap request
parameter := map[string]interface{}{
"grantType":      "client_credentials",
"additionalInfo": map[string]interface{}{},
}
reqToken := snap.RequestSnapAccessToken(parameter)

```

## 3. Examples

Integration test are available

- [Snap Sample Functional Test](snap/snap_test.go)

## Get help

- [Nicepay Docs](https://docs.nicepay.co.id/)
- [Nicepay Dashboard ](https://bo.nicepay.co.id/)
- [SNAP documentation](http://snap-docs.Nicepay.com)
- Can't find answer you looking for? email to [cs@nicepay.co.id](mailto:cs@nicepay.co.id)
