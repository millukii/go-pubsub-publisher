package pkg

import (
	"context"
	"encoding/base64"
	"log"
	"publisher/pkg/service"
)

type Runner interface {
	Do(ctx context.Context,req Request,publisher service.PublisherService) error
}
type runner struct{}

func NewRunner() Runner {
	return &runner{}
}
func (r *runner) Do(ctx context.Context,req Request, publisher service.PublisherService) error {

	log.Printf("%+v", req)
  raw, err := base64.StdEncoding.DecodeString(req.Message.Data)
    if err != nil {
        log.Fatalf("Base64: %v", err)
        return err
    }

	return 	publisher.Publish(ctx, raw, req.Message.Attributes, "topic-1")
}
