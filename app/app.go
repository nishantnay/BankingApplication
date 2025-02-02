package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nishantnay/bankingapplication/handler"
)

func Start() {

	router:=mux.NewRouter()

	router.HandleFunc("/hello", handler.Hello)
	router.HandleFunc("/getCust", handler.GetCustomer)
	router.HandleFunc("/api/time",handler.GetTime)

	http.ListenAndServe("localhost:8000", router)

}
