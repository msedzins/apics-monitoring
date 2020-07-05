package restapi

import (
	"apics-monitoring/configuration"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type authOk struct {
}
type authFailed struct {
}
type authFailed2 struct {
}

func (result *authOk) Do(req *http.Request) (*http.Response, error) {

	responseJSON := `{
		"access_token": "TOKEN",
		"token_type": "Bearer",
		"expires_in": 3600
	}
	`
	response := http.Response{}
	response.StatusCode = 200
	response.Body = ioutil.NopCloser(strings.NewReader(responseJSON))
	return &response, nil
}

func (result *authFailed) Do(req *http.Request) (*http.Response, error) {

	responseJSON := `{
		"token_type": "Bearer",
		"expires_in": 3600
	}
	`
	response := http.Response{}
	response.StatusCode = 200
	response.Body = ioutil.NopCloser(strings.NewReader(responseJSON))
	return &response, nil
}

func (result *authFailed2) Do(req *http.Request) (*http.Response, error) {

	responseJSON := ""
	response := http.Response{}
	response.StatusCode = 500
	response.Status = "500"
	response.Body = ioutil.NopCloser(strings.NewReader(responseJSON))
	return &response, nil
}

func Test_GetToken_OK(t *testing.T) {
	//given
	auth := Authentication{client: &authOk{}}
	conf := configuration.Configuration{}

	//when
	response, err := auth.GetToken(conf)
	if err != nil {
		t.Errorf("Error: " + err.Error())
		return
	}

	//then
	if response != "TOKEN" {
		t.Errorf("Response should == TOKEN. %s", response)
	}
}

func Test_GetToken_Failed(t *testing.T) {
	//given
	auth := Authentication{client: &authFailed{}}
	conf := configuration.Configuration{}

	//when
	response, err := auth.GetToken(conf)
	if err != nil {
		t.Errorf("Error: " + err.Error())
		return
	}

	//then
	if response != "" {
		t.Errorf("Response should be empty. %s", response)
	}
}

func Test_GetToken_Failed_2(t *testing.T) {
	//given
	auth := Authentication{client: &authFailed2{}}
	conf := configuration.Configuration{}

	//when
	_, err := auth.GetToken(conf)

	//then
	if err == nil || err.Error() != "Bad response status 500" {
		t.Errorf(err.Error())
	}
}
