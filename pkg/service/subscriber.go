package service

import (
	"context"
	"log"
	"publisher/pkg/model/in"
	"publisher/pkg/pubsub"

)

type SubscriberService interface {
	ReceiveMessages(ctx context.Context, subscription string, collection []in.Event) error
}
type suscriber struct {
	id int
	suscriber pubsub.Subscriber
}

func NewSubscriberService(s pubsub.Subscriber, id int) SubscriberService    {

	return &suscriber{
		id: id,
		suscriber: s,
	}
}

func (s suscriber) ReceiveMessages(ctx context.Context, subscription string, collection []in.Event) error {

	log.Println("Service ReceiveMessages")
	return 	s.suscriber.Pull(ctx,subscription,collection)
}