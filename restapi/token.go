package restapi

import (
	"apics-monitoring/configuration"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//Authentication gets token from IDCS
type Authentication struct {
	Client HTTPClient
}

// GetToken gets token from IDCS
func (auth *Authentication) GetToken(conf configuration.Configuration) (string, error) {

	req, _ := http.NewRequest(http.MethodPost, conf.GetIDCSHost(), nil)
	req.SetBasicAuth(conf.GetAPIPlatformClientID(), conf.GetAPIPlatformClientSecret())
	req.ParseForm()
	req.PostForm.Add("grant_type", "password")
	req.PostForm.Add("scope", conf.GetAPIPlatformScope())
	req.PostForm.Add("username", conf.GetAPIPlatformUser())
	req.PostForm.Add("password", conf.GetAPIPlatformUserPassword())

	response, err := auth.Client.Do(req)
	if err != nil {
		return "", err
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
