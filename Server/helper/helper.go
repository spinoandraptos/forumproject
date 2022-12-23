package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
)

// define a function that helper.Catches errors by printing error message and panicks to stop execution
// the Recoverer middleware then recovers the server, logs the error with a stack trace, and sends a 500 Internal Server Error response to the client
func Catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// define a function that sends responds to client reqeusts in JSON format
// the function takes in an interface (meaning any type that implement the interface can be passed) and marshal it into JSON format
// because the interface in this case is empty, it becomes a trick to be able to pass any types into the function
// the httyp response will have status code as header and the response in the body
func RespondwithJSON(w http.ResponseWriter, statusCode int, anytype interface{}) {
	response, _ := json.Marshal(anytype)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// respondwitherror is just a special type of respondwithjson that specifically takes in a error message
func RespondwithERROR(w http.ResponseWriter, statusCode int, message string) {
	response, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func unathenticatedresponse(w http.ResponseWriter, r *http.Request) {
	token, _, _ := jwtauth.FromContext(r.Context())
	if token == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}
}
