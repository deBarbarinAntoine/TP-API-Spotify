package spotifyAPI

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	token       string
	baseURL     string = "https://accounts.spotify.com/api/"
	albumJulURL string
	trackSdmURL string
	httpClient  = http.Client{
		Timeout: time.Second * 2,
	}
)

const (
	clientID     string = "73d4066635574286b9ad26767386dff6"
	clientSecret string = "84e78d80173c43dbb50fe267a5c9caf3"
)

func getToken() {

	const tokenEndpoint = "token"

	req, errReq := http.NewRequest(http.MethodPost, baseURL+tokenEndpoint, nil)
	if errReq != nil {
		log.Fatal("log: getToken()\tRequest failed!\n", errReq)
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
		log.Fatal("log: getToken()\tSending request error!\n", errRes)
	}

	var tokenResp map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&tokenResp); err != nil {
		log.Fatal("log: getToken()\tJson decode error!\n", err)
	}

	var ok bool
	token, ok = tokenResp["access_token"].(string)
	if !ok {
		log.Fatal("log: getToken()\ttoken d'accès introuvable!")
	}

	fmt.Printf("log: getToken()\ttoken: %#v\n", token)
}

func getAlbumsJul() AlbumsData {
	getToken()

	fmt.Println("log: getAlbumsJul()\tToken: ", token)

	var data = AlbumsData{}

	const urlRequest = "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy/albums"

	req, errReq := http.NewRequest(http.MethodGet, urlRequest, nil)
	if errReq != nil {
		log.Fatal("log: getAlbumsJul()\tRequest failed!\n", errReq)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		log.Fatal("log: getAlbumsJul()\tSending request error!\n", errRes)
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		log.Fatal("Reading body request error !\n", errBody)
	}

	errJSON := json.Unmarshal(body, &data)
	if errJSON != nil {
		log.Fatal("log: getAlbumsJul()\tData format error !\n", errJSON)
	}

	fmt.Printf("log: getAlbumsJul()\tdata.Items length: %#v\n", len(data.Items))

	return data
}

func getTrackSdm() TrackData {
	getToken()

	fmt.Println("log: getTrackSdm()\tToken: ", token)

	var data = TrackData{}

	const urlRequest = "https://api.spotify.com/v1/tracks/7A1nhuX64Y2JB206h3FjBK"

	req, errReq := http.NewRequest(http.MethodGet, urlRequest, nil)
	if errReq != nil {
		log.Fatal("log: getTrackSdm()\tRequest failed!\n", errReq)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		log.Fatal("log: getTrackSdm()\tSending request error!\n", errRes)
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		log.Fatal("log: getTrackSdm()\tReading body request error !\n", errBody)
	}

	errJSON := json.Unmarshal(body, &data)
	if errJSON != nil {
		log.Fatal("log: getTrackSdm()\tData format error !\n", errJSON)
	}

	fmt.Printf("log: getTrackSdm()\tdata: %#v\n", data)

	return data
}

func getArtistData(id string) ArtistData {
	var data = ArtistData{}

	urlRequest := "https://api.spotify.com/v1/artists/" + id

	req, errReq := http.NewRequest(http.MethodGet, urlRequest, nil)
	if errReq != nil {
		log.Fatal("log: getArtistData()\tRequest failed!\n", errReq)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		log.Fatal("log: getArtistData()\tSending request error!\n", errRes)
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		log.Fatal("log: getArtistData()\tReading body request error !\n", errBody)
	}

	errJSON := json.Unmarshal(body, &data)
	if errJSON != nil {
		log.Fatal("log: getArtistData()\tData format error !\n", errJSON)
	}

	fmt.Printf("log: getArtistData()\tdata: %#v\n", data)

	return data
}
