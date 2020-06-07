package modules

// Module contains set of methods that should be implemented by all modules
type Module interface {

	//Execute the action on the module and returns list of alerts, if any
	Execute() ([]Alert, error)
}

// Alert represents the notification that is returned by the module
type Alert struct {
	Message string
}
