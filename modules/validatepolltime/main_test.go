package validatepolltime

import (
	"apics-monitoring/modules"
	"apics-monitoring/restapi"
	"fmt"
	"testing"
	"time"
)

func Test_validateNodes(t *testing.T) {
	//TODO
}

func Test_validateNode(t *testing.T) {
	//given
	timeNow := time.Now().UTC().Format(dateLayout)
	tests := []struct {
		nodeData        restapi.Node
		delay           int
		expectedOutcome modules.Alert
	}{{
		nodeData:        restapi.Node{ContactedAt: "2020-06-17T20:16:18+0000"},
		delay:           120,
		expectedOutcome: modules.Alert{Message: "Gateway last polling time:2020-06-17T20:16:18+0000 breached the threshold:120 seconds"},
	}, {
		nodeData:        restapi.Node{ContactedAt: timeNow}, //-2 hours because time.Now gets time in UTC, dateLayout will move it to UTC+2
		delay:           120,
		expectedOutcome: modules.Alert{},
	}, {
		nodeData:        restapi.Node{ContactedAt: timeNow},
		delay:           -10,
		expectedOutcome: modules.Alert{Message: fmt.Sprintf(alertMessage, timeNow, -10)},
	}}
	mod := ValidatePollTime{}

	//when
	for _, test := range tests {
		alert, err := mod.validateNode(test.nodeData, test.delay)
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		//then
		if test.expectedOutcome != alert {
			t.Errorf("Not expected alert outcome: %v, %v", test.expectedOutcome, alert)
		}
	}
}
