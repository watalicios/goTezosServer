package goTezosServer

import (
  "time"
)

type Block struct{
  Protocol    string              `json:"protocol"`
	ChainID     string              `json:"chain_id"`
	Hash        string              `json:"hash"`
  Header      StructHeader        `json:"header"`
  Metadata    StructMetadata      `json:"metadata"`
  Operations  [][]StructOperations  `json:"operations"`
}

type StructHeader struct {
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
  Deactivated              []string                       `json:"deactivated"`
  BalanceUpdates           []StructBalanceUpdates         `json:"balance_updates"`
}

type StructTestChainStatus struct {
  Status string            `json:"status"`
}

type StructMaxOperationListLength struct {
  MaxSize int              `json:"max_size"`
  MaxOp   int              `json:"max_op,omitempty"`
}

type StructLevel struct {
  Level                int              `json:"level"`
  LevelPosition        int              `json:"level_position"`
  Cycle                int              `json:"cycle"`
  CyclePosition        int              `json:"cycle_position"`
  VotingPeriod         int              `json:"voting_period"`
  VotingPeriodPosition int              `json:"voting_period_position"`
  ExpectedCommitment   bool             `json:"expected_commitment"`
}

type StructBalanceUpdates struct {
  Kind     string           `json:"kind"`
  Contract string           `json:"contract,omitempty"`
  Change   string           `json:"change"`
  Category string           `json:"category,omitempty"`
  Delegate string           `json:"delegate,omitempty"`
  Level    int              `json:"level,omitempty"`
}

type StructOperations struct {
  Protocol  string           `json:"protocol"`
  ChainID   string           `json:"chain_id"`
  Hash      string           `json:"hash"`
  Branch    string           `json:"branch"`
  Contents  []StructContents       `json:"contents"`
  Signature string           `json:"signature"`
}

type StructContents struct {
  Kind     string             `json:"kind"`
  Level    int                `json:"level"`
  Metadata ContentsMetadata   `json:"metadata"`
}

type ContentsMetadata struct {
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
