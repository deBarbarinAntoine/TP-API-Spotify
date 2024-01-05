package spotifyAPI

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	token       string
	baseURL     string = "https://accounts.spotify.com/api/"
	albumJulURL string
	trackSdmURL string
)

const (
	clientID     string = "73d4066635574286b9ad26767386dff6"
	clientSecret string = "84e78d80173c43dbb50fe267a5c9caf3"
)

func GetToken() {

	tokenEndpoint := "token"
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodPost, baseURL+tokenEndpoint, nil)
	if errReq != nil {
		log.Fatal("Request failed!\n", errReq)
	}

	q := req.URL.Query()
	q.Add("grant_type", "client_credentials")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		log.Fatal("Sending request error!\n", errRes)
	}

	var tokenResp map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&tokenResp); err != nil {
		log.Fatal("Json decode error!\n", err)
	}

	var ok bool
	token, ok = tokenResp["access_token"].(string)
	if !ok {
		log.Fatal("token d'accès introuvable!")
	}

	fmt.Printf("log\tgetToken()\ttoken: %#v\n", token)
}
