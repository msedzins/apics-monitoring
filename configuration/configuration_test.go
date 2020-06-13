package configuration

import (
	"reflect"
	"testing"
)

func Test_LoadConfiguration(t *testing.T) {
	//given
	expectedConfig := Configuration{
		CoreConfiguration: CoreConfiguration{
			apiPlatformClientID:     "test1",
			apiPlatformHost:         "$test2",
			idcsHost:                "test3",
			apiPlatformClientSecret: "test4",
			apiPlatformUser:         "test5",
			apiPlatformUserPassword: "test6",
			apiPlatformScope:        "test7"},
		APIGWConditionMonitor: GatewayConditionMonitor{
			LastPoolTimeDelay: 360,
			Gateways:          []int{1, 2, 3},
		},
	}

	//when
	configFromFile, err := LoadConfiguration("configuration_test.json")
	if err != nil {
		t.Errorf(err.Error())
	}

	//then
	if !reflect.DeepEqual(configFromFile.APIGWConditionMonitor, expectedConfig.APIGWConditionMonitor) ||
		!reflect.DeepEqual(configFromFile.CoreConfiguration, expectedConfig.CoreConfiguration) {
		t.Errorf("configFromFile!=expectedConfig. %+v, %+v", configFromFile, expectedConfig)
	}
}
