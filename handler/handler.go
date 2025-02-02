package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Customer struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Pincode int
}

type TimeDataError struct {
	ErrorMessage string `json:"error_message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Word")
}

func GetTime(w http.ResponseWriter, r *http.Request) {
	timeZone := r.URL.Query().Get("tz")
	zonesRequested := strings.Split(timeZone, ",")
	w.Header().Add("content-type", "application/json")
	myResponse:= make(map[string] string)
	for _, mytimeZone := range zonesRequested {
		time_Location, err := time.LoadLocation(mytimeZone)
		if err != nil {
			json.NewEncoder(w).Encode(TimeDataError{
				ErrorMessage: fmt.Sprintf("Please provide valid location to get a time zone for %s", mytimeZone),
			})
			return
		}
		time_inLocation := time.Now().In(time_Location)
		if mytimeZone==""{
			mytimeZone="UTC"
		}
		myResponse[mytimeZone]=time_inLocation.String()
		
	}

	json.NewEncoder(w).Encode(myResponse)

}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	cust := []Customer{{
		Name:    "Nishant",
		Age:     34,
		Pincode: 412207,
	},
		{
			Name:    "Nayanam",
			Age:     34,
			Pincode: 412207,
		}}
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(cust)
	//xml.NewEncoder(w).Encode(cust)

}
