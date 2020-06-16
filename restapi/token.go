package restapi

import (
	"apics-monitoring/configuration"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Authentication gets token from IDCS
type Authentication struct {
	Client HTTPClient
}

// GetToken gets token from IDCS
func (auth *Authentication) GetToken(conf configuration.Configuration) (string, error) {

	data := url.Values{}
	data.Add("grant_type", "password")
	data.Add("scope", conf.GetAPIPlatformScope())
	data.Add("username", conf.GetAPIPlatformUser())
	data.Add("password", conf.GetAPIPlatformUserPassword())

	req, _ := http.NewRequest(http.MethodPost, conf.GetIDCSHost()+"/oauth2/v1/token", bytes.NewBufferString(data.Encode()))
	req.SetBasicAuth(conf.GetAPIPlatformClientID(), conf.GetAPIPlatformClientSecret())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	response, err := auth.Client.Do(req)
	if err != nil {
		return "", err
	} else if response.StatusCode != 200 {
		return "", errors.New(response.Status)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	token, err := parseResponse(bodyString)
	if err != nil {
		return "", err
	}

	return token, nil
}

//NewAuthentication returns initialized Authentication struct
func NewAuthentication() *Authentication {
	return &Authentication{Client: &http.Client{}}
}

func parseResponse(response string) (string, error) {

	type responseJSON struct {
		AccessToken string `json:"access_token"`
	}

	decoder := json.NewDecoder(strings.NewReader(response))

	json := &responseJSON{}
	err := decoder.Decode(json)
	if err != nil {
		return "", err
	}

	return json.AccessToken, nil
}
