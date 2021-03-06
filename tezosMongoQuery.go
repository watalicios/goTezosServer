package goTezosServer

import (
  "errors"
  "time"
  "fmt"
  "gopkg.in/mgo.v2/bson"
)

func GetBlock(arg interface{}) (Block, error){
  result := Block{}
  level := -1
  hash := ""

  switch arg.(type) {
    case int:
        level = arg.(int)
    case string:
        hash = arg.(string)
    default:
        return result, errors.New("GetBlock(arg interface{}) failed: Type not Supported")
  }

  if (level > -1){
    err := Collection.Find(bson.M{"header.level": level}).One(&result)
    if (err != nil) {
  		return result, err
  	}
  }

  if (hash != ""){
    err := Collection.Find(bson.M{"hash": hash}).One(&result)
    if (err != nil) {
      fmt.Println(err)
  		return result, err
  	}
  }
  //fmt.Println(result)

  return result, nil
}

func GetBlockHead() (Block, error){  //db.blocks.find().skip(db.blocks.count() - 1)
  result := Block{}
  count, err := Collection.Count()
  if (err != nil){
    return result, err
  }
  err = Collection.Find(bson.M{}).Skip(count - 1).One(&result)
  if (err != nil) {
		return result, err
	}
  //fmt.Println(result.Operations)
  return result, nil
}

func GetBlockProtocol(arg interface{}) (string, error){
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Protocol, nil
}

func GetBlockChainId(arg interface{}) (string, error){
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.ChainID, nil
}

func GetBlockHash(arg interface{}) (string, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Hash, nil
}

func GetBlockHeader(arg interface{}) (StructHeader, error) {
  var header StructHeader
  block, err := BlockCheck(arg)
  if (err != nil){
    return header, err
  }
  header.Level, _ = GetBlockHeaderLevel(block)
  header.Proto, _ = GetBlockHeaderProto(block)
  header.Predecessor, _ = GetBlockHeaderPredecessor(block)
  header.Timestamp, _ = GetBlockHeaderTimeStamp(block)
  header.ValidationPass, _ = GetBlockHeaderValidationPass(block)
  header.OperationsHash, _ = GetBlockHeaderOperationsHash(block)
  header.Fitness, _ = GetBlockHeaderFitness(block)
  header.Context, _ = GetBlockHeaderContext(block)
  header.Priority, _ = GetBlockHeaderPriority(block)
  header.ProofOfWorkNonce, _ = GetBlockHeaderProofOfWorkNonce(block)
  header.Signature, _ = GetBlockHeaderSignature(block)

  return header, nil
}

func GetBlockHeaderLevel(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Header.Level, nil
}

func GetBlockHeaderProto(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Header.Proto, nil
}

func GetBlockHeaderPredecessor(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Header.Predecessor, nil
}

func GetBlockHeaderTimeStamp(arg interface{}) (time.Time, error) {
  var t time.Time
  block, err := BlockCheck(arg)
  if (err != nil){
    return t, err
  }
  t = block.Header.Timestamp
  return t, nil
}

func GetBlockHeaderValidationPass(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Header.ValidationPass, nil
}

func GetBlockHeaderOperationsHash(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Header.OperationsHash, nil
}

func GetBlockHeaderFitness(arg interface{}) ([]string, error) {
  var str []string
  block, err := BlockCheck(arg)
  if (err != nil){
    return str, err
  }
  str = block.Header.Fitness
  return str, nil
}

func GetBlockHeaderContext(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Header.Context, nil
}

func GetBlockHeaderPriority(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Header.Priority, nil
}

func GetBlockHeaderProofOfWorkNonce(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Header.ProofOfWorkNonce, nil
}

func GetBlockHeaderSignature(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Header.Signature, nil
}

