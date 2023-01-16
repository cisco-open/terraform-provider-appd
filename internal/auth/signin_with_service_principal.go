package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
)

func loginWithServicePrincipal(clientId, clientSecret, tenantName, tenantId string) (string, error) {
	conf := clientcredentials.Config{}

	conf.ClientID = clientId
	conf.ClientSecret = clientSecret
	conf.TokenURL = fmt.Sprintf("https://%s.observe.appdynamics.com/auth/%s/default/oauth2/token", tenantName, tenantId)
	log.Printf("calling %s", conf.TokenURL)

	token, err := conf.Token(context.Background())

	// The error string contains response body as json.
	// parse it to display a more concise error message
	if err != nil {
		logger.Printf("Failed to get token using service principal: %v", err.Error())
		if strings.Contains(err.Error(), "Response") {
			var resp map[string]interface{}
			errRespJson := strings.Split(err.Error(), "Response: ")[1]
			json.Unmarshal([]byte(errRespJson), &resp)
			return "", fmt.Errorf("%v - %v", resp["cause"], resp["error_description"])
		}
		return "", fmt.Errorf(err.Error())
	}

	return token.AccessToken, nil
}
