package pubsub

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
)
type Publisher interface {
		Publish(ctx context.Context, message *pubsub.Message,w io.Writer, topicID string) error
		PublishWithSettings(ctx context.Context, messages []*pubsub.Message,w io.Writer, topicID string) error
		PublishWithFlowControlSettings(ctx context.Context, messages []*pubsub.Message,w io.Writer , topicID string) error
}

type publisher struct{
	client *pubsub.Client
}
func NewPublisher(client *pubsub.Client) Publisher{
	return &publisher{
		client:client,
	}
}

func (p publisher) Publish(ctx context.Context, message *pubsub.Message,w io.Writer, topicID string) error{

	  t := p.client.Topic(topicID)
		result :=	t.Publish(ctx, message)
		
		serverId, err :=	result.Get(ctx)
		if err != nil {
			return err
		}
		log.Printf("Server id %s ", serverId)
		return nil
}
func (p publisher)PublishWithSettings(ctx context.Context, messages []*pubsub.Message,w io.Writer, topicID string) error {

		var results []*pubsub.PublishResult
		var resultErrors []error
        t := p.client.Topic(topicID)
        t.PublishSettings.ByteThreshold = 5000
        t.PublishSettings.CountThreshold = 10
        t.PublishSettings.DelayThreshold = 100 * time.Millisecond

        for i := 0; i < len(messages); i++ {
                result := t.Publish(ctx, messages[i])
                results = append(results, result)
        }
        // The Get method blocks until a server-generated ID or
        // an error is returned for the published message.
        for i, res := range results {
                id, err := res.Get(ctx)
                if err != nil {
                        resultErrors = append(resultErrors, err)
                        fmt.Fprintf(w, "Failed to publish: %v", err)
                        continue
                }
                fmt.Fprintf(w, "Published message %d; msg ID: %v\n", i, id)
        }
        if len(resultErrors) != 0 {
                return fmt.Errorf("Get: %v", resultErrors[len(resultErrors)-1])
        }
        fmt.Fprintf(w, "Published messages with batch settings.")
        return nil
}

func (p publisher) PublishWithFlowControlSettings(ctx context.Context, messages []*pubsub.Message,w io.Writer, topicID string) error {


        t := p.client.Topic(topicID)
        t.PublishSettings.FlowControlSettings = pubsub.FlowControlSettings{
                MaxOutstandingMessages: 100,                     // default 1000
                MaxOutstandingBytes:    10 * 1024 * 1024,        // default 0 (unlimited)
                LimitExceededBehavior:  pubsub.FlowControlBlock, // default Block, other options: SignalError and Ignore
        }

        var wg sync.WaitGroup
        var totalErrors uint64

        numMsgs := 1000
        // Rapidly publishing 1000 messages in a loop may be constrained by flow control.
        for i := 0; i < len(messages); i++ {
                wg.Add(1)
                result := t.Publish(ctx, messages[i])

                go func(i int, res *pubsub.PublishResult) {
                        defer wg.Done()
                        // The Get method blocks until a server-generated ID or
                        // an error is returned for the published message.
                        id, err := res.Get(ctx)
                        if err != nil {
                                // Error handling code can be added here.
                                fmt.Fprintf(w, "Failed to publish: %v", err)
                                atomic.AddUint64(&totalErrors, 1)
                                return
                        }
                        fmt.Fprintf(w, "Published message %d; msg ID: %v\n", i, id)
                }(i, result)
        }

        wg.Wait()

        if totalErrors > 0 {
                return fmt.Errorf("%d of %d messages did not publish successfully", totalErrors, numMsgs)
        }
        return nil
			}