func GetBlockMetadata(arg interface{}) (StructMetadata, error) {
  var metadata StructMetadata
  block, err := BlockCheck(arg)
  if (err != nil){
    return metadata, err
  }
  metadata.Protocol, _ = GetBlockMetadataProtocol(block)
  metadata.NextProtocol, _ = GetBlockMetadataNextProtocol(block)
  metadata.TestChainStatus, _ = GetBlockMetadataTestChainStatus(block)
  metadata.MaxOperationsTTL, _ = GetBlockMetadataMaxOperationsTTL(block)
  metadata.MaxOperationDataLength, _ = GetBlockMetadataMaxOperationDataLength(block)
  metadata.MaxBlockHeaderLength, _ = GetBlockMetadataMaxBlockHeaderLength(block)
  metadata.MaxOperationListLength, _ = GetBlockMetadataMaxOperationListLength(block)
  metadata.Baker, _ = GetBlockMetadataBaker(block)
  metadata.Level, _ = GetBlockMetadataLevel(block)
  metadata.VotingPeriodKind, _ = GetBlockMetadataVotingPeriodKind(block)
  metadata.NonceHash, _ = GetBlockMetadataNonceHash(block)
  metadata.ConsumedGas, _ = GetBlockMetadataConsumedGas(block)
  metadata.Deactivated, _ = GetBlockMetadataDeactivated(block)
  metadata.BalanceUpdates, _ = GetBlockMetadataBalanceUpdates(block)

  return metadata, nil
}

func GetBlockMetadataProtocol(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.Protocol, nil
}

func GetBlockMetadataNextProtocol(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.NextProtocol, nil
}

func GetBlockMetadataTestChainStatus(arg interface{}) (StructTestChainStatus, error) {
  var testChainStatus StructTestChainStatus
  block, err := BlockCheck(arg)
  if (err != nil){
    return testChainStatus, err
  }
  testChainStatus.Status = block.Metadata.TestChainStatus.Status
  return testChainStatus, nil
}

func GetBlockMetadataMaxOperationsTTL(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.MaxOperationsTTL, nil
}

func GetBlockMetadataMaxOperationDataLength(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.MaxOperationDataLength, nil
}

func GetBlockMetadataMaxBlockHeaderLength(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.MaxBlockHeaderLength, nil
}

func GetBlockMetadataMaxOperationListLength(arg interface{}) ([]StructMaxOperationListLength, error) {
  var maxOperationListLength []StructMaxOperationListLength
  block, err := BlockCheck(arg)
  if (err != nil){
    return maxOperationListLength, err
  }

  for _, field := range block.Metadata.MaxOperationListLength{
    size := field.MaxSize
    op := field.MaxOp
    maxOperationListLength = append(maxOperationListLength, StructMaxOperationListLength{MaxSize: size, MaxOp: op})
  }

  return maxOperationListLength, nil
}

func GetBlockMetadataBaker(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.Baker, nil
}

func GetBlockMetadataLevel(arg interface{}) (StructLevel, error) {
  var level StructLevel
  block, err := BlockCheck(arg)
  if (err != nil){
    return level, err
  }

  level.Level = block.Metadata.Level.Level
  level.LevelPosition = block.Metadata.Level.LevelPosition
  level.Cycle = block.Metadata.Level.Cycle
  level.CyclePosition = block.Metadata.Level.CyclePosition
  level.VotingPeriod =  block.Metadata.Level.VotingPeriod
  level.VotingPeriodPosition = block.Metadata.Level.VotingPeriodPosition
  level.ExpectedCommitment = block.Metadata.Level.ExpectedCommitment

  return level, nil
}

func GetBlockMetadataLevelLevel(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.Level.Level, nil
}

func GetBlockMetadataLevelLevelPosition(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.Level.LevelPosition, nil
}

func GetBlockMetadataLevelCycle(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.Level.Cycle, nil
}

func GetBlockMetadataLevelCyclePosition(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.Level.CyclePosition, nil
}

func GetBlockMetadataLevelVotingPeriod(arg interface{}) (int, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.Level.VotingPeriod, nil
}

func GetBlockMetadataLevelExpectedCommitment(arg interface{}) (bool, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return false, err
  }
  return block.Metadata.Level.ExpectedCommitment, nil
}

func GetBlockMetadataVotingPeriodKind(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.VotingPeriodKind, nil
}

func GetBlockMetadataNonceHash(arg interface{}) (interface{}, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return arg, err
  }
  return block.Metadata.NonceHash, nil
}

func GetBlockMetadataConsumedGas(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.ConsumedGas, nil
}

