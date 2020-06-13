package configuration

import (
	"os"
	"testing"
)

func Test_loadConfiguration(t *testing.T) {
	//given
	expectedConfig := CoreConfiguration{
		apiPlatformClientID:     "test1",
		apiPlatformHost:         "$test2",
		idcsHost:                "test3",
		apiPlatformClientSecret: "test4",
		apiPlatformUser:         "test5",
		apiPlatformUserPassword: "test6",
	}
	config := CoreConfiguration{}

	//when
	err := config.loadConfiguration("configuration_test.json")
	if err != nil {
		t.Errorf(err.Error())
	}

	//then
	if config != expectedConfig {
		t.Errorf("config!=expectedConfig. %+v, %+v", config, expectedConfig)
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
			fileContent:     internalConfiguration{APIPlatformClientID: "1", APIPlatformHost: "2", APIPlatformClientSecret: "3", APIPlatformUserPassword: "4", IDCSHost: "5", APIPlatformUser: "6"},
			setEnv:          func() {},
			expectedOutcome: map[string]string{"APIPlatformClientID": "1", "APIPlatformHost": "2", "APIPlatformClientSecret": "3", "APIPlatformUserPassword": "4", "IDCSHost": "5", "APIPlatformUser": "6"},
		}, {
			fileContent: internalConfiguration{APIPlatformClientID: "$t1", APIPlatformHost: "$t2", APIPlatformClientSecret: "$t3", APIPlatformUserPassword: "$t4", IDCSHost: "$t5", APIPlatformUser: "$t6"},
			setEnv: func() {
				os.Setenv("t1", "test1")
				os.Setenv("t2", "test2")
				os.Setenv("t3", "test3")
				os.Setenv("t4", "test4")
				os.Setenv("t5", "test5")
				os.Setenv("t6", "test6")
			},
			expectedOutcome: map[string]string{"APIPlatformClientID": "test1", "APIPlatformHost": "test2", "APIPlatformClientSecret": "test3", "APIPlatformUserPassword": "test4", "IDCSHost": "test5", "APIPlatformUser": "test6"},
		}}

	for _, test := range tests {
		//when
		config := CoreConfiguration{}
		test.fileContent.setConfiguration(&config)
		test.setEnv()

		//then
		if config.GetAPIPlatformClientID() != test.expectedOutcome["APIPlatformClientID"] ||
			config.GetAPIPlatformHost() != test.expectedOutcome["APIPlatformHost"] ||
			config.GetAPIPlatformClientSecret() != test.expectedOutcome["APIPlatformClientSecret"] ||
			config.GetAPIPlatformUserPassword() != test.expectedOutcome["APIPlatformUserPassword"] ||
			config.GetIDCSHost() != test.expectedOutcome["IDCSHost"] {
			t.Errorf("Test failed: %+v", test)
		}
	}
}
