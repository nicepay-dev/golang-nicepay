package snap

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
	"errors"
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

func (h *Helper) ParsePrivateKey(privateKeyString string) (any, error) {
	block, _ := pem.Decode([]byte(privateKeyString))

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func (h *Helper) VerifySHA256RSA(stringToSign string, publicKeyString string, signatureString string) (bool, error) {

	// Format the public key in PEM format
	pemKey := fmt.Sprintf(`-----BEGIN PUBLIC KEY-----
%s
-----END PUBLIC KEY-----`, publicKeyString)

	// Decode the PEM block
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		return false, errors.New("failed to decode PEM block containing public key")
	}

	// Parse the public key
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, fmt.Errorf("failed to parse public key: %v", err)
	}

	// Assert that the parsed key is an *rsa.PublicKey
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return false, errors.New("not an RSA public key")
	}

	// Decode the base64 signature
	signature, err := base64.StdEncoding.DecodeString(signatureString)
	if err != nil {
		return false, fmt.Errorf("failed to decode base64 signature: %v", err)
	}

	// Create a SHA-256 hash of the stringToSign
	hash := sha256.New()
	hash.Write([]byte(stringToSign))
	digest := hash.Sum(nil)

	// Verify the signature using RSA-PSS
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, digest, signature)
	if err != nil {
		return false, fmt.Errorf("verification failed: %v", err)
	}

	// If no error, the signature is verified
	return true, nil
}

func (h *Helper) SHA256Encrypt(stringToSign string) string {

	hasher := sha256.New()
	hasher.Write([]byte(stringToSign))
	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	return hash
}

func (h *Helper) GetTimestampFormat() string {
	now := time.Now()

	return now.Format("20060102150405")
}
