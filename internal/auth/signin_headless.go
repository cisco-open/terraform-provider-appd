package auth

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func loginHeadless(username, password string) (string, error) {
	var (
		loc   string
		nodes []*cdp.Node

		inputUser      = `//input[@name="username"]`
		inputPassword  = `//input[@name="password"]`
		emailSubmit    = `idp-discovery-submit`
		oktaSignIn     = `okta-signin-submit`
		errorContainer = `.o-form-error-container`

		uri = authUrl + "?response_type=code&client_id=default&state=abcd&scope=introspect_tokens openid offline_access&redirect_uri=https://accounts.appdynamics.com/redirect.html&code_challenge=8XqBxDPVLPcA0RuyrrRwm8l1JvAJgxh6akne5oOz2QU&code_challenge_method=S256"
	)

	allocatorContext, cancel := chromedp.NewExecAllocator(context.Background(), chromedp.Headless)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	// navigate to the auth page
	if err := chromedp.Run(ctx, chromedp.Navigate(uri), chromedp.Location(&loc)); err != nil {
		logger.Printf("Failed to navigate to the auth url: %v", err)
		cancel()
		return "", fmt.Errorf("unable to sign in: failed to load the auth page")
	}

	log.Printf("on auth page, beginning to inject credentials")

	isRedirectUri := isRedirectUrl(loc)

	if !isRedirectUri {
		if err := chromedp.Run(ctx,
			chromedp.WaitVisible(inputUser),
			chromedp.Clear(inputUser),
			chromedp.SendKeys(inputUser, username),

			chromedp.WaitReady(emailSubmit, chromedp.ByID),
			chromedp.Click(emailSubmit, chromedp.ByID),

			chromedp.WaitVisible(inputPassword),
			chromedp.Clear(inputPassword),
			chromedp.SendKeys(inputPassword, password),

			chromedp.WaitReady(oktaSignIn, chromedp.ByID),
			chromedp.Click(oktaSignIn, chromedp.ByID),

			chromedp.Nodes(errorContainer, &nodes, chromedp.ByQuery),
		); err != nil {
			cancel()
			logger.Printf("unable to sign in: %v", err)
			return "", fmt.Errorf("unable to sign in")
		}

		logger.Printf("credentials injected successfully, waiting for sign in")

		// now that we have injected the credentials
		// wait for the sign in button to get activated.
		// if it gets activated again, it means that we are still on the same page
		// and possible that user has entered invalid credentials
		//
		// if the button does not get activated again
		// indicating that we are successfully redirected.
		ctx10Sec, cancel10Sec := context.WithTimeout(ctx, 15*time.Second)
		defer cancel10Sec()

		if err := chromedp.Run(ctx10Sec,
			chromedp.WaitEnabled(oktaSignIn, chromedp.ByID),
			// error node will contain the sign in error message, if present
			chromedp.Nodes(errorContainer, &nodes, chromedp.ByQuery),
		); err != nil {
			// we are expecting deadline to exceed, as it mostly means that we have been redirected
			// there is no explicit way to check for redirection with chromedp
			// if the page is redirected, the button no longer exist and will never get enabled
			// so the timeout will occur.
			// but report is any other error has occurred.
			if err.Error() != "context deadline exceeded" {
				logger.Printf("error happened waiting for submit button to get activated: %v", err)
			}
		}

		// check if error occurred signing in
		isErrorPresent := nodes[0].ChildNodeCount > 0

		if isErrorPresent {
			logger.Printf("some error happened trying to sign in")

			if err := chromedp.Run(ctx10Sec, chromedp.Nodes(`p`, &nodes, chromedp.ByQuery)); err != nil {
				cancel10Sec()
				logger.Println("unable to get sign in error message")
				return "", fmt.Errorf("unable to sign in")
			}

			errorMessage := nodes[0].Children[0].NodeValue

			logger.Printf("failed to sign in: %v", errorMessage)
			return "", fmt.Errorf(errorMessage)
		}

		logger.Printf("signed in successfully, proceeding to acquire redirect code")

		// TODO: check if go routine would work here
		// wait for redirect until until timeout
		for !isRedirectUri {
			isRedirectUri = isRedirectUrl(loc)
			if err := chromedp.Run(ctx, chromedp.Location(&loc), chromedp.WaitReady("body")); err != nil {
				cancel()
				logger.Printf("failed to get the auth code: %v", err)
				return "", fmt.Errorf("sign in was successful, but failed to get the auth code")
			}
		}

	}

	url, _ := url.ParseRequestURI(loc)
	code := url.Query().Get("code")

	logger.Printf("redirect code acquired")

	at, err := getAccessToken("default", codeVerifier, code, redirectUrl)
	if err != nil {
		logger.Printf("error fetching access token: %v", err)
	} else {
		log.Printf("access token fetch successful")
	}

	return at, nil
}
