package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-pascal/iban"
)

const applicationPort = "8080"

func main() {
	http.HandleFunc("/validate", IBANValidationHandler)
	if err := http.ListenAndServe(":"+applicationPort, nil); err != nil {
		log.Fatalf("Couldn't run server. Got error: %s", err.Error())
	}
}

//Response for request response
type Response struct {
	Status int         `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func sendResp(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Could not encode %+v to json. Got error: %s", resp, err.Error()))
	}
}

func IBANValidationHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{}
	IBAN := r.FormValue("iban")
	ok, _, err := iban.IsCorrectIban(IBAN, true)
	if err != nil || !ok {
		resp.Status = http.StatusBadRequest
		resp.Error = fmt.Sprintf("Given IBAN is not valid: %v", IBAN)
		sendResp(w, resp, resp.Status)
		return
	}
	resp.Data = "Given IBAN is valid"
	resp.Status = http.StatusOK
	sendResp(w, resp, resp.Status)
	return
}
