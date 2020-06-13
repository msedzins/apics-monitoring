package configuration

import (
	"encoding/json"
	"os"
	"strings"
)

// CoreConfiguration of the application
type CoreConfiguration struct {
	apiPlatformClientID     string
	apiPlatformHost         string
	idcsHost                string
	apiPlatformClientSecret string
	apiPlatformUser         string
	apiPlatformUserPassword string
	apiPlatformScope        string
}

// GetAPIPlatformScope returns value from configuration file or env. variable.
func (conf *CoreConfiguration) GetAPIPlatformScope() string {
	if strings.HasPrefix(conf.apiPlatformScope, "$") {
		return os.Getenv(conf.apiPlatformScope[1:])
	}
	return conf.apiPlatformScope
}

// GetAPIPlatformClientID returns value from configuration file or env. variable.
func (conf *CoreConfiguration) GetAPIPlatformClientID() string {
	if strings.HasPrefix(conf.apiPlatformClientID, "$") {
		return os.Getenv(conf.apiPlatformClientID[1:])
	}
	return conf.apiPlatformClientID
}

// GetAPIPlatformHost returns value from configuration file or env. variable.
func (conf *CoreConfiguration) GetAPIPlatformHost() string {
	if strings.HasPrefix(conf.apiPlatformHost, "$") {
		return os.Getenv(conf.apiPlatformHost[1:])
	}
	return conf.apiPlatformHost
}

// GetIDCSHost returns value from configuration file or env. variable.
func (conf *CoreConfiguration) GetIDCSHost() string {
	if strings.HasPrefix(conf.idcsHost, "$") {
		return os.Getenv(conf.idcsHost[1:])
	}
	return conf.idcsHost
}

// GetAPIPlatformClientSecret returns value from configuration file or env. variable.
func (conf *CoreConfiguration) GetAPIPlatformClientSecret() string {
	if strings.HasPrefix(conf.apiPlatformClientSecret, "$") {
		return os.Getenv(conf.apiPlatformClientSecret[1:])
	}
	return conf.apiPlatformClientSecret
}

// GetAPIPlatformUserPassword returns value from configuration file or env. variable.
func (conf *CoreConfiguration) GetAPIPlatformUserPassword() string {
	if strings.HasPrefix(conf.apiPlatformUserPassword, "$") {
		return os.Getenv(conf.apiPlatformUserPassword[1:])
	}
	return conf.apiPlatformUserPassword
}

// GetAPIPlatformUser returns value from configuration file or env. variable.
func (conf *CoreConfiguration) GetAPIPlatformUser() string {
	if strings.HasPrefix(conf.apiPlatformUser, "$") {
		return os.Getenv(conf.apiPlatformUser[1:])
	}
	return conf.apiPlatformUser
}

//loads the configuration from the file to CoreConfiguration struct
func (conf *CoreConfiguration) loadConfiguration(fullPath string) error {

	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	internalConf := internalConfiguration{}
	err = decoder.Decode(&internalConf)
	if err != nil {
		return err
	}

	internalConf.setConfiguration(conf)
	return nil
}

//temporary structure for holding data loaded from configuration file
type internalConfiguration struct {
	APIPlatformClientID     string
	APIPlatformHost         string
	IDCSHost                string
	APIPlatformClientSecret string
	APIPlatformUser         string
	APIPlatformUserPassword string
	APIPlatformScope        string
}

//copies data from temporary structure to CoreConfiguration
func (config *internalConfiguration) setConfiguration(newConf *CoreConfiguration) {

	newConf.apiPlatformClientID = config.APIPlatformClientID
	newConf.apiPlatformHost = config.APIPlatformHost
	newConf.idcsHost = config.IDCSHost
	newConf.apiPlatformClientSecret = config.APIPlatformClientSecret
	newConf.apiPlatformUser = config.APIPlatformUser
	newConf.apiPlatformUserPassword = config.APIPlatformUserPassword
	newConf.apiPlatformScope = config.APIPlatformScope
}
