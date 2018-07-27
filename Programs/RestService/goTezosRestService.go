package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: A ReST API to Query the MongoDB database
License: MIT
*/

import (
  "strconv"
  "net/http"
  "log"
  "time"
  "fmt"
  "flag"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/DefinitelyNotAGoat/goTezosServer"
)

func main(){
  connection := flag.String("connection", "127.0.0.1", "URL or IP to the MongoDB Database")
  db := flag.String("db", "TEZOS", "Use the TEZOS Database")
  collection := flag.String("collection", "blocks", "Use the blocks collection")
  user := flag.String("user", "", "If using authentication, set the user")
  pass := flag.String("pass", "", "If using authentication, set the password")

  flag.Parse()

  var dbCon string

  if (*user != "" && *pass != ""){
    dbCon = "mongodb://" + *user + ":" + *pass + "@" + *connection
  } else{
    dbCon = "mongodb://"+ *connection
  }

  goTezosServer.SetDatabaseConnection(dbCon, *db, *collection)

  r := mux.NewRouter()
	r.HandleFunc("/head/", GetBlockHead).Methods("GET")
	r.HandleFunc("/block/{id}/", GetBlock).Methods("GET")
  r.HandleFunc("/block/{id}/protocol/", GetBlockProtocol).Methods("GET")
  r.HandleFunc("/block/{id}/chain_id/", GetBlockChainId).Methods("GET")
  r.HandleFunc("/block/{id}/hash/", GetBlockHash).Methods("GET")
  r.HandleFunc("/block/{id}/header/", GetBlockHeader).Methods("GET")
  r.HandleFunc("/block/{id}/level/", GetBlockLevel).Methods("GET")
  r.HandleFunc("/block/{id}/proto/", GetBlockProto).Methods("GET")
  r.HandleFunc("/block/{id}/predecessor/", GetBlockPredecessor).Methods("GET")
  r.HandleFunc("/block/{id}/timestamp/", GetBlockTimeStamp).Methods("GET")
  r.HandleFunc("/block/{id}/validation_pass/", GetBlockValidationPass).Methods("GET")
  r.HandleFunc("/block/{id}/operation_hash/", GetBlockOperationsHash).Methods("GET")
  r.HandleFunc("/block/{id}/fitness/", GetBlockFitness).Methods("GET")
  r.HandleFunc("/block/{id}/context/", GetBlockContext).Methods("GET")
  r.HandleFunc("/block/{id}/priority/", GetBlockPriority).Methods("GET")
  r.HandleFunc("/block/{id}/proof_of_work_nonce/", GetBlockProofOfWorkNonce).Methods("GET")
  r.HandleFunc("/block/{id}/signature/", GetBlockSignature).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/", GetBlockMetadata).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/protocol/", GetBlockMetadataProtocol).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/next_protocol/", GetBlockMetadataNextProtocol).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/test_chain_status/", GetBlockMetadataTestChainStatus).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/max_operations_ttl/", GetBlockMetadataMaxOperationsTTL).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/max_operation_data_length/", GetBlockMetadataMaxOperationDataLength).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/max_block_header_length/", GetBlockMetadataMaxBlockHeaderLength).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/max_operation_list_length/", GetBlockMetadataMaxOperationListLength).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/baker/", GetBlockMetadataBaker).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/", GetBlockMetadataLevel).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/level/", GetBlockMetadataLevelLevel).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/position/", GetBlockMetadataLevelLevelPosition).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/cycle/", GetBlockMetadataLevelCycle).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/voting_period/", GetBlockMetadataLevelVotingPeriod).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/expected_commitment/", GetBlockMetadataLevelExpectedCommitment).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/voting_period_kind/", GetBlockMetadataVotingPeriodKind).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/nonce_hash/", GetBlockMetadataNonceHash).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/consumed_gas/", GetBlockMetadataConsumedGas).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/deactivated/", GetBlockMetadataDeactivated).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/balance_updates/", GetBlockMetadataBalanceUpdates).Methods("GET")
  r.HandleFunc("/block/{id}/operations/", GetBlockOperations).Methods("GET")
  r.HandleFunc("/block/operation/{id}/", GetBlockOperation).Methods("GET")
  r.HandleFunc("/block/operation/{id}/protocol/", GetBlockOperationProtocol).Methods("GET")
  r.HandleFunc("/block/operation/{id}/branch/", GetBlockOperationProtocol).Methods("GET")
  r.HandleFunc("/block/operation/{id}/contents/", GetBlockOperationsContents).Methods("GET")
  r.HandleFunc("/block/operation/{id}/signature/", GetBlockOperationsContents).Methods("GET")
  r.HandleFunc("/block/{id}/operations/kind/{kind}/", GetBlockOperationsByKind).Methods("GET")

  r.HandleFunc("/head", GetBlockHead).Methods("GET")
	r.HandleFunc("/block/{id}", GetBlock).Methods("GET")
  r.HandleFunc("/block/{id}/protocol", GetBlockProtocol).Methods("GET")
  r.HandleFunc("/block/{id}/chain_id", GetBlockChainId).Methods("GET")
  r.HandleFunc("/block/{id}/hash", GetBlockHash).Methods("GET")
  r.HandleFunc("/block/{id}/header", GetBlockHeader).Methods("GET")
  r.HandleFunc("/block/{id}/level", GetBlockLevel).Methods("GET")
  r.HandleFunc("/block/{id}/proto", GetBlockProto).Methods("GET")
  r.HandleFunc("/block/{id}/predecessor", GetBlockPredecessor).Methods("GET")
  r.HandleFunc("/block/{id}/timestamp", GetBlockTimeStamp).Methods("GET")
  r.HandleFunc("/block/{id}/validation_pass", GetBlockValidationPass).Methods("GET")
  r.HandleFunc("/block/{id}/operation_hash", GetBlockOperationsHash).Methods("GET")
  r.HandleFunc("/block/{id}/fitness", GetBlockFitness).Methods("GET")
  r.HandleFunc("/block/{id}/context", GetBlockContext).Methods("GET")
  r.HandleFunc("/block/{id}/priority", GetBlockPriority).Methods("GET")
  r.HandleFunc("/block/{id}/proof_of_work_nonce", GetBlockProofOfWorkNonce).Methods("GET")
  r.HandleFunc("/block/{id}/signature", GetBlockSignature).Methods("GET")
  r.HandleFunc("/block/{id}/metadata", GetBlockMetadata).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/protocol", GetBlockMetadataProtocol).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/next_protocol", GetBlockMetadataNextProtocol).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/test_chain_status", GetBlockMetadataTestChainStatus).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/max_operations_ttl", GetBlockMetadataMaxOperationsTTL).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/max_operation_data_length", GetBlockMetadataMaxOperationDataLength).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/max_block_header_length", GetBlockMetadataMaxBlockHeaderLength).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/max_operation_list_length", GetBlockMetadataMaxOperationListLength).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/baker", GetBlockMetadataBaker).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level", GetBlockMetadataLevel).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/level", GetBlockMetadataLevelLevel).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/position", GetBlockMetadataLevelLevelPosition).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/cycle", GetBlockMetadataLevelCycle).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/voting_period", GetBlockMetadataLevelVotingPeriod).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/level/expected_commitment", GetBlockMetadataLevelExpectedCommitment).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/voting_period_kind", GetBlockMetadataVotingPeriodKind).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/nonce_hash", GetBlockMetadataNonceHash).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/consumed_gas", GetBlockMetadataConsumedGas).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/deactivated", GetBlockMetadataDeactivated).Methods("GET")
  r.HandleFunc("/block/{id}/metadata/balance_updates", GetBlockMetadataBalanceUpdates).Methods("GET")
  r.HandleFunc("/block/{id}/operations", GetBlockOperations).Methods("GET")
  r.HandleFunc("/block/operation/{id}", GetBlockOperation).Methods("GET")
  r.HandleFunc("/block/operation/{id}/protocol", GetBlockOperationProtocol).Methods("GET")
  r.HandleFunc("/block/operation/{id}/branch", GetBlockOperationProtocol).Methods("GET")
  r.HandleFunc("/block/operation/{id}/contents", GetBlockOperationsContents).Methods("GET")
  r.HandleFunc("/block/operation/{id}/signature", GetBlockOperationsContents).Methods("GET")
  r.HandleFunc("/block/{id}/operations/kind/{kind}", GetBlockOperationsByKind).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func parseID(id string) interface{}{
  blockid, isInt := strconv.Atoi(id)
  if (isInt != nil){
    return id
  }
  return blockid
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
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  block, err := goTezosServer.GetBlock(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, block)
}

