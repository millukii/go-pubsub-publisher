package in

import "time"

type Event struct {
	Data            []byte
	Attributes      map[string]string
	DeliveryAttempt *int
	PublishTime     time.Time
	ConsumedTime time.Time
}