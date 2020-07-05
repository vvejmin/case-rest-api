package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type promotion struct {
	id              string `json:"id"`
	price           string `json:price`
	expiration_date string `json:"expiration_date"`
}

type allPromotions []promotion

var promotions = allPromotions{

	{
		id:              "172FFC14-D229-4C93-B06B-F48B8C095512",
		price:           "9.68",
		expiration_date: "2018-06-4 06:01:20",
	},
}

func getOnePromotion(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singlePromotion := range promotions {
		if singlePromotion.id == eventID {
			json.NewEncoder(w).Encode(singlePromotion)
		}
	}
}

func updatePromotion(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedPromotion promotion

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Not Found")
	}
	json.Unmarshal(reqBody, &updatedPromotion)

	for i, singleEvent := range promotions {
		if singleEvent.id == eventID {
			promotions = append(promotions[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}

}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/promotions/{id}", getOnePromotion).Methods("GET")
	r.HandleFunc("/promotions/{id}", updatePromotion).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":1321", r))

}
