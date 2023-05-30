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
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"io"
	"os"
	"time"

	client "github.com/aniketk-crest/appdynamicscloud-go-client"
	cloudconnectionapi "github.com/aniketk-crest/appdynamicscloud-go-client/apis/v1/cloudconnections"
)

var key = []byte("testtesttesttest")
var tokenFile = "./.token.enc"

type TokenStore struct {
	Token     string
	ExpiresAt time.Time
}

func storeToken(token string) {
	logger.Println("writing token to", tokenFile)

	file, err := os.Create(tokenFile)
	if err != nil {
		logger.Println("an error occurred writing token file", err)
		return
	}

	defer file.Close()

	// encoder := gob.NewEncoder(file)
	tokenStore := TokenStore{}

	// token expiry at the time of writing this is 1 hour. thus, storing the time+1 hour here
	// this is expecting that the API behavior will not change that often.
	tokenStore.Token = token
	tokenStore.ExpiresAt = time.Now().Add(1 * time.Hour)

	logger.Println("written token will expire at", tokenStore.ExpiresAt)

	file.Write([]byte(encrypt(tokenStore)))

	// encoder.Encode(tokenStore)
}

func getTokenExists() (string, bool) {
	logger.Println("checking if the token is already present at", tokenFile)

	file, err := os.Open(tokenFile)

	if os.IsNotExist(err) {
		logger.Println("token file is not available, will continue to sign in")
		return "", false
	} else if err != nil {
		logger.Println("an error occurred reading token file", err)
		return "", false
	} else {
		defer file.Close()

		tokenStore := TokenStore{}

		logger.Printf("reading encrypted file")
		ciphertext, err := os.ReadFile(tokenFile)
		if err != nil {
			logger.Println("error reading encrypted token file", err)
			return "", false
		}

		logger.Println("file read successful. decrypting...")
		plainText := decrypt(string(ciphertext))
		var plainTextBuffer bytes.Buffer
		plainTextBuffer.WriteString(plainText)

		logger.Println("file decryption successful. decoding...")
		decoder := gob.NewDecoder(&plainTextBuffer)
		decoder.Decode(&tokenStore)

		tokenExpiresIn := time.Since(tokenStore.ExpiresAt).Minutes()
		if tokenExpiresIn > -10 {
			logger.Println("access token is present, but has expired, invalid or tampered")
			logger.Println("continuing to sign in")

			return "", false
		} else {
			logger.Println("access token is present, it expires at", tokenStore.ExpiresAt)
			return tokenStore.Token, true
		}
	}
}

func checkTokenValid(token, tenantName string) bool {
	configuration := client.NewConfiguration()
	configuration.Debug = true

	apiClient := cloudconnectionapi.NewAPIClient(configuration)

	myctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
		"tenant-name": tenantName,
	})
	myctx = context.WithValue(myctx, client.ContextServerIndex, client.SERVER_INDEX_CLOUD_CONNECTION)
	myctx = context.WithValue(myctx, client.ContextAccessToken, token)

	_, _, err := apiClient.ConnectionsApi.GetConnections(myctx).Execute()
	if err != nil {
		logger.Println(err)
		return false
	}

	return true
}

// === Cryptography Utils ===

func decrypt(cipherstring string) string {
	// Byte array of the string
	ciphertext := []byte(cipherstring)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Println("error occurred creating the AES cipher", err)
		return ""
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		logger.Println("Text is too short")
		return ""
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

func encrypt(token TokenStore) string {
	var tokenBytes bytes.Buffer
	enc := gob.NewEncoder(&tokenBytes)

	err := enc.Encode(token)
	if err != nil {
		logger.Println("error occurred encoding token during encryption", err)
		return ""
	}

	plaintext := tokenBytes.Bytes()

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Println("error occurred creating the AES cipher", err)
		return ""
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		logger.Println("error occurred filling initialization vector during encryption", err)
		return ""
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}
