package snap

type ResponseAccessToken struct {
	AccessToken     string `json:"accessToken"`
	ResponseCode    string `json:"responseCode,omitempty"`
	ResponseMessage string `json:"responseMessage,omitempty"`
	TokenType       string `json:"tokenType"`
	ExpiresIn       string `json:"expiresIn"`
}
