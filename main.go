package main

import (
	"fmt"
	//"m-hospital/calculation"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Item struct {
	UID string `json:"UID"`
	Name string`json:"Name"`
	Desc string `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

/*
func Multiply(num1 int, num2 int) int {
	return mul(num1, num2)
}

func mul(num1 int, num2 int) int {
	return num1 * num2
}
*/

//Batasan endpoint
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homePage()")
}

//Memanggil http port :8000
func handlerRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/",homePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

//Halaman Utama
func main() {
	handlerRequest()
}
