package main

import (
	"bank-di/db"
	"bank-di/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/register", h.Registration).Methods(http.MethodPost)
	router.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	router.HandleFunc("/create_bill", h.CreateBill).Methods(http.MethodPost)
	router.HandleFunc("/close_bill", h.CloseBill).Methods(http.MethodPost)
	router.HandleFunc("/get_bill_info", h.GetBillInfo).Methods(http.MethodGet)
	//
	router.HandleFunc("/create_card", h.CreateCard).Methods(http.MethodPost)
	//router.HandleFunc("/transfer", h.Transfer).Methods(http.MethodPost)
	router.HandleFunc("/set_limit", h.SetLimit).Methods(http.MethodPost)
	router.HandleFunc("/set_limit", h.SetCardLimit).Methods(http.MethodPost)

	log.Println("API is Running")

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Println(err)
	}

}
