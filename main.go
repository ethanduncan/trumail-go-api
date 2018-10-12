package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	trumail "github.com/sdwolfe32/trumail/verifier"
)

type Email struct {
	Email string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/email", VerifyEmail).Methods("POST")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	v := trumail.NewVerifier("DOMAIN.COM", "test@domain.COM")
	decoder := json.NewDecoder(r.Body)
	var email Email
	err := decoder.Decode(&email)
	if err != nil {
		panic(err)
	}
	val := Must(v.Verify(email.Email))
	log.Println(val)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(val)
}

func Must(l *trumail.Lookup, err error) *trumail.Lookup {
	if err != nil {
		panic(err)
	}
	return l
}
