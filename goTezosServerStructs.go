package goTezosServer

import (
  "time"
  "gopkg.in/mgo.v2/bson"
)

func BsonBlockToBlock(bsonBlock BsonBlock) Block{
  var block Block
  block.Protocol = bsonBlock.Protocol
  block.ChainID = bsonBlock.ChainID
  block.Hash = bsonBlock.Hash

  var header StructHeader
  header.Level = bsonBlock.Header.Level
  header.Proto = bsonBlock.Header.Proto
  header.Predecessor = bsonBlock.Header.Predecessor
  header.Timestamp = bsonBlock.Header.Timestamp
  header.ValidationPass = bsonBlock.Header.ValidationPass
  header.OperationsHash = bsonBlock.Header.OperationsHash
  header.Fitness = bsonBlock.Header.Fitness
  header.Context = bsonBlock.Header.Context
  header.Priority = bsonBlock.Header.Priority
  header.ProofOfWorkNonce = bsonBlock.Header.ProofOfWorkNonce
  header.Signature = bsonBlock.Header.Signature
  block.Header = header

  var metadata StructMetadata
  metadata.Protocol = bsonBlock.Metadata.Protocol
  metadata.NextProtocol = bsonBlock.Metadata.NextProtocol
  var testChainStatus StructTestChainStatus
  testChainStatus.Status = bsonBlock.Metadata.TestChainStatus.Status
  metadata.TestChainStatus = testChainStatus
  metadata.MaxOperationsTTL = bsonBlock.Metadata.MaxOperationsTTL
  metadata.MaxOperationDataLength = bsonBlock.Metadata.MaxOperationDataLength
  metadata.MaxBlockHeaderLength = bsonBlock.Metadata.MaxBlockHeaderLength
  var maxOperationListLengths []StructMaxOperationListLength
  for _, field := range bsonBlock.Metadata.MaxOperationListLength{
    var maxOperationListLength StructMaxOperationListLength
    maxOperationListLength.MaxSize = field.MaxSize
    maxOperationListLength.MaxSize = field.MaxOp
    maxOperationListLengths = append(maxOperationListLengths, maxOperationListLength)
  }
  metadata.MaxOperationListLength = maxOperationListLengths
  metadata.Baker = bsonBlock.Metadata.Baker
  var level StructLevel
  level.Level = bsonBlock.Metadata.Level.Level
  level.LevelPosition = bsonBlock.Metadata.Level.LevelPosition
  level.Cycle = bsonBlock.Metadata.Level.Cycle
  level.CyclePosition = bsonBlock.Metadata.Level.CyclePosition
  level.VotingPeriod = bsonBlock.Metadata.Level.VotingPeriod
  level.VotingPeriodPosition = bsonBlock.Metadata.Level.VotingPeriodPosition
  level.ExpectedCommitment = bsonBlock.Metadata.Level.ExpectedCommitment
  metadata.Level = level
  metadata.VotingPeriodKind = bsonBlock.Metadata.VotingPeriodKind
  metadata.NonceHash = bsonBlock.Metadata.NonceHash
  metadata.ConsumedGas = bsonBlock.Metadata.ConsumedGas
  metadata.Deactivated = bsonBlock.Metadata.Deactivated
  var balanceUpdates []StructBalanceUpdates
  for _, field := range bsonBlock.Metadata.BalanceUpdates{
    var balanceUpdate StructBalanceUpdates
    balanceUpdate.Kind = field.Kind
    balanceUpdate.Contract = field.Contract
    balanceUpdate.Change = field.Change
    balanceUpdate.Category = field.Category
    balanceUpdate.Delegate = field.Delegate
    balanceUpdate.Level = field.Level
    balanceUpdates = append(balanceUpdates, balanceUpdate)
  }
  metadata.BalanceUpdates = balanceUpdates

  var operations []StructOperations
  for _, op := range bsonBlock.Operations{
    for _, op2 := range op {
      var operation StructOperations
      operation.Protocol = op2.Protocol
      operation.ChainID = op2.ChainID
      operation.Hash = op2.Hash
      operation.Branch = op2.Branch
      var contents []StructContents
      for _, cont := range op2.Contents{
        var content StructContents
        content.Kind = cont.Kind
        content.Level = cont.Level
        var contentMetadata ContentsMetadata
        var balanceUpdates []StructBalanceUpdates
        for _, bup := range cont.Metadata.BalanceUpdates{
          var balanceUpdate StructBalanceUpdates
          balanceUpdate.Kind = bup.Kind
          balanceUpdate.Contract = bup.Contract
          balanceUpdate.Change = bup.Change
          balanceUpdate.Category = bup.Category
          balanceUpdate.Delegate = bup.Delegate
          balanceUpdate.Level = bup.Level
          balanceUpdates = append(balanceUpdates, balanceUpdate)
        }
        contentMetadata.BalanceUpdates = balanceUpdates
        contentMetadata.Delegate = cont.Metadata.Delegate
        contentMetadata.Slots = cont.Metadata.Slots
        content.Metadata = contentMetadata
        contents = append(contents, content)
      }
      operation.Contents = contents
      operation.Signature = op2.Signature
      operations = append(operations, operation)
    }
  }

  block.Metadata = metadata
  block.Operations = operations

  return block
}

