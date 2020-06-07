package configuration

import (
	"encoding/json"
	"os"
	"strings"
)

// Configuration of the application
type Configuration struct {
	apiPlatformClientID string
	apiPlatformHost     string
}

// GetAPIPlatformClientID returns value from configuration file or env. variable.
func (conf *Configuration) GetAPIPlatformClientID() string {
	if strings.HasPrefix(conf.apiPlatformClientID, "$") {
		return os.Getenv(conf.apiPlatformClientID[1:])
	}
	return conf.apiPlatformClientID
}

// GetAPIPlatformHost returns value from configuration file or env. variable.
func (conf *Configuration) GetAPIPlatformHost() string {
	if strings.HasPrefix(conf.apiPlatformHost, "$") {
		return os.Getenv(conf.apiPlatformHost[1:])
	}
	return conf.apiPlatformHost
}

// LoadConfiguration loads configuration from the file
func LoadConfiguration(fullPath string) (*Configuration, error) {

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := internalConfiguration{}
	err = decoder.Decode(&conf)
	if err != nil {
		return nil, err
	}

	return conf.getConfiguration(), err
}

type internalConfiguration struct {
	APIPlatformClientID string
	APIPlatformHost     string
}

func (config *internalConfiguration) getConfiguration() *Configuration {
	return &Configuration{
		apiPlatformClientID: config.APIPlatformClientID,
		apiPlatformHost:     config.APIPlatformHost,
	}
}
