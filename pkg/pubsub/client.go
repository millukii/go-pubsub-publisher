package pubsub

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/pubsub"
)

func NewPubSubClient(ctx context.Context, projectID string) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %v", err)
	}

	return client, nil
}


func CreateTopic(ctx context.Context,w io.Writer, client *pubsub.Client, topicID string) error {

        t, err := client.CreateTopic(ctx, topicID)
        if err != nil {
                return fmt.Errorf("CreateTopic: %v", err)
        }
        fmt.Fprintf(w, "Topic created: %v\n", t)
        return nil
} 
func GetTopic(ctx context.Context,w io.Writer, client *pubsub.Client, topicID string) *pubsub.Topic {

        t := client.Topic(topicID)
        return t
} 
func CreateSubscription(ctx context.Context,w io.Writer, client *pubsub.Client, subscription string, topic *pubsub.Topic) error {

        t, err := client.CreateSubscription(ctx, subscription, pubsub.SubscriptionConfig{
					Topic: topic,
				})
        if err != nil {
                return fmt.Errorf("CreateTopic: %v", err)
        }
        fmt.Fprintf(w, "Topic created: %v\n", t)
        return nil
} 