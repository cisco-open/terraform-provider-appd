package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

const (
	timeout = 120

	// TODO: generate later
	codeChallenge = "8XqBxDPVLPcA0RuyrrRwm8l1JvAJgxh6akne5oOz2QU"
	codeVerifier  = "RXZ6eUhodnNjVmx6dzZmdTBpcnlyOEtNcnZLeWNZMmg"
)

var logger = log.Default()

// URLs
var (
	redirectUrl string
	authDomain  string
	tokenUrl    string
	authUrl     string
)

func setCommonUrls(tenantName, tenantId string) {
	redirectUrl = "https://accounts.appdynamics.com/redirect.html"
	authDomain = fmt.Sprintf("https://%s.observe.appdynamics.com", tenantName)

	tokenUrl = fmt.Sprintf("%s/auth/%s/default/oauth2/token", authDomain, tenantId)
	authUrl = fmt.Sprintf("%s/auth/%s/default/oauth2/authorize", authDomain, tenantId)
}

func getAccessToken(clientID string, codeVerifier string, authorizationCode string, callbackURL string) (string, error) {
	// set the url and form-encoded data for the POST to the access token endpoint
	url := tokenUrl

	data := fmt.Sprintf(
		"grant_type=authorization_code&client_id=%s"+
			"&code_verifier=%s"+
			"&code=%s"+
			"&redirect_uri=%s",
		clientID, codeVerifier, authorizationCode, callbackURL)
	payload := strings.NewReader(data)

	// create the request and execute it
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("snap: HTTP error: %s", err)
		return "", err
	}

	// process the response
	defer res.Body.Close()
	var responseData map[string]interface{}
	body, _ := io.ReadAll(res.Body)

	// unmarshal the json into a string map
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Printf("snap: JSON error: %s", err)
		return "", err
	}

	// retrieve the access token out of the map, and return to caller
	accessToken := responseData["access_token"]

	if !reflect.ValueOf(accessToken).IsValid() {
		return "", fmt.Errorf("unable to login")
	}

	return accessToken.(string), nil
}

func isRedirectUrl(url string) bool {
	return strings.Contains(url, "redirect.html")
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
