package goTezosServer

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

// func GetBlockHeader(arg interface{}) (Block.Header, error) {
//   block, err := GetBlock(arg)
//   if (err != nil){
//     return "", err
//   }
//   return block.Header, nil
// }

func GetBlockHeaderLevel(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Header.Level, nil
}

func GetBlockHeaderProto(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Header.Proto, nil
}

func GetBlockHeaderPredecessor(arg interface{}) (string, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Header.Predecessor, nil
}

func GetBlockHeaderTimeStamp(arg interface{}) (time.Time, error) {
  var t time.Time
  block, err := GetBlock(arg)
  if (err != nil){
    return t, err
  }
  t = block.Header.Timestamp
  return t, nil
}

func GetBlockHeaderValidationPass(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Header.ValidationPass, nil
}

func GetBlockHeaderOperationsHash(arg interface{}) (string, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Header.OperationsHash, nil
}

func GetBlockHeaderFitness(arg interface{}) ([]string, error) {
  var str []string
  block, err := GetBlock(arg)
  if (err != nil){
    return str, err
  }
  str = block.Header.Fitness
  return str, nil
}

func GetBlockHeaderContext(arg interface{}) (string, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Header.Context, nil
}

func GetBlockHeaderPriority(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Header.Priority, nil
}

func GetBlockHeaderProofOfWorkNonce(arg interface{}) (string, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Header.ProofOfWorkNonce, nil
}

func GetBlockHeaderSignature(arg interface{}) (string, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Header.ProofOfWorkNonce, nil
}

// func GetBlockMetadata(arg interface{}) (Block.Metadata, error) {
//
// }

func GetBlockMetadataProtocol(arg interface{}) (string, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.Protocol, nil
}

func GetBlockMetadataNextProtocol(arg interface{}) (string, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.NextProtocol, nil
}

// func GetBlockMetadataTestChainStatus(arg interface{}) (Block.Metadata.TestChainStatus, error) {
//
// }

func GetBlockMetadataMaxOperationsTTL(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.MaxOperationsTTL, nil
}

func GetBlockMetadataMaxOperationDataLength(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.MaxOperationDataLength, nil
}

func GetBlockMetadataMaxBlockHeaderLength(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.MaxBlockHeaderLength, nil
}

// // func GetBlockMetadataMaxOperationListLength(arg interface{}) ([]Block.Metadata.MaxOperationListLength, error) {
// //
// // }
//
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
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.Baker, nil
}

// // func GetBlockMetadataLevel(arg interface{}) ([]Block.Metadata.Level, error) {
// //
// // }
//
func GetBlockMetadataLevelLevel(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.Level.Level, nil
}

func GetBlockMetadataLevelLevelPosition(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.Level.LevelPosition, nil
}

func GetBlockMetadataLevelCycle(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.Level.Cycle, nil
}

func GetBlockMetadataLevelCyclePosition(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.Level.CyclePosition, nil
}

func GetBlockMetadataLevelVotingPeriod(arg interface{}) (int, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return 0, err
  }
  return block.Metadata.Level.VotingPeriod, nil
}

func GetBlockMetadataLevelExpectedCommitment(arg interface{}) (bool, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return false, err
  }
  return block.Metadata.Level.ExpectedCommitment, nil
}

func GetBlockMetadataVotingPeriodKind(arg interface{}) (string, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.VotingPeriodKind, nil
}

func GetBlockMetadataNonceHash(arg interface{}) (interface{}, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return arg, err
  }
  return block.Metadata.NonceHash, nil
}

func GetBlockMetadataConsumedGas(arg interface{}) (string, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return "", err
  }
  return block.Metadata.ConsumedGas, nil
}

func GetBlockMetadataDeactivated(arg interface{}) (interface{}, error) {
  block, err := GetBlock(arg)
  if (err != nil){
    return arg, err
  }
  return block.Metadata.Deactivated, nil
}

// // func GetBlockMetadataBalanceUpdates(arg interface{}) ([]Block.Metadata.BalanceUpdates, error) {
// //
// // }
//
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

// // func GetBlockOperations(arg interface{}) (Block.Operations, error){
// //
// // }

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

// func GetBlockOperationsBranch(arg interface{}) (string, error){
//
// }
//
// // func GetBlockOperationsContents(arg interface{}) ([]Block.Operations.Contents, error){
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
// // func GetBlockOperationsContentsMetadata(arg interface{}) (Block.Operations.Contents.Metadata, error){
// //
// // }
// //
// // func GetBlockOperationsContentsMetadataBalanceUpdates(arg interface{}) ([]Block.Operations.Contents.Metadata.BalanceUpdates, error){
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
