package library

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"time"
)

type Helper struct {
}

func (h *Helper) GetSignatureAccessToken(privateKey *rsa.PrivateKey, stringToSign []byte) (string, error) {
	hashed := sha256.Sum256(stringToSign)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}

	encodedSignature := base64.StdEncoding.EncodeToString(signature)

	return encodedSignature, nil
}

func (h *Helper) GetEncodePayload(requestBody interface{}) (string, error) {
	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	hashed := sha256.Sum256(jsonString)
	return fmt.Sprintf("%x", hashed), nil
}

func (h *Helper) GetRegistSignature(stringToSign string, clientSecret string) string {
	hmac := hmac.New(sha512.New, []byte(clientSecret))
	hmac.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(hmac.Sum(nil))
}

func (h *Helper) GetFormattedDate() string {
	now := time.Now()
	formattedDate := now.Format("2006-01-02T15:04:05-07:00")
	return formattedDate
}

func (h *Helper) GetConvertFormattedDate(timeStamp string) (string, error) {
	t, err := time.Parse("20060102150405", timeStamp)
	if err != nil {
		return "", err
	}
	formattedDate := t.Format("2006-01-02T15:04:05-07:00")
	return formattedDate, nil
}

func (h *Helper) ParsePrivateKey(privateKeyString string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyString))

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
