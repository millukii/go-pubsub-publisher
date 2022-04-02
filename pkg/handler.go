package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"publisher/pkg/service"

)
type Handler interface{
	EventHandler(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	publisher service.PublisherService
	service Runner
}
func NewHandler(publisher service.PublisherService, service Runner) Handler {

	return &handler{
		publisher: publisher,
		service: service,
	}
}
func (h handler)EventHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/events/"):]
    w.Write([]byte("The ID is " + id))
		var pr Request
    if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
                http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
                return
    }
		ctx := context.Background()
		if err:=	h.service.Do(ctx, pr,h.publisher); err != nil {
			                http.Error(w, fmt.Sprintf("Error : %v", err), http.StatusInternalServerError)
                return
		}

		w.Write([]byte(" ok"))
		return
}

