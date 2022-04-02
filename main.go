package main

import (
	"bytes"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"publisher/pkg"
	"publisher/pkg/pubsub"
	"publisher/pkg/service"
	"runtime"
	"runtime/pprof"

	gcp "cloud.google.com/go/pubsub"

)
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var buf bytes.Buffer
var client *gcp.Client
var ctx context.Context
var config service.Configuration


func main() {

	    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

		log.Printf("CPU profile: %s", *cpuprofile)
    // ... rest of the program ...

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
    }
		log.Printf("Memory profile: %s", *memprofile)

		os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	// get Configuration
 config,	err := service.ReadConfig()

	if err != nil {
		panic(err)
	}
	log.Printf("config %v", &config)
		// Wrap context for cancellation, and cancel the first

	ctx = context.Background()
	//Cliente pubsub
	client, err := pubsub.NewPubSubClient(ctx, "sistema")

	if err != nil {
		panic(err)
	}
	// creates all topics
	for i,v := range config.Topics {
	log.Printf("Create topic %d %s", i, v.Topic)
 	err :=	pubsub.CreateTopic(ctx,&buf,client, v.Topic)
		if err != nil {
			log.Printf("%s",err)
		} 
	}
	// creates all subscriptions
	for i,v := range config.Subscriptions {
	log.Printf("Create Subscription %d %s", i, v.Subscription)
	topic :=	pubsub.GetTopic(ctx, &buf,client, v.Topic)
 	err :=	pubsub.CreateSubscription(ctx,&buf,client,v.Subscription, topic)
		if err != nil {
			log.Printf("%s",err)
		} 
	}

// types of events 

	log.Printf("config events %v", &config.Events)


	publisher := pubsub.NewPublisher(client)

	publisherService := service.NewPublisherService(publisher,1)

	runner := pkg.NewRunner()
	handler := pkg.NewHandler(publisherService, runner)

	http.HandleFunc("/events/", func(w http.ResponseWriter, r *http.Request){

		handler.EventHandler(w,r)
	})
	
	  http.ListenAndServe(":8090", nil)
	//		defer client.Close()
}