func GetBlockProtocol(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  protocol, err := goTezosServer.GetBlockProtocol(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, protocol)
}

func GetBlockChainId(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  blockChainId, err := goTezosServer.GetBlockChainId(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, blockChainId)
}

func GetBlockHash(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  blockHash, err := goTezosServer.GetBlockHash(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, blockHash)
}

func GetBlockHeader(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  blockHeader, err := goTezosServer.GetBlockHeader(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, rtnBlockHeader)
}

func GetBlockLevel(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  level, err := goTezosServer.GetBlockHeaderLevel(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, level)
}

func GetBlockProto(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  proto, err := goTezosServer.GetBlockHeaderProto(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, proto)
}

func GetBlockPredecessor(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  predecessor, err := goTezosServer.GetBlockHeaderPredecessor(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, predecessor)
}

func GetBlockTimeStamp(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  timestamp, err := goTezosServer.GetBlockHeaderTimeStamp(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, timestamp)
}

func GetBlockValidationPass(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  validation, err := goTezosServer.GetBlockHeaderValidationPass(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, rtnValidation)
}

func GetBlockOperationsHash(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  opHash, err := goTezosServer.GetBlockHeaderOperationsHash(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, opHash)
}

func GetBlockFitness(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  fitness, err := goTezosServer.GetBlockHeaderFitness(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, fitness)
}

