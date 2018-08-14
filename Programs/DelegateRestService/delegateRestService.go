package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: A ReST API to Query the MongoDB database
License: MIT
*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/DefinitelyNotAGoat/goTezosServer"
	"github.com/gorilla/mux"
)

var (
	httpP  string
	httpsP string
)

func main() {
	connection := flag.String("connection", "127.0.0.1", "URL or IP to the MongoDB Database")
	db := flag.String("db", "TEZOS", "Use the TEZOS Database")
	collection := flag.String("collection", "delegates", "Use the blocks collection")
	user := flag.String("user", "", "If using authentication, set the user")
	pass := flag.String("pass", "", "If using authentication, set the password")
	httpPort := flag.String("httpport", "3000", "The HTTP port to listen on, will be redirected to HTTPS")
	httpsPort := flag.String("httpsport", "3001", "The HTTPs port to listen on")
	cert := flag.String("cert", "./cert.pem", "TLS certificate.")
	key := flag.String("key", "./key.pem", "TLS key.")

	flag.Parse()

	var dbCon string

	if *user != "" && *pass != "" {
		dbCon = "mongodb://" + *user + ":" + *pass + "@" + *connection
	} else {
		dbCon = "mongodb://" + *connection
	}

	goTezosServer.SetDatabaseConnection(dbCon, *db, *collection)

	httpP = ":" + *httpPort
	httpsP = ":" + *httpsPort

	go http.ListenAndServe(httpP, http.HandlerFunc(redirect))

	r := mux.NewRouter()
	r.HandleFunc("/delegate/{id}/", GetDelegate).Methods("GET")
	r.HandleFunc("/delegate/{id}", GetDelegate).Methods("GET")
	r.HandleFunc("/delegate/{id}/report/rate/{rate}/cycles/{cycles}/", GetDelegateReport).Methods("GET")
	r.HandleFunc("/delegate/{id}/report/rate/{rate}/cycles/{cycles}", GetDelegateReport).Methods("GET")

	err := http.ListenAndServeTLS(":3001", *cert, *key, r)
	if err != nil {
		fmt.Println(err)
	}
}

func GetDelegate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	delegateAddr := params["id"]
	delegate, err := goTezosServer.GetDelegateInfo(delegateAddr)
	errorControl(err, w)
	respondWithJson(w, http.StatusOK, delegate)
}

func GetDelegateReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	delegateAddr := params["id"]
	rate, err := strconv.ParseFloat(params["rate"], 64)
	errorControl(err, w)
	cycles := params["cycles"]
	delegate, err := goTezosServer.ComputeDelegateServiceReport(rate, delegateAddr, cycles)
	errorControl(err, w)
	respondWithJson(w, http.StatusOK, delegate)
}

func errorControl(err error, w http.ResponseWriter) {
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	oldhost := req.Host
	newhost := strings.Replace(oldhost, httpP, httpsP, -1)

	target := "https://" + newhost + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see @andreiavrammsd comment: often 307 > 301
		http.StatusTemporaryRedirect)
}
