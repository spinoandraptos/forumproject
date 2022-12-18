package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func catch(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
