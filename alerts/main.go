package alerts

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/ons"
)

//Send publish the message to the topic
func Send(topicID string, body string) error {

	config := common.DefaultConfigProvider()
	client, err := ons.NewNotificationDataPlaneClientWithConfigurationProvider(config)
	if err != nil {
		return err
	}

	message := ons.PublishMessageRequest{
		TopicId:        &topicID,
		MessageDetails: ons.MessageDetails{Body: &body},
	}

	_, err = client.PublishMessage(context.TODO(), message)
	if err != nil {
		return err
	}

	return nil
}
