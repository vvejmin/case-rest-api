package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Model
type Promotion struct {
	ID             string `json:"id"`
	Price          string `json:price`
	ExpirationDate string `json:"expiration_date"`
}

// Init promotions var as a slice Promotion struct
var promotions []Promotion

//Get Single Promotion
func getPromotion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Loop through promotions and find with id
	for _, item := range promotions {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Promotion{})
}

func updatePromotion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range promotions {
		if item.ID == params["id"] {
			promotions = append(promotions[:index], promotions[index+1:]...)
			var promotion Promotion
			_ = json.NewDecoder(r.Body).Decode(&promotion)
			//
			promotion.ID = params["id"]
			promotions = append(promotions, promotion)
			json.NewEncoder(w).Encode(promotion)
			return
		}
	}
	json.NewEncoder(w).Encode(promotions)
}

func main() {
	//Init Router
	r := mux.NewRouter().StrictSlash(true)
	// Mock Data - @todo - implement DB
	promotions = append(promotions, Promotion{ID: "172FFC14-D229-4C93-B06B-F48B8C095512", Price: "9.68", ExpirationDate: "2018-06-4 06:01:20"})
	//Roue Handlers / Endpoints
	r.HandleFunc("/promotions/{id}", getPromotion).Methods("GET")
	r.HandleFunc("/promotions/{id}", updatePromotion).Methods("PUT")
	log.Fatal(http.ListenAndServe(":1321", r))

}
