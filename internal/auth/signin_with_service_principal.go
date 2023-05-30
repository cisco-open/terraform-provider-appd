// Copyright 2023 Cisco Systems, Inc.
//
// Licensed under the MPL License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.mozilla.org/en-US/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
