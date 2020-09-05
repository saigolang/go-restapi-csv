package main

import (
	"encoding/csv"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/getUser", handler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", router))

}

func handler(w http.ResponseWriter, req *http.Request) {

	var finalResult []Customer

	finalResult = getUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalResult)

}

func getUsers() []Customer {

	// opening the csv file
	reader, err := os.Open("customer.csv")

	if err != nil {
		log.Fatal("error in reading from csv file", err.Error())
	}

	records, err := csv.NewReader(reader).ReadAll()

	if err != nil {
		log.Fatal("error in reading the records from csv file", err.Error())
	}

	var finalList []Customer

	for _, r := range records[1:] {
		result := Customer{
			Id:        r[0],
			FirstName: r[1],
			LastName:  r[2],
		}

		finalList = append(finalList, result)

	}

	return finalList
}

// Customer struct
// Fields in the structs should start with upper case inorder to make them included in
// the result
type Customer struct {
	Id        string
	FirstName string
	LastName  string
}
