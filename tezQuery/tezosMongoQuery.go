package tezQuery

import (
  "errors"
  "github.com/DefinitelyNotAGoat/goTezosServer"
  "gopkg.in/mgo.v2/bson"
)


func GetBlock(arg interface{}) (goTezosServer.Block, error){
  result := goTezosServer.Block{}
  level := -1
  hash := ""

  switch v := arg.(type) {
    case int:
        level = arg.(int)
    case string:
        hash = arg.(string)
    default:
        return result, errors.New("GetBlock(arg interface{}) failed: Type not Supported")
  }

  if (level > -1){
    err := goTezosServer.Collection.Find(bson.M{"header.level": level}).One(&result)
    if (err != nil) {
  		return result, err
  	}
  }

  if (hash != ""){
    err := goTezosServer.Collection.Find(bson.M{"hash": hash}).One(&result)
    if (err != nil) {
  		return result, err
  	}
  }

  return result, errors.New("GetBlock(arg interface{}) failed: Level and Hash Empty")
}

func GetBlockHead() (goTezosServer.Block, error){  //db.blocks.find().skip(db.blocks.count() - 1)
  result := goTezosServer.Block{}
  Count, err := goTezosServer.Collection.Count()
  if (err != nil){
    return result, err
  }
  err = goTezosServer.Collection.Find(bson.M{}).Skip(goTezosServer.Collection.Count() -1).One(&result)
  if (err != nil) {
		return result, err
	}
  return result, nil
}

// func GetBlockProtocol(arg interface{}) (string, error){
//
// }
//
// func GetBlockChainId(arg interface{}) (string, error){
//
// }
//
// func GetBlockHash(arg interface{}) (string, error) {
//
// }
//
// // func GetBlockHeader(arg interface{}) (goTezosServer.Block.Header, error) {
// //
// // }
//
//
// func GetBlockHeaderLevel(arg interface{}) (int, error) {
//
// }
//
// func GetBlockHeaderProto(arg interface{}) (int, error) {
//
// }
//
// func GetBlockHeaderPredecessor(arg interface{}) (string, error) {
//
// }
//
// func GetBlockHeaderTimeStamp(arg interface{}) (Time.time, error) {
//
// }
//
// func GetBlockHeaderValidationPass(arg interface{}) (int, error) {
//
// }
//
// func GetBlockHeaderOperationsHash(arg interface{}) (string, error) {
//
// }
//
// func GetBlockHeaderFitness(arg interface{}) ([]string, error) {
//
// }
//
// func GetBlockHeaderContext(arg interface{}) ([]string, error) {
//
// }
//
// func GetBlockHeaderPriority(arg interface{}) (int, error) {
//
// }
//
// func GetBlockHeaderProofOfWorkNonce(arg interface{}) (int, error) {
//
// }
//
// func GetBlockHeaderSignature(arg interface{}) (string, error) {
//
// }
//
// // func GetBlockMetadata(arg interface{}) (goTezosServer.Block.Metadata, error) {
// //
// // }
//
// func GetBlockMetadataProtocol(arg interface{}) (string, error) {
//
// }
//
// func GetBlockMetadataNextProtocol(arg interface{}) (string, error) {
//
// }
//
// // func GetBlockMetadataTestChainStatus(arg interface{}) (goTezosServer.Block.Metadata.TestChainStatus, error) {
// //
// // }
//
// func GetBlockMetadataMaxOperationsTTL(arg interface{}) (int, error) {
//
// }
//
// func GetBlockMetadataMaxOperationDataLength(arg interface{}) (int, error) {
//
// }
//
// func GetBlockMetadataMaxBlockHeaderLength(arg interface{}) (int, error) {
//
// }
//
// // func GetBlockMetadataMaxOperationListLength(arg interface{}) ([]goTezosServer.Block.Metadata.MaxOperationListLength, error) {
// //
// // }
//
// func GetBlockMetadataMaxOperationDataLengthMaxSize(arg interface{}) (int, error) {
//
// }
//
// func GetBlockMetadataMaxOperationDataLengthMaxOp(arg interface{}) (int, error) {
//
// }
//
// func GetBlockMetadataBaker(arg interface{}) (string, error) {
//
// }
//
// // func GetBlockMetadataLevel(arg interface{}) ([]goTezosServer.Block.Metadata.Level, error) {
// //
// // }
//
// func GetBlockMetadataLevelLevel(arg interface{}) (int, error) {
//
// }
//
// func GetBlockMetadataLevelPosition(arg interface{}) (int, error) {
//
// }
//
// func GetBlockMetadataLevelCycle(arg interface{}) (int, error) {
//
// }
//
// func GetBlockMetadataLevelCyclePosition(arg interface{}) (int, error) {
//
// }
//
// func GetBlockMetadataLevelVotingPeriod(arg interface{}) (int, error) {
//
// }
//
// func GetBlockMetadataLevelExpectedCommitment(arg interface{}) (bool, error) {
//
// }
//
// func GetBlockMetadataVotingPeriodKind(arg interface{}) (string, error) {
//
// }
//
// func GetBlockMetadataNonceHash(arg interface{}) (interface{}, error) {
//
// }
//
// func GetBlockMetadataConsumedGas(arg interface{}) (string, error) {
//
// }
//
// func GetBlockMetadataDeactivated(arg interface{}) (interface{}, error) {
//
// }
//
// // func GetBlockMetadataBalanceUpdates(arg interface{}) ([]goTezosServer.Block.Metadata.BalanceUpdates, error) {
// //
// // }
//
// func GetBlockMetadataBalanceKind(arg interface{}) (string, error) {
//
// }
//
// func GetBlockMetadataBalanceUpdatesContract(arg interface{}) (string, error) {
//
// }
//
// func GetBlockMetadataBalanceUpdatesChange(arg interface{}) (string, error) {
//
// }
//
// func GetBlockMetadataBalanceUpdatesCategory(arg interface{}) (string, error) {
//
// }
//
// func GetBlockMetadataBalanceUpdatesDelegate(arg interface{}) (string, error) {
//
// }
//
// func GetBlockMetadataBalanceUpdatesLevel(arg interface{}) (int, error) {
//
// }
//
// // func GetBlockOperations(arg interface{}) (goTezosServer.Block.Operations, error){
// //
// // }
//
// func GetBlockOperationsProtocol(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsChainID(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsHash(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsBranch(arg interface{}) (string, error){
//
// }
//
// // func GetBlockOperationsContents(arg interface{}) ([]goTezosServer.Block.Operations.Contents, error){
// //
// // }
//
// func GetBlockOperationsContentsKind(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsContentsLevel(arg interface{}) (int, error){
//
// }
//
// // func GetBlockOperationsContentsMetadata(arg interface{}) (goTezosServer.Block.Operations.Contents.Metadata, error){
// //
// // }
// //
// // func GetBlockOperationsContentsMetadataBalanceUpdates(arg interface{}) ([]goTezosServer.Block.Operations.Contents.Metadata.BalanceUpdates, error){
// //
// // }
//
// func GetBlockOperationsContentsMetadataBalanceUpdatesKind(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsContentsMetadataBalanceUpdatesContract(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsContentsMetadataBalanceUpdatesChange(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsContentsMetadataBalanceUpdatesCategory(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsContentsMetadataBalanceUpdatesDelegate(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsContentsMetadataBalanceUpdatesLevel(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsContentsMetadataDelegate(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsContentsMetadataSlots(arg interface{}) ([]int, error){
//
// }
//
// func GetBlockOperationsSignature(arg interface{}) (string, error){
//
// }
