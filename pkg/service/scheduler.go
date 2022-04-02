package service

import (
	"context"
	"log"
	"publisher/pkg/model/in"

)
type Service struct {
	Subscribers []SubscriberService
	Publishers  []PublisherService
}

type SchedulerService interface {
	Start()
}

type scheduler struct {
	Svc    Service
	config *Configuration
}

func NewSchedulerService(svc Service, config *Configuration) SchedulerService {
	return &scheduler{
		Svc: svc,
		config: config,
	}
}
func (s *scheduler) Start() {
	log.Printf("Subscribers  %v", s.Svc.Subscribers)

	for i, svc := range s.Svc.Subscribers {
	log.Print("subscriber process ", i)

	log.Print("subscriber process go routine ", i)
			ctx := context.Background()
			//input

		events := []map[string]interface{}{
		{
			"data": map[string]interface{}{
				"id": "1",
			},
			"attributes": map[string]string{"eventType":"creation"},
		},
	}
		log.Printf("%s", events)
		err := s.Svc.Publishers[i].PublishMessages(ctx, events,s.config.Topics[i].Topic)

		if err != nil {
			log.Printf("%s",err)
		} 

		collection := []in.Event{}
		err =	svc.ReceiveMessages(ctx, s.config.Subscriptions[i].Subscription, collection)
		log.Printf("Events consumed %+v",collection)
		if err != nil {
					log.Printf("%s",err)
		} 
		log.Printf("Filters %+v",s.config.Events[0].Filters )
		//mapper 
		for _, event := range collection {
			//filters
			for _, configuredEvent := range s.config.Events {
				valid := false
				log.Println(configuredEvent.Filters.Attributes.EventType)
				if (configuredEvent.Filters.Attributes.EventType !="" &&
				 event.Attributes["eventType"]==configuredEvent.Filters.Attributes.EventType){

						if (configuredEvent.Filters.Attributes.EventName !="" && event.Attributes["eventName"]==configuredEvent.Filters.Attributes.EventName){
							valid = true
						}
				}
				if (valid){
					// load mapper
					log.Print("mapper")
					// publish output message
				}
			}

		}

	}
}