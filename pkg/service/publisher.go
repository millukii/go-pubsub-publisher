package service

import (
	"bytes"
	"context"
	"fmt"
	"publisher/pkg/pubsub"

	gcp "cloud.google.com/go/pubsub"
)


type PublisherService interface {
	Publish(ctx context.Context,event []byte, attributes map[string]string,topic string) error 
   PublishMessages(ctx context.Context,events []map[string]interface{}, topic string) error
}

type publisher struct{
	id int
	publisher pubsub.Publisher
}

func NewPublisherService(p pubsub.Publisher, id int) PublisherService {
	
	return &publisher{
		id: id,
		publisher:p,
	}
}

func (p publisher) Publish(ctx context.Context,event []byte, attributes map[string]string,topic string) error {

	var buf bytes.Buffer


	err := p.publisher.Publish(ctx,&gcp.Message{
		Data: event,
		Attributes: attributes,
	}, &buf,topic)
	if err != nil {
		return err
	}
	results := buf.String()

	fmt.Print(results)
	return nil
}

func (p publisher) PublishMessages(ctx context.Context,events []map[string]interface{}, topic string) error {

	var buf bytes.Buffer


	messages, err := pubsub.ConvertPubsubMessages(events)
	if err != nil {
		return err
	}
	err = p.publisher.PublishWithSettings(ctx,messages, &buf,topic)
	if err != nil {
		return err
	}
	results := buf.String()

	fmt.Print(results)
	return nil
}