package validatepolltime

import (
	"apics-monitoring/modules"
	"apics-monitoring/restapi"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test_validateNodes(t *testing.T) {
	//given
	tests := []struct {
		nodes           []restapi.Node
		delay           int
		expectedOutcome []modules.Alert
		expectedError   error
	}{
		{
			nodes:           nil,
			delay:           1,
			expectedOutcome: make([]modules.Alert, 0),
			expectedError:   nil,
		},
		{
			nodes:           []restapi.Node{{ContactedAt: "2006-01-02T15:04:05+0000"}},
			delay:           1,
			expectedOutcome: []modules.Alert{{Message: "Gateway last polling time:2006-01-02T15:04:05+0000 breached the threshold:1 seconds"}},
			expectedError:   nil,
		},
		{
			nodes:           []restapi.Node{{ContactedAt: "2006-01-02T15:04:05+0000"}, {ContactedAt: "2006-01-02T15:04:05+0000"}},
			delay:           1,
			expectedOutcome: []modules.Alert{{Message: "Gateway last polling time:2006-01-02T15:04:05+0000 breached the threshold:1 seconds"}, {Message: "Gateway last polling time:2006-01-02T15:04:05+0000 breached the threshold:1 seconds"}},
			expectedError:   nil,
		},
	}
	mod := ValidatePollTime{}
	//when
	for _, test := range tests {
		alerts, err := mod.validateNodes(test.nodes, test.delay)

		//then
		if !reflect.DeepEqual(alerts, test.expectedOutcome) {
			t.Errorf("Not expected outcome: %v, %v", test.expectedOutcome, alerts)
		}

		if test.expectedError != err {
			t.Errorf("Not expected outcome: %+v, %+v", test.expectedError, err)
		}
	}
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
		nodeData:        restapi.Node{ContactedAt: timeNow},
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
