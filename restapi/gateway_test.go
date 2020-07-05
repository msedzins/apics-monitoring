package restapi

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type gatewayOk struct {
}
type gatewayOkNoNodes struct {
}
type gatewayFailed struct {
}

func (result *gatewayOk) Do(req *http.Request) (*http.Response, error) {

	responseJSON := `{
		"items": [
			{
				"contactedAt": "2020-06-17T20:16:18+0000",
				"name": "devapigateway01"
			},
			{
				"contactedAt": "2020-06-19T21:16:18+0000",
				"name": "devapigateway02"
			}
		]
	}
	`
	response := http.Response{}
	response.StatusCode = 200
	response.Body = ioutil.NopCloser(strings.NewReader(responseJSON))
	return &response, nil
}

func (result *gatewayOkNoNodes) Do(req *http.Request) (*http.Response, error) {

	responseJSON := `{
		"items": [
		]
	}
	`
	response := http.Response{}
	response.StatusCode = 200
	response.Body = ioutil.NopCloser(strings.NewReader(responseJSON))
	return &response, nil
}

func (result *gatewayFailed) Do(req *http.Request) (*http.Response, error) {

	responseJSON := ""
	response := http.Response{}
	response.StatusCode = 500
	response.Status = "500"
	response.Body = ioutil.NopCloser(strings.NewReader(responseJSON))
	return &response, nil
}

func Test_GetNodes_OK(t *testing.T) {
	//given
	gateway := Gateway{
		client:          &gatewayOk{},
		apiPlatformHost: "DUMMY",
		id:              0}

	//when
	response, err := gateway.GetNodesInfo("DUMMY_TOKEN")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	//then
	if response == nil || len(response) != 2 ||
		response[0].Name != "devapigateway01" || response[0].ContactedAt != "2020-06-17T20:16:18+0000" ||
		response[1].Name != "devapigateway02" || response[1].ContactedAt != "2020-06-19T21:16:18+0000" {
		t.Errorf("Bad response returned. %+v", response)
	}
}

func Test_GetNodes_OK_NoNodes(t *testing.T) {
	//given
	gateway := Gateway{
		client:          &gatewayOkNoNodes{},
		apiPlatformHost: "DUMMY",
		id:              0}

	//when
	response, err := gateway.GetNodesInfo("DUMMY_TOKEN")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	//then
	if response == nil || len(response) != 0 {
		t.Errorf("Bad response returned. %+v", response)
	}
}

func Test_GetNodes_Failed(t *testing.T) {
	//given
	gateway := Gateway{
		client:          &gatewayFailed{},
		apiPlatformHost: "DUMMY",
		id:              0}

	//when
	_, err := gateway.GetNodesInfo("DUMMY_TOKEN")

	//then
	if err == nil || err.Error() != "Bad response status 500" {
		t.Errorf(err.Error())
	}
}
