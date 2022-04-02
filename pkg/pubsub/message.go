package pubsub

import (
	"encoding/json"

	"cloud.google.com/go/pubsub"
)

func ConvertPubsubMessages(messages []map[string]interface{}) ([]*pubsub.Message, error) {
	var events []*pubsub.Message
		for _, msg := range messages{
			raw, err := json.Marshal(msg["data"])
			if err != nil {
				return nil, err
			}
			attributes := msg["attributes"].(map[string]string)
			pubsubMessage := &pubsub.Message{
				Data:raw, 
				Attributes: attributes,
			}
			events = append(events, pubsubMessage)
		}
	return events, nil
}

