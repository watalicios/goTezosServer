package main

import (
  "strconv"
  "net/http"
  "log"
  "time"
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
  r.HandleFunc("/block/header/{id}", GetBlockHeader).Methods("GET")
  r.HandleFunc("/block/level/{id}", GetBlockLevel).Methods("GET")
  r.HandleFunc("/block/proto/{id}", GetBlockProto).Methods("GET")
  r.HandleFunc("/block/predecessor/{id}", GetBlockPredecessor).Methods("GET")
  r.HandleFunc("/block/timestamp/{id}", GetBlockTimeStamp).Methods("GET")
  r.HandleFunc("/block/validationpass/{id}", GetBlockValidationPass).Methods("GET")
  r.HandleFunc("/block/operationhash/{id}", GetBlockOperationsHash).Methods("GET")
  r.HandleFunc("/block/fitness/{id}", GetBlockFitness).Methods("GET")
  r.HandleFunc("/block/context/{id}", GetBlockContext).Methods("GET")
  r.HandleFunc("/block/priority/{id}", GetBlockPriority).Methods("GET")
  r.HandleFunc("/block/proofofworknonce/{id}", GetBlockProofOfWorkNonce).Methods("GET")
  r.HandleFunc("/block/signature/{id}", GetBlockSignature).Methods("GET")
  r.HandleFunc("/block/metadata/{id}", GetBlockMetadata).Methods("GET")
  r.HandleFunc("/block/metadata/protocol/{id}", GetBlockMetadataProtocol).Methods("GET")
  r.HandleFunc("/block/metadata/nextprotocol/{id}", GetBlockMetadataNextProtocol).Methods("GET")
  r.HandleFunc("/block/metadata/testchainstatus/{id}", GetBlockMetadataTestChainStatus).Methods("GET")
  r.HandleFunc("/block/metadata/maxoperationsttl/{id}", GetBlockMetadataMaxOperationsTTL).Methods("GET")
  r.HandleFunc("/block/metadata/maxoperationdatalength/{id}", GetBlockMetadataMaxOperationDataLength).Methods("GET")
  r.HandleFunc("/block/metadata/maxblockheaderlength/{id}", GetBlockMetadataMaxBlockHeaderLength).Methods("GET")
  r.HandleFunc("/block/metadata/maxoperationlistlength/{id}", GetBlockMetadataMaxOperationListLength).Methods("GET")
  //GetBlockMetadataMaxOperationListLength
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
    blockHeader, err := goTezosServer.GetBlockHeader(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnBlockHeader = blockHeader
  } else {
    blockHeader, err := goTezosServer.GetBlockHeader(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnBlockHeader = blockHeader
  }
	respondWithJson(w, http.StatusOK, rtnBlockHeader)
}

func GetBlockLevel(w http.ResponseWriter, r *http.Request){
  var rtnLevel int
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    level, err := goTezosServer.GetBlockHeaderLevel(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnLevel = level
  } else {
    level, err := goTezosServer.GetBlockHeaderLevel(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnLevel = level
  }
	respondWithJson(w, http.StatusOK, rtnLevel)
}

func GetBlockProto(w http.ResponseWriter, r *http.Request){
  var rtnProto int
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    proto, err := goTezosServer.GetBlockHeaderProto(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnProto = proto
  } else {
    proto, err := goTezosServer.GetBlockHeaderProto(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnProto = proto
  }
	respondWithJson(w, http.StatusOK, rtnProto)
}

func GetBlockPredecessor(w http.ResponseWriter, r *http.Request){
  var rtnPredecessor string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    predecessor, err := goTezosServer.GetBlockHeaderPredecessor(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnPredecessor = predecessor
  } else {
    predecessor, err := goTezosServer.GetBlockHeaderPredecessor(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnPredecessor = predecessor
  }
	respondWithJson(w, http.StatusOK, rtnPredecessor)
}

func GetBlockTimeStamp(w http.ResponseWriter, r *http.Request){
  var rtnTimestamp time.Time
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    timestamp, err := goTezosServer.GetBlockHeaderTimeStamp(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnTimestamp = timestamp
  } else {
    timestamp, err := goTezosServer.GetBlockHeaderTimeStamp(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnTimestamp = timestamp
  }
	respondWithJson(w, http.StatusOK, rtnTimestamp)
}

func GetBlockValidationPass(w http.ResponseWriter, r *http.Request){
  var rtnValidation int
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    validation, err := goTezosServer.GetBlockHeaderValidationPass(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnValidation = validation
  } else {
    validation, err := goTezosServer.GetBlockHeaderValidationPass(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnValidation = validation
  }
	respondWithJson(w, http.StatusOK, rtnValidation)
}

func GetBlockOperationsHash(w http.ResponseWriter, r *http.Request){
  var rtnOpHash string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    opHash, err := goTezosServer.GetBlockHeaderOperationsHash(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnOpHash = opHash
  } else {
    opHash, err := goTezosServer.GetBlockHeaderOperationsHash(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnOpHash = opHash
  }
	respondWithJson(w, http.StatusOK, rtnOpHash)
}

func GetBlockFitness(w http.ResponseWriter, r *http.Request){
  var rtnFitness []string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    fitness, err := goTezosServer.GetBlockHeaderFitness(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnFitness = fitness
  } else {
    fitness, err := goTezosServer.GetBlockHeaderFitness(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnFitness = fitness
  }
	respondWithJson(w, http.StatusOK, rtnFitness)
}

func GetBlockContext(w http.ResponseWriter, r *http.Request){
  var rtnContext string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    context, err := goTezosServer.GetBlockHeaderContext(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnContext = context
  } else {
    context, err := goTezosServer.GetBlockHeaderContext(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnContext = context
  }
	respondWithJson(w, http.StatusOK, rtnContext)
}

func GetBlockPriority(w http.ResponseWriter, r *http.Request){
  var rtnPriority int
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    priority, err := goTezosServer.GetBlockHeaderPriority(params["id"])
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnPriority = priority
  } else {
    priority, err := goTezosServer.GetBlockHeaderPriority(blockid)
  	if err != nil {
  		respondWithError(w, http.StatusInternalServerError, err.Error())
  		return
  	}
    rtnPriority = priority
  }
	respondWithJson(w, http.StatusOK, rtnPriority)
}

func GetBlockProofOfWorkNonce(w http.ResponseWriter, r *http.Request){
  var rtnProofOfWork string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    proofOfWork, err := goTezosServer.GetBlockHeaderProofOfWorkNonce(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnProofOfWork = proofOfWork
  } else {
    proofOfWork, err := goTezosServer.GetBlockHeaderProofOfWorkNonce(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnProofOfWork = proofOfWork
  }
  respondWithJson(w, http.StatusOK, rtnProofOfWork)
}

func GetBlockSignature(w http.ResponseWriter, r *http.Request){
  var rtnSig string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    sig, err := goTezosServer.GetBlockHeaderSignature(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnSig = sig
  } else {
    sig, err := goTezosServer.GetBlockHeaderSignature(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnSig = sig
  }
  respondWithJson(w, http.StatusOK, rtnSig)
}

func GetBlockMetadata(w http.ResponseWriter, r *http.Request){
  var rtnMeta StructMetadata
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    meta, err := goTezosServer.GetBlockMetadata(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMeta = meta
  } else {
    meta, err := goTezosServer.GetBlockMetadata(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMeta = meta
  }
  respondWithJson(w, http.StatusOK, rtnSig)
}

func GetBlockMetadataProtocol(w http.ResponseWriter, r *http.Request){
  var rtnProtocol string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    protocol, err := goTezosServer.GetBlockMetadataProtocol(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnProtocol = protocol
  } else {
    protocol, err := goTezosServer.GetBlockMetadataProtocol(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnProtocol = protocol
  }
  respondWithJson(w, http.StatusOK, rtnProtocol)
}

func GetBlockMetadataNextProtocol(w http.ResponseWriter, r *http.Request){
  var rtnProtocol string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    protocol, err := goTezosServer.GetBlockMetadataNextProtocol(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnProtocol = protocol
  } else {
    protocol, err := goTezosServer.GetBlockMetadataNextProtocol(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnProtocol = protocol
  }
  respondWithJson(w, http.StatusOK, rtnProtocol)
}

func GetBlockMetadataTestChainStatus(w http.ResponseWriter, r *http.Request){
  var rtnTest StructTestChainStatus
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    test, err := goTezosServer.GetBlockMetadataTestChainStatus(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnTest = test
  } else {
    test, err := goTezosServer.GetBlockMetadataTestChainStatus(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnTest = test
  }
  respondWithJson(w, http.StatusOK, rtnTest)
}

func GetBlockMetadataMaxOperationsTTL(w http.ResponseWriter, r *http.Request){
  var rtnMaxOperationsTTL StructTestChainStatus
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    maxOperationsTTL, err := goTezosServer.GetBlockMetadataMaxOperationsTTL(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMaxOperationsTTL = maxOperationsTTL
  } else {
    maxOperationsTTL, err := goTezosServer.GetBlockMetadataMaxOperationsTTL(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMaxOperationsTTL = maxOperationsTTL
  }
  respondWithJson(w, http.StatusOK, rtnMaxOperationsTTL)
}

func GetBlockMetadataMaxOperationDataLength(w http.ResponseWriter, r *http.Request){
  var rtnMaxOperationDataLength StructTestChainStatus
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    maxOperationDataLength, err := goTezosServer.GetBlockMetadataMaxOperationDataLength(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMaxOperationDataLength = maxOperationDataLength
  } else {
    maxOperationDataLength, err := goTezosServer.GetBlockMetadataMaxOperationDataLength(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMaxOperationDataLength = maxOperationDataLength
  }
  respondWithJson(w, http.StatusOK, rtnMaxOperationDataLength)
}

func GetBlockMetadataMaxBlockHeaderLength(w http.ResponseWriter, r *http.Request){
  var rtnMaxBlockHeaderLength StructTestChainStatus
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    maxBlockHeaderLength, err := goTezosServer.GetBlockMetadataMaxBlockHeaderLength(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMaxBlockHeaderLength = maxBlockHeaderLength
  } else {
    maxBlockHeaderLength, err := goTezosServer.GetBlockMetadataMaxBlockHeaderLength(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMaxBlockHeaderLength = maxBlockHeaderLength
  }
  respondWithJson(w, http.StatusOK, rtnMaxBlockHeaderLength)
}

func GetBlockMetadataMaxOperationListLength(w http.ResponseWriter, r *http.Request){
  var rtnMaxOperationListLength StructTestChainStatus
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    maxOperationListLength, err := goTezosServer.GetBlockMetadataMaxOperationListLength(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMaxOperationListLength = maxOperationListLength
  } else {
    maxBlockHeaderLength, err := goTezosServer.GetBlockMetadataMaxOperationListLength(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMaxOperationListLength = maxOperationListLength
  }
  respondWithJson(w, http.StatusOK, rtnMaxOperationListLength)
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
