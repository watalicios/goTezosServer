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
  r.HandleFunc("/block/chainid/{id}", GetBlockChainId).Methods("GET")
  r.HandleFunc("/block/hash/{id}", GetBlockHash).Methods("GET")
  r.HandleFunc("/block/hash/header/{id}", GetBlockHeader).Methods("GET")
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
  var rtnProtocol string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    protocol, err := goTezosServer.GetBlockProtocol(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnProtocol = protocol
  } else {
    protocol, err := goTezosServer.GetBlockProtocol(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnProtocol = protocol
  }
	respondWithJson(w, http.StatusOK, rtnProtocol)
}

func GetBlockChainId(w http.ResponseWriter, r *http.Request){
  var rtnBlockChainId string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    blockChainId, err := goTezosServer.GetBlockChainId(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnBlockChainId = blockChainId
  } else {
    blockChainId, err := goTezosServer.GetBlockChainId(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnBlockChainId = blockChainId
  }
	respondWithJson(w, http.StatusOK, rtnBlockChainId)
}

func GetBlockHash(w http.ResponseWriter, r *http.Request){
  var rtnBlockHash string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    blockHash, err := goTezosServer.GetBlockHash(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnBlockHash = blockHash
  } else {
    blockHash, err := goTezosServer.GetBlockHash(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnBlockHash = blockHash
  }
	respondWithJson(w, http.StatusOK, rtnBlockHash)
}

func GetBlockHeader(w http.ResponseWriter, r *http.Request){
  var rtnBlockHeader goTezosServer.StructHeader
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    blockHeader, err := goTezosServer.GetBlockHash(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnBlockHeader = blockHeader
  } else {
    blockHash, err := goTezosServer.GetBlockHash(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnBlockHeader = blockHeader
  }
	respondWithJson(w, http.StatusOK, rtnBlockHeader)
}

// func CheckType(v interface{}) (int, error) {
//   switch v.(type){
//   case int:
//   }
// }

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