func GetBlockContext(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  context, err := goTezosServer.GetBlockHeaderContext(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, context)
}

func GetBlockPriority(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  priority, err := goTezosServer.GetBlockHeaderPriority(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, priority)
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
  var rtnMeta goTezosServer.StructMetadata
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
  respondWithJson(w, http.StatusOK, rtnMeta)
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
  var rtnTest goTezosServer.StructTestChainStatus
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
  var rtnMaxOperationsTTL int
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
  var rtnMaxOperationDataLength int
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
  var rtnMaxBlockHeaderLength int
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
  var rtnMaxOperationListLength []goTezosServer.StructMaxOperationListLength
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
    maxOperationListLength, err := goTezosServer.GetBlockMetadataMaxOperationListLength(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMaxOperationListLength = maxOperationListLength
  }
  respondWithJson(w, http.StatusOK, rtnMaxOperationListLength)
}

func GetBlockMetadataBaker(w http.ResponseWriter, r *http.Request){
  var rtnMetadataBaker string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    metaBaker, err := goTezosServer.GetBlockMetadataBaker(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataBaker = metaBaker
  } else {
    metaBaker, err := goTezosServer.GetBlockMetadataBaker(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataBaker = metaBaker
  }
  respondWithJson(w, http.StatusOK, rtnMetadataBaker)
}

func GetBlockMetadataLevel(w http.ResponseWriter, r *http.Request){
  var rtnMetadataLevel goTezosServer.StructLevel
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    level, err := goTezosServer.GetBlockMetadataLevel(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevel = level
  } else {
    level, err := goTezosServer.GetBlockMetadataLevel(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevel = level
  }
  respondWithJson(w, http.StatusOK, rtnMetadataLevel)
}

func GetBlockMetadataLevelLevel(w http.ResponseWriter, r *http.Request){
  var rtnMetadataLevel int
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    level, err := goTezosServer.GetBlockMetadataLevelLevel(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevel = level
  } else {
    level, err := goTezosServer.GetBlockMetadataLevelLevel(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevel = level
  }
  respondWithJson(w, http.StatusOK, rtnMetadataLevel)
}

func GetBlockMetadataLevelLevelPosition(w http.ResponseWriter, r *http.Request){
  var rtnMetadataLevelPosition int
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    levelPosition, err := goTezosServer.GetBlockMetadataLevelLevelPosition(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevelPosition = levelPosition
  } else {
    levelPosition, err := goTezosServer.GetBlockMetadataLevelLevelPosition(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevelPosition = levelPosition
  }
  respondWithJson(w, http.StatusOK, rtnMetadataLevelPosition)
}

func GetBlockMetadataLevelCycle(w http.ResponseWriter, r *http.Request){
  var rtnMetadataLevelCycle int
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    levelCycle, err := goTezosServer.GetBlockMetadataLevelCycle(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevelCycle = levelCycle
  } else {
    levelCycle, err := goTezosServer.GetBlockMetadataLevelCycle(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevelCycle = levelCycle
  }
  respondWithJson(w, http.StatusOK, rtnMetadataLevelCycle)
}

func GetBlockMetadataLevelCyclePosition(w http.ResponseWriter, r *http.Request){
  var rtnMetadataLevelCyclePosition int
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    levelCyclePosition, err := goTezosServer.GetBlockMetadataLevelCyclePosition(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevelCyclePosition = levelCyclePosition
  } else {
    levelCyclePosition, err := goTezosServer.GetBlockMetadataLevelCyclePosition(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevelCyclePosition = levelCyclePosition
  }
  respondWithJson(w, http.StatusOK, rtnMetadataLevelCyclePosition)
}

func GetBlockMetadataLevelVotingPeriod(w http.ResponseWriter, r *http.Request){
  var rtnMetadataLevelVotingPeriod int
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    levelVotingPeriod, err := goTezosServer.GetBlockMetadataLevelVotingPeriod(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevelVotingPeriod = levelVotingPeriod
  } else {
    levelVotingPeriod, err := goTezosServer.GetBlockMetadataLevelVotingPeriod(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevelVotingPeriod = levelVotingPeriod
  }
  respondWithJson(w, http.StatusOK, rtnMetadataLevelVotingPeriod)
}

func GetBlockMetadataLevelExpectedCommitment(w http.ResponseWriter, r *http.Request){
  var rtnMetadataLevelExpectedCommitment bool
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    levelExpectedCommitment, err := goTezosServer.GetBlockMetadataLevelExpectedCommitment(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevelExpectedCommitment = levelExpectedCommitment
  } else {
    levelExpectedCommitment, err := goTezosServer.GetBlockMetadataLevelExpectedCommitment(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataLevelExpectedCommitment = levelExpectedCommitment
  }
  respondWithJson(w, http.StatusOK, rtnMetadataLevelExpectedCommitment)
}

func GetBlockMetadataVotingPeriodKind(w http.ResponseWriter, r *http.Request){
  var rtnMetadataVotingPeriodKind string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    votingPeriodKind, err := goTezosServer.GetBlockMetadataVotingPeriodKind(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataVotingPeriodKind = votingPeriodKind
  } else {
    votingPeriodKind, err := goTezosServer.GetBlockMetadataVotingPeriodKind(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataVotingPeriodKind = votingPeriodKind
  }
  respondWithJson(w, http.StatusOK, rtnMetadataVotingPeriodKind)
}

func GetBlockMetadataNonceHash(w http.ResponseWriter, r *http.Request){
  var rtnMetadataNonceHash interface{}
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    nonceHash, err := goTezosServer.GetBlockMetadataNonceHash(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataNonceHash = nonceHash
  } else {
    nonceHash, err := goTezosServer.GetBlockMetadataNonceHash(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataNonceHash = nonceHash
  }
  respondWithJson(w, http.StatusOK, rtnMetadataNonceHash)
}

func GetBlockMetadataConsumedGas(w http.ResponseWriter, r *http.Request){
  var rtnMetadataConsumedGas string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    consumedGas, err := goTezosServer.GetBlockMetadataConsumedGas(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataConsumedGas = consumedGas
  } else {
    consumedGas, err := goTezosServer.GetBlockMetadataConsumedGas(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataConsumedGas = consumedGas
  }
  respondWithJson(w, http.StatusOK, rtnMetadataConsumedGas)
}

func GetBlockMetadataDeactivated(w http.ResponseWriter, r *http.Request){
  var rtnMetadataDeactivated []string
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    deactivated, err := goTezosServer.GetBlockMetadataDeactivated(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataDeactivated = deactivated
  } else {
    deactivated, err := goTezosServer.GetBlockMetadataDeactivated(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataDeactivated = deactivated
  }
  respondWithJson(w, http.StatusOK, rtnMetadataDeactivated)
}

func GetBlockMetadataBalanceUpdates(w http.ResponseWriter, r *http.Request){
  var rtnMetadataBalanceUpdates []goTezosServer.StructBalanceUpdates
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    balanceUpdate, err := goTezosServer.GetBlockMetadataBalanceUpdates(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataBalanceUpdates = balanceUpdate
  } else {
    balanceUpdate, err := goTezosServer.GetBlockMetadataBalanceUpdates(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnMetadataBalanceUpdates = balanceUpdate
  }
  respondWithJson(w, http.StatusOK, rtnMetadataBalanceUpdates)
}

func GetBlockOperations(w http.ResponseWriter, r *http.Request){
  var rtnOperations []goTezosServer.StructOperations
  params := mux.Vars(r)
  blockid, isInt := strconv.Atoi(params["id"])
  if (isInt != nil){
    operations, err := goTezosServer.GetBlockOperations(params["id"])
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnOperations = operations
  } else {
    operations, err := goTezosServer.GetBlockOperations(blockid)
    if err != nil {
      respondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    rtnOperations = operations
  }
  respondWithJson(w, http.StatusOK, rtnOperations)
}

func GetBlockOperation(w http.ResponseWriter, r *http.Request){
  var rtnOperation goTezosServer.StructOperations
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperation(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  rtnOperation = operation

  respondWithJson(w, http.StatusOK, rtnOperation)
}

func GetBlockOperationProtocol(w http.ResponseWriter, r *http.Request){
  var rtnOperationProtocol string
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperationProtocol(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  rtnOperationProtocol = operation

  respondWithJson(w, http.StatusOK, rtnOperationProtocol)
}

func GetBlockOperationsBranch(w http.ResponseWriter, r *http.Request){
  var rtnOperationBranch string
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperationsBranch(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  rtnOperationBranch = operation

  respondWithJson(w, http.StatusOK, rtnOperationBranch)
}

func GetBlockOperationsContents(w http.ResponseWriter, r *http.Request){
  var rtnOperationContents []goTezosServer.StructContents
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperationsContents(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  rtnOperationContents = operation

  respondWithJson(w, http.StatusOK, rtnOperationContents)
}

func GetBlockOperationsSignature(w http.ResponseWriter, r *http.Request){
  var rtnOperationSig string
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperationsSignature(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  rtnOperationSig = operation

  respondWithJson(w, http.StatusOK, rtnOperationSig)
}

func GetBlockOperationsByKind(w http.ResponseWriter, r *http.Request){
  var rtnOperationsKind []goTezosServer.StructOperations
  params := mux.Vars(r)
  fmt.Println(params)
  operation, err := goTezosServer.GetBlockOperationsByKind(params["id"], params["kind"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  rtnOperationsKind = operation

  respondWithJson(w, http.StatusOK, rtnOperationsKind)
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