type Block struct{
  ID          bson.ObjectId       `bson:"_id,omitempty"`
  Protocol    string              `json:"protocol"`
	ChainID     string              `json:"chain_id"`
	Hash        string              `json:"hash"`
  Header      StructHeader        `json:"header"`
  Metadata    StructMetadata      `json:"metadata"`
  Operations  []StructOperations  `json:"operations"`
}

type BsonBlock struct {
	Protocol string `json:"protocol"`
	ChainID  string `json:"chain_id"`
	Hash     string `json:"hash"`
	Header   struct {
		Level            int       `json:"level"`
		Proto            int       `json:"proto"`
		Predecessor      string    `json:"predecessor"`
		Timestamp        time.Time `json:"timestamp"`
		ValidationPass   int       `json:"validation_pass"`
		OperationsHash   string    `json:"operations_hash"`
		Fitness          []string  `json:"fitness"`
		Context          string    `json:"context"`
		Priority         int       `json:"priority"`
		ProofOfWorkNonce string    `json:"proof_of_work_nonce"`
		Signature        string    `json:"signature"`
	} `json:"header"`
  Metadata struct {
		Protocol        string `json:"protocol"`
		NextProtocol    string `json:"next_protocol"`
		TestChainStatus struct {
			Status string `json:"status"`
		} `json:"test_chain_status"`
		MaxOperationsTTL       int `json:"max_operations_ttl"`
		MaxOperationDataLength int `json:"max_operation_data_length"`
		MaxBlockHeaderLength   int `json:"max_block_header_length"`
		MaxOperationListLength []struct {
			MaxSize int `json:"max_size"`
			MaxOp   int `json:"max_op,omitempty"`
		} `json:"max_operation_list_length"`
		Baker string `json:"baker"`
		Level struct {
			Level                int  `json:"level"`
			LevelPosition        int  `json:"level_position"`
			Cycle                int  `json:"cycle"`
			CyclePosition        int  `json:"cycle_position"`
			VotingPeriod         int  `json:"voting_period"`
			VotingPeriodPosition int  `json:"voting_period_position"`
			ExpectedCommitment   bool `json:"expected_commitment"`
		} `json:"level"`
		VotingPeriodKind string        `json:"voting_period_kind"`
		NonceHash        interface{}   `json:"nonce_hash"`
		ConsumedGas      string        `json:"consumed_gas"`
		Deactivated      []interface{} `json:"deactivated"`
		BalanceUpdates   []struct {
			Kind     string `json:"kind"`
			Contract string `json:"contract,omitempty"`
			Change   string `json:"change"`
			Category string `json:"category,omitempty"`
			Delegate string `json:"delegate,omitempty"`
			Level    int    `json:"level,omitempty"`
		} `json:"balance_updates"`
	} `json:"metadata"`
	Operations [][]struct {
		Protocol string `json:"protocol"`
		ChainID  string `json:"chain_id"`
		Hash     string `json:"hash"`
		Branch   string `json:"branch"`
		Contents []struct {
			Kind     string `json:"kind"`
			Level    int    `json:"level"`
			Metadata struct {
				BalanceUpdates []struct {
					Kind     string `json:"kind"`
					Contract string `json:"contract,omitempty"`
					Change   string `json:"change"`
					Category string `json:"category,omitempty"`
					Delegate string `json:"delegate,omitempty"`
					Level    int    `json:"level,omitempty"`
				} `json:"balance_updates"`
				Delegate string `json:"delegate"`
				Slots    []int  `json:"slots"`
			} `json:"metadata"`
		} `json:"contents"`
		Signature string `json:"signature"`
	} `json:"operations"`
}

