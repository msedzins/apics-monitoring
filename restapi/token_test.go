package restapi

import (
	"apics-monitoring/configuration"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type ok struct {
}

func (result *ok) Do(req *http.Request) (*http.Response, error) {

	responseJSON := `{
		"access_token": "TOKEN",
		"token_type": "Bearer",
		"expires_in": 3600
	}
	`
	response := http.Response{}
	response.Body = ioutil.NopCloser(strings.NewReader(responseJSON))
	return &response, nil
}

func Test_GetToken_OK(t *testing.T) {
	//given
	auth := Authentication{Client: &ok{}}
	conf := configuration.Configuration{}

	//when
	response, err := auth.GetToken(conf)
	if err != nil {
		t.Errorf(err.Error())
	}

	//then
	if response != "TOKEN" {
		t.Errorf("Response should == TOKEN. %s", response)
	}
}

type failed struct {
}

func (result *failed) Do(req *http.Request) (*http.Response, error) {

	responseJSON := `{
		"token_type": "Bearer",
		"expires_in": 3600
	}
	`
	response := http.Response{}
	response.Body = ioutil.NopCloser(strings.NewReader(responseJSON))
	return &response, nil
}

func Test_GetToken_Failed(t *testing.T) {
	//given
	auth := Authentication{Client: &failed{}}
	conf := configuration.Configuration{}

	//when
	response, err := auth.GetToken(conf)
	if err != nil {
		t.Errorf(err.Error())
	}

	//then
	if response != "" {
		t.Errorf("Response should be empty. %s", response)
	}
}
