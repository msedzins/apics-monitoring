package alerts

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity" // Identity or any other service you wish to make requests to
)

func test() {

	config := common.DefaultConfigProvider()
	_, err := identity.NewIdentityClientWithConfigurationProvider(config)
	if err != nil {
		panic(err)
	}
}
