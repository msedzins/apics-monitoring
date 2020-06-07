package configuration

import (
	"os"
	"testing"
)

func Test_LoadConfiguration(t *testing.T) {
	//given
	expectedConfig := Configuration{
		apiPlatformClientID: "test1",
		apiPlatformHost:     "$test2",
	}

	//when
	config, err := LoadConfiguration("configuration_test.json")
	if err != nil {
		t.Errorf(err.Error())
	}

	//then
	if *config != expectedConfig {
		t.Errorf("config!=expectedConfig. %+v, %+v", *config, expectedConfig)
	}
}

func Test_Configuration(t *testing.T) {
	//given
	tests := []struct {
		fileContent     internalConfiguration
		setEnv          func()
		expectedOutcome map[string]string
	}{
		{
			fileContent:     internalConfiguration{APIPlatformClientID: "1", APIPlatformHost: "2"},
			setEnv:          func() {},
			expectedOutcome: map[string]string{"APIPlatformClientID": "1", "APIPlatformHost": "2"},
		}, {
			fileContent:     internalConfiguration{APIPlatformClientID: "$t1", APIPlatformHost: "$t2"},
			setEnv:          func() { os.Setenv("t1", "test1"); os.Setenv("t2", "test2") },
			expectedOutcome: map[string]string{"APIPlatformClientID": "test1", "APIPlatformHost": "test2"},
		}}

	for _, test := range tests {
		config := test.fileContent.getConfiguration()
		test.setEnv()

		if config.GetAPIPlatformClientID() != test.expectedOutcome["APIPlatformClientID"] ||
			config.GetAPIPlatformHost() != test.expectedOutcome["APIPlatformHost"] {
			t.Errorf("Test failed: %+v", test)
		}
	}
}
