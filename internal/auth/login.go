package auth

import (
	"context"
	"fmt"
)

const (
	LOGIN_HEADLESS               = "headless"
	LOGIN_WITH_BROWSER           = "browser"
	LOGIN_WITH_SERVICE_PRINCIPAL = "service_principal"
)

type credential string

const (
	USERNAME      credential = "username"
	PASSWORD      credential = "password"
	CLIENT_ID     credential = "client_id"
	CLIENT_SECRET credential = "client_secret"
	MODE          credential = "mode"
)

var (
	accessToken string
	err         error
)

func Login(tenantName, tenantId string, saveToken bool, ctx context.Context) (string, error) {
	setCommonUrls(tenantName, tenantId)
	logger.SetPrefix("auth: ")

	// check if token is already present
	// and the present token is valid (not tampered with)
	token, ok := getTokenExists()
	if ok {
		logger.Println("checking token validity")

		valid := checkTokenValid(token, tenantName)
		if valid {
			logger.Println("token is valid")
			return token, nil
		} else {
			logger.Println("token is invalid")
		}
	}

	loginMode := ctx.Value(MODE).(string)

	if !isValidLoginMode(loginMode) {
		logger.Println("unsupported login mode:", loginMode)
		return "", fmt.Errorf("not a valid login mode, supported login modes are: 'headless' and 'browser'")
	}

	logger.Println("signing in with mode:", loginMode)
	logger.SetPrefix(loginMode + ": ")

	if loginMode == LOGIN_HEADLESS {
		username := ctx.Value(USERNAME).(string)
		password := ctx.Value(PASSWORD).(string)

		if !isEmailValid(username) {
			logger.Println("the provided email id seems to be invalid")
			return "", fmt.Errorf("the provided email id seems to be invalid, please recheck or try logging in through browser mode")
		}

		accessToken, err = loginHeadless(username, password)

	} else if loginMode == LOGIN_WITH_BROWSER {
		accessToken, err = loginWithBrowser()
	} else if loginMode == LOGIN_WITH_SERVICE_PRINCIPAL {
		clientId := ctx.Value(CLIENT_ID).(string)
		clientSecret := ctx.Value(CLIENT_SECRET).(string)

		accessToken, err = loginWithServicePrincipal(clientId, clientSecret, tenantName, tenantId)
	}

	if err != nil {
		logger.Println("will not modify token file due to error")
	} else if !saveToken {
		logger.Println("saving token is disabled")
	} else {
		storeToken(accessToken)
	}

	return accessToken, err
}

func isValidLoginMode(mode string) bool {
	return mode == LOGIN_HEADLESS || mode == LOGIN_WITH_BROWSER || mode == LOGIN_WITH_SERVICE_PRINCIPAL
}
