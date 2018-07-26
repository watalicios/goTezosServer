package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: The Tezos API written in GO, for easy development.
License: MIT
*/

import (
//  "sync"
  //"flag"
  //"fmt"
  // "log"
  // "html"
  // "net/http"
  "github.com/gorilla/mux"
  "github.com/DefinitelyNotAGoat/goTezosServer"
)

var wg sync.WaitGroup

func main(){
  r := mux.NewRouter()
	r.HandleFunc("/head", GetBlockHead).Methods("GET")
//	r.HandleFunc("/block/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}

}

func GetBlockHead(w http.ResponseWriter, r *http.Request) {
  block, err := goTezosServer.GetBlockHead()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, block)
}
