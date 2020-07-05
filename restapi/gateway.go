package restapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//Gateway represents and interacs with logical gateway
type Gateway struct {
	client          HTTPClient
	id              int
	apiPlatformHost string
}

//Node represents one physical node of logical gateway
type Node struct {
	ContactedAt string `json:"contactedAt"`
	Name        string `json:"name"`
}

// GetNodesInfo gets information about nodes of logical gateway from API Platform
func (gt *Gateway) GetNodesInfo(authToken string) ([]Node, error) {

	url := gt.apiPlatformHost + "/apiplatform/management/v1/gateways/" + strconv.Itoa(gt.id) + "/nodes"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", "Bearer")

	response, err := gt.client.Do(req)
	if err != nil {
		return nil, err
	} else if response.StatusCode != 200 {
		return nil, errors.New("Bad response status " + response.Status)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)

	nodes, err := parseGatewayResponse(bodyString)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

//NewGateway returns initialized Gateway struct
func NewGateway(id int, apiPlatformHost string) *Gateway {
	return &Gateway{
		client:          &http.Client{},
		id:              id,
		apiPlatformHost: apiPlatformHost}
}

func parseGatewayResponse(response string) ([]Node, error) {

	type responseJSON struct {
		Items []Node `json:"items"`
	}

	decoder := json.NewDecoder(strings.NewReader(response))

	json := &responseJSON{}
	err := decoder.Decode(json)
	if err != nil {
		return nil, err
	}

	return json.Items, nil
}
