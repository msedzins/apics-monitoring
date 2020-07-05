package modules

import "apics-monitoring/configuration"

// Module contains set of methods that should be implemented by all modules
type Module interface {

	//Execute the action on the module and returns list of alerts, if any
	Execute(token string, config configuration.Configuration) ([]Alert, error)

	//Returns module name, for auditing purposes
	GetName() string
}

// Alert represents the notification that is returned by the module
type Alert struct {
	Message string
}
