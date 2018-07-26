package random

import (
  "errors"
  "time"
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
  		return result, err
  	}
  }

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

func GetBlockHeader(arg interface{}) (Header, error) {
  var header Header
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
  return block.Header.ProofOfWorkNonce, nil
}

func GetBlockMetadata(arg interface{}) (Metadata, error) {
  var metadata Metadata
  block, err := BlockCheck(arg)
  if (err != nil){
    return metadata, err
  }
  metadata.Protocol, _ = GetBlockMetadataProtocol(block)
  metadata.NextProtocol, _ = GetBlockMetadataNextProtocol(block)
  metadata.TstChainStatus, _ = GetBlockMetadataTestChainStatus(block)
  metadata.MaxOperationsTTL, _ = GetBlockMetadataMaxOperationsTTL(block)
  metadata.MaxOperationDataLength, _ = GetBlockMetadataMaxOperationDataLength(block)
  metadata.MaxBlockHeaderLength, _ = GetBlockMetadataMaxBlockHeaderLength(block)
  metadata.MxOperationListLength, _ = GetBlockMetadataMaxOperationListLength(block)
  metadata.Baker, _ = GetBlockMetadataBaker(block)
  metadata.Lvl, _ = GetBlockMetadataLevel(block)
  metadata.VotingPeriodKind, _ = GetBlockMetadataVotingPeriodKind(block)
  metadata.NonceHash, _ = GetBlockMetadataNonceHash(block)
  metadata.ConsumedGas, _ = GetBlockMetadataConsumedGas(block)
  metadata.Deactivated, _ = GetBlockMetadataDeactivated(block)
  metadata.BlncUpdates, _ = GetBlockMetadataBalanceUpdates(block)

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

func GetBlockMetadataTestChainStatus(arg interface{}) (TestChainStatus, error) {
  var testChainStatus TestChainStatus
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

func GetBlockMetadataMaxOperationListLength(arg interface{}) ([]MaxOperationListLength, error) {
  var maxOperationListLength []MaxOperationListLength
  block, err := BlockCheck(arg)
  if (err != nil){
    return maxOperationListLength, err
  }

  for _, field := range block.Metadata.MaxOperationListLength{
    size := field.MaxSize
    op := field.MaxOp
    maxOperationListLength = append(maxOperationListLength, MaxOperationListLength{MaxSize: size, MaxOp: op})
  }

  return maxOperationListLength, nil
}

// func GetBlockMetadataMaxOperationDataLengthMaxSize(arg interface{}) (int, error) {
//   block, err := GetBlock(arg)
//   if (err != nil){
//     return 0, err
//   }
//   return block.Metadata.MaxOperationListLength.MaxSize, nil
// }
//
// func GetBlockMetadataMaxOperationDataLengthMaxOp(arg interface{}) (int, error) {
//   block, err := GetBlock(arg)
//   if (err != nil){
//     return 0, err
//   }
//   return block.Metadata.MaxOperationListLength.MaxOP, nil
// }

func GetBlockMetadataBaker(arg interface{}) (string, error) {
  block, err := BlockCheck(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.Baker, nil
}

func GetBlockMetadataLevel(arg interface{}) (Level, error) {
  var level Level
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

func GetBlockMetadataDeactivated(arg interface{}) ([]interface{}, error) {
  var interf []interface{}
  block, err := BlockCheck(arg)
  if (err != nil){
    return interf, err
  }
  interf = block.Metadata.Deactivated
  return interf, nil
}

func GetBlockMetadataBalanceUpdates(arg interface{}) ([]BalanceUpdates, error) {
  var balanceUpdates []BalanceUpdates
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
    balanceUpdates = append(balanceUpdates, BalanceUpdates{Kind: kind, Contract: contract, Change: change, Category: category, Delegate: delegate, Level: level})
  }

  return balanceUpdates, nil
}

// func GetBlockMetadataBalanceUpdatesKind(arg interface{}) (string, error) {
//   block, err := GetBlock(arg)
//   if (err != nil){
//     return "", err
//   }
//   return block.Metadata.BalanceUpdates.Kind, nil
// }
//
// func GetBlockMetadataBalanceUpdatesContract(arg interface{}) (string, error) {
//   block, err := GetBlock(arg)
//   if (err != nil){
//     return "", err
//   }
//   return block.Metadata.BalanceUpdates.Contract, nil
// }
//
// func GetBlockMetadataBalanceUpdatesChange(arg interface{}) (string, error) {
//   block, err := GetBlock(arg)
//   if (err != nil){
//     return "", err
//   }
//   return block.Metadata.BalanceUpdates.Change, nil
// }
//
// func GetBlockMetadataBalanceUpdatesCategory(arg interface{}) (string, error) {
//   block, err := GetBlock(arg)
//   if (err != nil){
//     return "", err
//   }
//   return block.Metadata.BalanceUpdates.Category, nil
// }
//
// func GetBlockMetadataBalanceUpdatesDelegate(arg interface{}) (string, error) {
//   block, err := GetBlock(arg)
//   if (err != nil){
//     return "", err
//   }
//   return block.Metadata.BalanceUpdates.Delegate, nil
// }
//
// func GetBlockMetadataBalanceUpdatesLevel(arg interface{}) (int, error) {
//   block, err := GetBlock(arg)
//   if (err != nil){
//     return 0, err
//   }
//   return block.Metadata.BalanceUpdates.Level, nil
// }

func GetBlockOperations(arg interface{}) ([]Operations, error){
  var operations []Operations
  block, err := BlockCheck(arg)
  if (err != nil){
    return operations, err
  }

  for _, field := range block.Operations[0]{
    protocol := field.Protocol
    chainId := field.ChainID
    hash := field.Hash
    branch := field.Branch
    //contents := nil //GetBlockOperationsContents(block)
    signature := field.Signature
    operations = append(operations, Operations{Protocol: protocol, ChainID: chainId, Hash: hash, Branch: branch, Signature: signature})
  }

  return operations, nil
}

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

// func GetBlockOperationsContents(arg interface{}) ([]Contents, error){
//   var contents []Contents
//   block, err := BlockCheck(arg)
//   if (err != nil){
//     return contents, err
//   }
// }

// func GetBlockOperationsContentsKind(arg interface{}) (string, error){
//
// }
//
// func GetBlockOperationsContentsLevel(arg interface{}) (int, error){
//
// }

// func GetBlockOperationsContentsMetadata(arg interface{}) (Block.Operations.Contents.Metadata, error){
//
// }
//
// func GetBlockOperationsContentsMetadataBalanceUpdates(arg interface{}) ([]Block.Operations.Contents.Metadata.BalanceUpdates, error){
//
// }

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

// func GetBlockOperationsSignature(arg interface{}) (string, error){
//
// }

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
