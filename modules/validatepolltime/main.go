package validatepolltime

import (
	"apics-monitoring/configuration"
	"apics-monitoring/modules"
	"apics-monitoring/restapi"
	"fmt"
	"time"
)

const dateLayout = "2006-01-02T15:04:05+0000"
const alertMessage = "Gateway last polling time:%s breached the threshold:%d seconds"

//ValidatePollTime checks last poll time for the node and generates alert if needed
type ValidatePollTime struct {
}

//Execute holds the main logic
func (v *ValidatePollTime) Execute(token string, config configuration.Configuration) ([]modules.Alert, error) {

	var alerts []modules.Alert

	for _, id := range config.APIGWConditionMonitor.Gateways {
		gateway := restapi.NewGateway(id, config.GetAPIPlatformHost())
		nodes, err := gateway.GetNodesInfo(token)
		if err != nil {
			return nil, err
		}

		alerts, err = v.validateNodes(nodes, config.APIGWConditionMonitor.LastPoolTimeDelay)
	}
	return alerts, nil
}

//GetName returns the name of the module
func (*ValidatePollTime) GetName() string {
	return "ValidatePollTime"
}

func (v *ValidatePollTime) validateNodes(nodes []restapi.Node, delay int) ([]modules.Alert, error) {

	alerts := make([]modules.Alert, 0)
	for _, node := range nodes {
		if alert, error := v.validateNode(node, delay); error != nil {
			return nil, error
		} else {
			alerts = append(alerts, alert)
		}
	}
	return alerts, nil
}

func (*ValidatePollTime) validateNode(node restapi.Node, delay int) (modules.Alert, error) {

	var alert modules.Alert

	pollTime, error := time.Parse(dateLayout, node.ContactedAt)
	if error != nil {
		return alert, error
	}
	pollTime = pollTime.Add(time.Second * time.Duration(delay))
	today := time.Now()

	if pollTime.Before(today) {
		alert = modules.Alert{Message: fmt.Sprintf(alertMessage, node.ContactedAt, delay)}
	}

	return alert, nil
}
