package server

import (
	"encoding/json"
	"fmt"
	"github.com/owulveryck/CSAaaS/subscription"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")

}

func CsaCreate(w http.ResponseWriter, r *http.Request) {
	var s map[string]interface{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)

	}
	if err := r.Body.Close(); err != nil {
		panic(err)

	}
	if err := json.Unmarshal(body, &s); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	// Create the subscription request
	//Generate the uuid
	u1 := uuid.NewV4()
	var sr subscription.SubscriptionRequest
	sd := Subscriptions[fmt.Sprintf("%v/%v", r.URL, r.Method)]
	sr.CatalogID = sd.CatalogID
	sr.CategoryName = sd.Category.Name
	sr.SubscriptionName = fmt.Sprintf("Request %s generated with API", u1)
	sr.SubscriptionDescription = "Send by API..."
	sr.Action = "ORDER"
	sr.StartDate = subscription.JSONTime(time.Now())
	sr.EndDate = subscription.JSONTime(time.Now().Add(8760 * time.Hour))
	sr.Fields = make(map[string]interface{}, len(sd.Fields))
	for _, field := range sd.Fields {
		if val, ok := s[field.Name]; ok {
			sr.Fields[field.ID] = val
		} else {
			sr.Fields[field.ID] = field.Value
		}
	}
	o, _ := json.MarshalIndent(sr, " ", "  ")
	log.Printf("Sending this to CSA: %s", o)
	type response struct {
		Id string `json:"ID"`
	}
	t := response{
		Id: u1.String(),
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)

	}
}

func GetServices(w http.ResponseWriter, r *http.Request) {
	log.Println("GetServices")
	if err := json.NewEncoder(w).Encode(Subscriptions); err != nil {
		panic(err)
	}

}
func CsaGet(w http.ResponseWriter, r *http.Request) {
}
func CsaDelete(w http.ResponseWriter, r *http.Request) {
}
