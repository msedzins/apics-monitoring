package restapi

import (
	"apics-monitoring/configuration"
	"io/ioutil"
	"net/http"
)

//Authentication gets token from IDCS
type Authentication struct {
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

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	return bodyString, nil
}
