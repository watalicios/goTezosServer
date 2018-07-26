package main

import (
  "strconv"
  "net/http"
  "log"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/DefinitelyNotAGoat/goTezosServer"
)


func main(){
  r := mux.NewRouter()
	r.HandleFunc("/head", GetBlockHead).Methods("GET")
	r.HandleFunc("/block/{id}", GetBlock).Methods("GET")
  r.HandleFunc("/block/protocol/{id}", GetBlockProtocol).Methods("GET")
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

func GetBlock(w http.ResponseWriter, r *http.Request) {
  var rtbBlock goTezosServer.Block
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    block, err := goTezosServer.GetBlock(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtbBlock = block
  } else {
    block, err := goTezosServer.GetBlock(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtbBlock = block
  }
	respondWithJson(w, http.StatusOK, rtbBlock)
}

func GetBlockProtocol(w http.ResponseWriter, r *http.Request){
  var rtbBlock goTezosServer.Block
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    block, err := goTezosServer.GetBlockProtocol(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtbBlock = block
  } else {
    block, err := goTezosServer.GetBlockProtocol(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtbBlock = block
  }
	respondWithJson(w, http.StatusOK, rtbBlock)
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
