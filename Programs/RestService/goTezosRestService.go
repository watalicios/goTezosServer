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

func errorControl(err error, w http.ResponseWriter){
  if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func GetBlockHead(w http.ResponseWriter, r *http.Request) {
  block, err := goTezosServer.GetBlockHead()
	errorControl(err, w)
	respondWithJson(w, http.StatusOK, block)
}

func GetBlock(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  block, err := goTezosServer.GetBlock(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, block)
}

func GetBlockProtocol(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  protocol, err := goTezosServer.GetBlockProtocol(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, protocol)
}

func GetBlockChainId(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  blockChainId, err := goTezosServer.GetBlockChainId(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, blockChainId)
}

func GetBlockHash(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  blockHash, err := goTezosServer.GetBlockHash(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, blockHash)
}

func GetBlockHeader(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  blockHeader, err := goTezosServer.GetBlockHeader(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, blockHeader)
}

func GetBlockLevel(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  level, err := goTezosServer.GetBlockHeaderLevel(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, level)
}

func GetBlockProto(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  proto, err := goTezosServer.GetBlockHeaderProto(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, proto)
}

func GetBlockPredecessor(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  predecessor, err := goTezosServer.GetBlockHeaderPredecessor(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, predecessor)
}

func GetBlockTimeStamp(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  timestamp, err := goTezosServer.GetBlockHeaderTimeStamp(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, timestamp)
}

func GetBlockValidationPass(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  validation, err := goTezosServer.GetBlockHeaderValidationPass(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, validation)
}

func GetBlockOperationsHash(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  opHash, err := goTezosServer.GetBlockHeaderOperationsHash(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, opHash)
}

func GetBlockFitness(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  fitness, err := goTezosServer.GetBlockHeaderFitness(blockid)
  errorControl(err, w)
  respondWithJson(w, http.StatusOK, fitness)
}

func GetBlockContext(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  context, err := goTezosServer.GetBlockHeaderContext(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, context)
}

func GetBlockPriority(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  priority, err := goTezosServer.GetBlockHeaderPriority(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, priority)
}

func GetBlockProofOfWorkNonce(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  proofOfWork, err := goTezosServer.GetBlockHeaderProofOfWorkNonce(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, proofOfWork)
}

func GetBlockSignature(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  sig, err := goTezosServer.GetBlockHeaderSignature(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, sig)
}

func GetBlockMetadata(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  meta, err := goTezosServer.GetBlockMetadata(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, meta)
}

func GetBlockMetadataProtocol(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  protocol, err := goTezosServer.GetBlockMetadataProtocol(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, protocol)
}

func GetBlockMetadataNextProtocol(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  protocol, err := goTezosServer.GetBlockMetadataNextProtocol(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, protocol)
}

func GetBlockMetadataTestChainStatus(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  test, err := goTezosServer.GetBlockMetadataTestChainStatus(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, test)
}

func GetBlockMetadataMaxOperationsTTL(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  maxOperationsTTL, err := goTezosServer.GetBlockMetadataMaxOperationsTTL(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, maxOperationsTTL)
}

func GetBlockMetadataMaxOperationDataLength(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  maxOperationDataLength, err := goTezosServer.GetBlockMetadataMaxOperationDataLength(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, maxOperationDataLength)
}

func GetBlockMetadataMaxBlockHeaderLength(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  maxBlockHeaderLength, err := goTezosServer.GetBlockMetadataMaxBlockHeaderLength(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, maxBlockHeaderLength)
}

func GetBlockMetadataMaxOperationListLength(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  maxOperationListLength, err := goTezosServer.GetBlockMetadataMaxOperationListLength(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, maxOperationListLength)
}

func GetBlockMetadataBaker(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  metaBaker, err := goTezosServer.GetBlockMetadataBaker(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, metaBaker)
}

func GetBlockMetadataLevel(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  level, err := goTezosServer.GetBlockMetadataLevel(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, level)
}

func GetBlockMetadataLevelLevel(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  level, err := goTezosServer.GetBlockMetadataLevelLevel(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, level)
}

func GetBlockMetadataLevelLevelPosition(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  levelPosition, err := goTezosServer.GetBlockMetadataLevelLevelPosition(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, levelPosition)
}

func GetBlockMetadataLevelCycle(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  levelCycle, err := goTezosServer.GetBlockMetadataLevelCycle(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, levelCycle)
}

func GetBlockMetadataLevelCyclePosition(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  levelCyclePosition, err := goTezosServer.GetBlockMetadataLevelCyclePosition(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, levelCyclePosition)
}

func GetBlockMetadataLevelVotingPeriod(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  levelVotingPeriod, err := goTezosServer.GetBlockMetadataLevelVotingPeriod(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, levelVotingPeriod)
}

func GetBlockMetadataLevelExpectedCommitment(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  levelExpectedCommitment, err := goTezosServer.GetBlockMetadataLevelExpectedCommitment(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, levelExpectedCommitment)
}

func GetBlockMetadataVotingPeriodKind(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  votingPeriodKind, err := goTezosServer.GetBlockMetadataVotingPeriodKind(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, votingPeriodKind)
}

func GetBlockMetadataNonceHash(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  nonceHash, err := goTezosServer.GetBlockMetadataNonceHash(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, nonceHash)
}

func GetBlockMetadataConsumedGas(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  consumedGas, err := goTezosServer.GetBlockMetadataConsumedGas(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, consumedGas)
}

func GetBlockMetadataDeactivated(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  deactivated, err := goTezosServer.GetBlockMetadataDeactivated(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, deactivated)
}

func GetBlockMetadataBalanceUpdates(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  balanceUpdate, err := goTezosServer.GetBlockMetadataBalanceUpdates(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, balanceUpdate)
}

func GetBlockOperations(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  blockid := parseID(params["id"])
  operations, err := goTezosServer.GetBlockOperations(blockid)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, operations)
}

func GetBlockOperation(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperation(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, operation)
}

func GetBlockOperationProtocol(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperationProtocol(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, operation)
}

func GetBlockOperationsBranch(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperationsBranch(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, operation)
}

func GetBlockOperationsContents(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperationsContents(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, operation)
}

func GetBlockOperationsSignature(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperationsSignature(params["id"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, operation)
}

func GetBlockOperationsByKind(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  operation, err := goTezosServer.GetBlockOperationsByKind(params["id"], params["kind"])
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, operation)
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
