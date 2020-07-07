package configuration

import (
	"encoding/json"
	"os"
)

//Configuration of the application
type Configuration struct {
	CoreConfiguration
	APIGWConditionMonitor GatewayConditionMonitor
	APIGWAlerts           Alerts
}

//GatewayConditionMonitor monitors condition of the gateways
type GatewayConditionMonitor struct {
	LastPoolTimeDelay int
	Gateways          []int
}

//Alerts contains details of Notification service
type Alerts struct {
	TopicID string
}

//LoadConfiguration loads configuration from the file
func LoadConfiguration(fullPath string) (*Configuration, error) {

	config := Configuration{}
	err := config.loadConfiguration(fullPath) //here we load "core" configuraton
	if err != nil {
		return nil, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
