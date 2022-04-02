package pubsub

import (
	"context"
	"log"
	"publisher/pkg/model/in"
	"time"

	"cloud.google.com/go/pubsub"

)

type Subscriber interface {
Pull(ctx context.Context, subscription string,  collection []in.Event) error 
}

type subscriber struct {
	client *pubsub.Client
}

func NewSuscriber(client *pubsub.Client)  Subscriber {
	return &subscriber{
		client: client,
	}
}

func (s *subscriber) Pull(ctx context.Context, subscription string, collection []in.Event) error {
  
	sub :=	s.client.Subscription(subscription)

	sub.ReceiveSettings.MaxExtension = -1
	sub.ReceiveSettings.MaxExtensionPeriod = -1
	sub.ReceiveSettings.MaxOutstandingMessages = 3
	sub.ReceiveSettings.MaxOutstandingBytes = 1000 
	sub.ReceiveSettings.NumGoroutines = 2
	sub.ReceiveSettings.Synchronous = false
	
 err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
	log.Printf("Got message %s at %s with data %s and attibutes: %s", m.ID, m.PublishTime,m.Data, m.Attributes)
	collection = append(collection,in.Event{
		Data: m.Data,
		PublishTime: m.PublishTime,
		Attributes: m.Attributes,
		DeliveryAttempt: m.DeliveryAttempt,
		ConsumedTime: time.Now(),
	})
		log.Println("ack")
	m.Ack()

 })
 if err != nil {
	return err
 }

 log.Println("Got messages without  error")
 return nil
}