type StructHeader struct {
  ID               bson.ObjectId `bson:"_id,omitempty"`
  Level            int           `json:"level"`
  Proto            int           `json:"proto"`
  Predecessor      string        `json:"Predecessor"`
  Timestamp        time.Time     `json:"timestamp"`
  ValidationPass   int           `json:"validation_pass"`
  OperationsHash   string        `json:"operations_hash"`
  Fitness          []string      `json:"fitness"`
  Context          string        `json:"context"`
  Priority         int           `json:"priority"`
  ProofOfWorkNonce string        `json:"proof_of_work_nonce"`
  Signature        string        `json:"signature"`
}

type StructMetadata struct {
  ID                       bson.ObjectId                  `bson:"_id,omitempty"`
  Protocol                 string                         `json:"protocol"`
  NextProtocol             string                         `json:"next_protocol"`
  TestChainStatus          StructTestChainStatus          `json:"test_chain_status"`
  MaxOperationsTTL         int                            `json:"max_operations_ttl"`
  MaxOperationDataLength   int                            `json:"max_operation_data_length"`
  MaxBlockHeaderLength     int                            `json:"max_block_header_length"`
  MaxOperationListLength   []StructMaxOperationListLength `json:"max_operation_list_length"`
  Baker                    string                         `json:"baker"`
  Level                    StructLevel                    `json:"level"`
  VotingPeriodKind         string                         `json:"voting_period_kind"`
  NonceHash                interface{}                    `json:"nonce_hash"`
  ConsumedGas              string                         `json:"consumed_gas"`
  Deactivated              []interface{}                  `json:"deactivated"`
  BalanceUpdates           []StructBalanceUpdates         `json:"balance_updates"`
}

type StructTestChainStatus struct {
  ID     bson.ObjectId     `bson:"_id,omitempty"`
  Status string            `json:"status"`
}

type StructMaxOperationListLength struct {
  ID     bson.ObjectId     `bson:"_id,omitempty"`
  MaxSize int              `json:"max_size"`
  MaxOp   int              `json:"max_op,omitempty"`
}

type StructLevel struct {
  ID                   bson.ObjectId    `bson:"_id,omitempty"`
  Level                int              `json:"level"`
  LevelPosition        int              `json:"level_position"`
  Cycle                int              `json:"cycle"`
  CyclePosition        int              `json:"cycle_position"`
  VotingPeriod         int              `json:"voting_period"`
  VotingPeriodPosition int              `json:"voting_period_position"`
  ExpectedCommitment   bool             `json:"expected_commitment"`
}

type StructBalanceUpdates struct {
  ID       bson.ObjectId    `bson:"_id,omitempty"`
  Kind     string           `json:"kind"`
  Contract string           `json:"contract,omitempty"`
  Change   string           `json:"change"`
  Category string           `json:"category,omitempty"`
  Delegate string           `json:"delegate,omitempty"`
  Level    int              `json:"level,omitempty"`
}

type StructOperations struct {
  ID        bson.ObjectId    `bson:"_id,omitempty"`
  Protocol  string           `json:"protocol"`
  ChainID   string           `json:"chain_id"`
  Hash      string           `json:"hash"`
  Branch    string           `json:"branch"`
  Contents  []StructContents       `json:"contents"`
  Signature string           `json:"signature"`
}

type StructContents struct {
  ID        bson.ObjectId    `bson:"_id,omitempty"`
  Kind     string
  Level    int
  Metadata ContentsMetadata
}

type ContentsMetadata struct {
  ID        bson.ObjectId               `bson:"_id,omitempty"`
  BalanceUpdates []StructBalanceUpdates `json:"balance_updates"`
  Delegate string                       `json:"delegate"`
  Slots    []int                        `json:"slots"`
}

type SnapShot struct {
  Cycle int
  RandomSeed string
  Number string
  Level int
}
