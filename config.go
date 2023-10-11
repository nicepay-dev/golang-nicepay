package library

import (
	"crypto/rsa"
)

type Config struct {
	IsProduction bool
	PrivateKey   *rsa.PrivateKey
	ClientSecret string
	ClientID     string
}


func (c *Config) GetConfiguration() map[string]interface{} {
	return map[string]interface{}{
		"isProduction": c.IsProduction,
		"privateKey":   c.PrivateKey,
		"clientSecret": c.ClientSecret,
		"clientId":     c.ClientID,
	}
}

func (c *Config) SetConfiguration(options map[string]interface{}) {
	mergedConfig := c.GetConfiguration()

	for key, value := range options {
		if key == "isProduction" {
			if isProduction, ok := value.(bool); ok {
				mergedConfig["isProduction"] = isProduction
			}
		} else if key == "privateKey" {
			if privateKey, ok := value.(*rsa.PrivateKey); ok {
				mergedConfig["privateKey"] = privateKey
			}
		} else if key == "clientSecret" {
			if clientSecret, ok := value.(string); ok {
				mergedConfig["clientSecret"] = clientSecret
			}
		} else if key == "clientId" {
			if clientID, ok := value.(string); ok {
				mergedConfig["clientId"] = clientID
			}
		}
	}

	c.IsProduction = mergedConfig["isProduction"].(bool)
	c.PrivateKey = mergedConfig["privateKey"].(*rsa.PrivateKey)
	c.ClientSecret = mergedConfig["clientSecret"].(string)
	c.ClientID = mergedConfig["clientId"].(string)
}

func (c *Config) GetSnapAPIBaseURL() string {
	if c.IsProduction {
		return "https://www.nicepay.co.id/nicepay"
	} else {
		return "https://dev.nicepay.co.id/nicepay"
	}
}


