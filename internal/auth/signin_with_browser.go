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
	"fmt"
	"net/url"
	"time"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func loginWithBrowser() (string, error) {
	var loc string
	var uri = authUrl + "?response_type=code&client_id=default&state=abcd&scope=introspect_tokens openid offline_access&redirect_uri=https://accounts.appdynamics.com/redirect.html&code_challenge=8XqBxDPVLPcA0RuyrrRwm8l1JvAJgxh6akne5oOz2QU&code_challenge_method=S256"

	allocatorContext, cancel := chromedp.NewExecAllocator(context.Background())
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	logger.Printf("beginning authentication")

	if err := chromedp.Run(
		ctx,
		chromedp.Navigate(uri),
		chromedp.Location(&loc),
		chromedp.ActionFunc(func(ctx context.Context) error {
			deadline, _ := ctx.Deadline()
			remaining := time.Until(deadline)
			remainingSec := int32(remaining.Seconds())
			_, exp, err := runtime.Evaluate(
				addTimer(int(remainingSec)),
			).Do(ctx)

			if err != nil {
				logger.Printf("error injecting script: %v", err)
				return err
			}

			if exp != nil {
				logger.Printf("error injecting script: %v", exp)
				return exp
			}

			return nil
		}),
	); err != nil {
		cancel()
		logger.Printf("failed to sign in: %v", err)
		return "", fmt.Errorf("failed to sign in")
	}

	isRedirectUri := isRedirectUrl(loc)

	for !isRedirectUri {
		isRedirectUri = isRedirectUrl(loc)
		if err := chromedp.Run(ctx, chromedp.Location(&loc), chromedp.WaitReady("body")); err != nil {
			cancel()
			logger.Printf("failed to get the redirect code: %v", err)
			return "", fmt.Errorf("sign in unsuccessful")
		}
	}

	logger.Printf("sign in successful, proceeding to acquire redirect code")

	url, _ := url.ParseRequestURI(loc)
	code := url.Query().Get("code")

	logger.Printf("redirect code acquired")

	at, err := getAccessToken("default", codeVerifier, code, redirectUrl)
	if err != nil {
		logger.Printf("error fetching access token: %v", err)
		return "", err
	} else {
		logger.Printf("access token fetch successful")
	}

	return at, nil
}

func addTimer(timeRemainingSeconds int) string {
	return fmt.Sprintf(`
		const node = document.createElement("span");
		const textnode = document.createTextNode("");
		node.setAttribute("id", "timer")
		node.style.color = "#fff"
		node.appendChild(textnode);
		
		body = document.querySelector("body")
		body.insertBefore(node, body.children[0])
		
		timeout = %v
		
		setInterval(() => {
			timeout -= 1
			textnode.textContent = "This window will expires in about " + timeout + " seconds"
		}, 1000);
	`, timeRemainingSeconds)
}
