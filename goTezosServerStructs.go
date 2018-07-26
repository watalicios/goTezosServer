package goTezosServer

import (
  "time"
)

type Block struct {
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
	Operations []struct {
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

type Header struct {
  Level            int
  Proto            int
  Predecessor      string
  Timestamp        time.Time
  ValidationPass   int
  OperationsHash   string
  Fitness          []string
  Context          string
  Priority         int
  ProofOfWorkNonce string
  Signature        string
}

type Metadata struct {
  Protocol                 string
  NextProtocol             string
  TstChainStatus           TestChainStatus
  MaxOperationsTTL         int
  MaxOperationDataLength   int
  MaxBlockHeaderLength     int
  MxOperationListLength    []MaxOperationListLength
  Baker                    string
  Lvl                      Level
  VotingPeriodKind         string
  NonceHash                interface{}
  ConsumedGas              string
  Deactivated              []interface{}
  BlncUpdates              []BalanceUpdates
}

type TestChainStatus struct {
  Status string
}

type MaxOperationListLength struct {
  MaxSize int
  MaxOp   int
}

type Level struct {
  Level                int
  LevelPosition        int
  Cycle                int
  CyclePosition        int
  VotingPeriod         int
  VotingPeriodPosition int
  ExpectedCommitment   bool
}

type BalanceUpdates struct {
  Kind     string
  Contract string
  Change   string
  Category string
  Delegate string
  Level    int
}

type Operations struct {
  Protocol string
  ChainID  string
  Hash     string
  Branch   string
  Cont     []Contents
  Signature string
}

type Contents struct {
  Kind     string
  Level    int
  Metadata ContentsMetadata
}

type ContentsMetadata struct {
  balUpdates []BalanceUpdates
  Delegate string
  Slots    []int
}

type OpBalanceUpdates struct {
  Kind     string
  Contract string
  Change   string
  Category string
  Delegate string
  Level    int
}

type SnapShot struct {
  Cycle int
  RandomSeed string
  Number string
  Level int
}