func GetBlockMetadataDeactivated(arg interface{}) ([]string, error) {
  var interf []string
  block, err := BlockCheck(arg)
  if (err != nil){
    return interf, err
  }
  interf = block.Metadata.Deactivated
  return interf, nil
}

func GetBlockMetadataBalanceUpdates(arg interface{}) ([]StructBalanceUpdates, error) {
  var balanceUpdates []StructBalanceUpdates
  block, err := BlockCheck(arg)
  if (err != nil){
    return balanceUpdates, err
  }
  for _, field := range block.Metadata.BalanceUpdates{
    kind := field.Kind
    contract := field.Contract
    change := field.Change
    category := field.Category
    delegate := field.Delegate
    level := field.Level
    balanceUpdates = append(balanceUpdates, StructBalanceUpdates{Kind: kind, Contract: contract, Change: change, Category: category, Delegate: delegate, Level: level})
  }

  return balanceUpdates, nil
}

func GetBlockOperations(arg interface{}) ([]StructOperations, error){
  var operations []StructOperations
  block, err := BlockCheck(arg)
  if (err != nil){
    return operations, err
  }

  for _, operation := range block.Operations{
    for _, field := range operation{
      protocol := field.Protocol
      chainId := field.ChainID
      hash := field.Hash
      branch := field.Branch
      var contents []StructContents
      for _,cont := range field.Contents{
        contents = append(contents, cont)
      }
      signature := field.Signature
      operations = append(operations, StructOperations{Protocol: protocol, ChainID: chainId, Hash: hash, Contents: contents, Branch: branch, Signature: signature})
    }
  }

  return operations, nil
}

func GetBlockOperation(opHash string) (StructOperations, error){
  var op StructOperations
  block, err := GetBlockByOp(opHash)
  if (err != nil){
    return op, nil
  }

  for _, operation := range block.Operations{
    for _, field := range operation{
      if (opHash == field.Hash) {
        protocol := field.Protocol
        chainId := field.ChainID
        branch := field.Branch
        var contents []StructContents
        for _,cont := range field.Contents{
          contents = append(contents, cont)
        }
        signature := field.Signature
        op = StructOperations{Protocol: protocol, ChainID: chainId, Hash: opHash, Contents: contents, Branch: branch, Signature: signature}
        break
      }
    }
    break
  }
  return op, nil
}

func GetBlockOperationProtocol(opHash string) (string, error){
  operation, err := GetBlockOperation(opHash)
  if (err != nil){
    return "", err
  }
  return operation.Protocol, nil
}

func GetBlockOperationsBranch(opHash string) (string, error){
  operation, err := GetBlockOperation(opHash)
  if (err != nil){
    return "", err
  }
  return operation.Branch, nil
}

func GetBlockOperationsContents(opHash string) ([]StructContents, error){
  operation, err := GetBlockOperation(opHash)
  if (err != nil){
    return operation.Contents, err
  }

  return operation.Contents, nil
}

func GetBlockOperationsByKind(arg interface{}, kind string) ([]StructOperations, error){
  operations, err := GetBlockOperations(arg)
  if (err != nil){
    return operations, err
  }
  var ops []StructOperations
  for _, op := range operations{
    for _, cont := range op.Contents{
      if (cont.Kind == kind){
        ops = append(ops, op)
      }
    }
  }

  return ops, nil
}

func GetBlockOperationsSignature(opHash string) (string, error){
  operation, err := GetBlockOperation(opHash)
  if (err != nil){
    return "", err
  }

  return operation.Signature, nil
}

func BlockCheck(arg interface{}) (Block, error){
  var block Block
  var err error

  switch arg.(type) {
    case Block:
      block = arg.(Block)
    default:
      block, err = GetBlock(arg)
      if (err != nil){
        return block, err
      }
  }

  return block, nil
}

func GetBlockByOp(opHash string) (Block, error) {
  result := Block{}
  err := Collection.Find(bson.M{"operations":bson.M{"$elemMatch":bson.M{"$elemMatch":bson.M{"hash":opHash}}}}).One(&result)
  if (err != nil) {
		return result, err
	}
  return result, nil
}
