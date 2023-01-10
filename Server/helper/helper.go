//this package defines custom helper functions to aid in the functioning of the server

package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

// define a function that catches errors by printing error message and panicking to stop execution
// the go-chi Recoverer middleware then recovers the server, logs the error, and sends a 500 Internal Server Error response to the client

func Catch(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

// define a function that sends responses to client reqeusts in JSON format
// the function takes in an interface (meaning any type that implement the interface can be passed) and marshal it into JSON format
// because the interface in this case is empty, it becomes a trick to be able to pass any types into the function
// the http response will have status code as header and the response in the body

func RespondwithJSON(w http.ResponseWriter, statusCode int, anytype interface{}) {
	response, _ := json.Marshal(anytype)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// respondwitherror is just a special type of respondwithjson that specifically takes in and sends an error message

func RespondwithERROR(w http.ResponseWriter, statusCode int, message string) {
	response, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
