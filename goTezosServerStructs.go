package goTezosServer

import (
	"time"
)

type Block struct {
	Protocol   string               `json:"protocol"`
	ChainID    string               `json:"chain_id"`
	Hash       string               `json:"hash"`
	Header     StructHeader         `json:"header"`
	Metadata   StructMetadata       `json:"metadata"`
	Operations [][]StructOperations `json:"operations"`
}

type StructHeader struct {
	Level            int       `json:"level"`
	Proto            int       `json:"proto"`
	Predecessor      string    `json:"Predecessor"`
	Timestamp        time.Time `json:"timestamp"`
	ValidationPass   int       `json:"validation_pass"`
	OperationsHash   string    `json:"operations_hash"`
	Fitness          []string  `json:"fitness"`
	Context          string    `json:"context"`
	Priority         int       `json:"priority"`
	ProofOfWorkNonce string    `json:"proof_of_work_nonce"`
	Signature        string    `json:"signature"`
}

type StructMetadata struct {
	Protocol               string                         `json:"protocol"`
	NextProtocol           string                         `json:"next_protocol"`
	TestChainStatus        StructTestChainStatus          `json:"test_chain_status"`
	MaxOperationsTTL       int                            `json:"max_operations_ttl"`
	MaxOperationDataLength int                            `json:"max_operation_data_length"`
	MaxBlockHeaderLength   int                            `json:"max_block_header_length"`
	MaxOperationListLength []StructMaxOperationListLength `json:"max_operation_list_length"`
	Baker                  string                         `json:"baker"`
	Level                  StructLevel                    `json:"level"`
	VotingPeriodKind       string                         `json:"voting_period_kind"`
	NonceHash              interface{}                    `json:"nonce_hash"`
	ConsumedGas            string                         `json:"consumed_gas"`
	Deactivated            []string                       `json:"deactivated"`
	BalanceUpdates         []StructBalanceUpdates         `json:"balance_updates"`
}

type StructTestChainStatus struct {
	Status string `json:"status"`
}

type StructMaxOperationListLength struct {
	MaxSize int `json:"max_size"`
	MaxOp   int `json:"max_op,omitempty"`
}

type StructLevel struct {
	Level                int  `json:"level"`
	LevelPosition        int  `json:"level_position"`
	Cycle                int  `json:"cycle"`
	CyclePosition        int  `json:"cycle_position"`
	VotingPeriod         int  `json:"voting_period"`
	VotingPeriodPosition int  `json:"voting_period_position"`
	ExpectedCommitment   bool `json:"expected_commitment"`
}

type StructBalanceUpdates struct {
	Kind     string `json:"kind"`
	Contract string `json:"contract,omitempty"`
	Change   string `json:"change"`
	Category string `json:"category,omitempty"`
	Delegate string `json:"delegate,omitempty"`
	Level    int    `json:"level,omitempty"`
}

type StructOperations struct {
	Protocol  string           `json:"protocol"`
	ChainID   string           `json:"chain_id"`
	Hash      string           `json:"hash"`
	Branch    string           `json:"branch"`
	Contents  []StructContents `json:"contents"`
	Signature string           `json:"signature"`
}

type StructContents struct {
	Kind             string           `json:"kind"`
	Source           string           `json:"source"`
	Fee              string           `json:"fee"`
	Counter          string           `json:"counter"`
	GasLimit         string           `json:"gas_limit"`
	StorageLimit     string           `json:"storage_limit"`
	Amount           string           `json:"amount"`
	Destination      string           `json:"destination"`
	Delegate         string           `json:"delegate"`
	Phk              string           `json:"phk"`
	Secret           string           `json:"secret"`
	Level            int              `json:"level"`
	ManagerPublicKey string           `json:"managerPubkey"`
	Balance          string           `json:"balance"`
	Metadata         ContentsMetadata `json:"metadata"`
}

type ContentsMetadata struct {
	BalanceUpdates []StructBalanceUpdates `json:"balance_updates"`
	Slots          []int                  `json:"slots"`
}

type StructDelegate struct {
	Balance              string                       `json:"balance"`
	FrozenBalance        string                       `json:"frozen_balance"`
	FrozenBalanceByCycle []StructFrozenBalanceByCycle `json:"frozen_balance_by_cycle"`
	StakingBalance       string                       `json:"staking_balance"`
	DelegateContracts    []string                     `json:"delegated_contracts"`
	DelegatedBalance     string                       `json:"delegated_balance"`
	Deactivated          bool                         `json:"deactivated"`
	GracePeriod          int                          `json:"grace_period"`
	Address              string                       `json:"address"`
	ContractsByCycle     []StructContractsByCycle     `json:"contracts_by_cycle"`
}

type StructContractsByCycle struct {
	Cycle             int                       `json:"cycle"`
	DelegateContracts []StructDelegateContracts `json:"delegate_contracts"`
}

type StructDelegateContracts struct {
	ContractAddress string `json:"contract_address"`
	Balance         int    `json:"balance"`
}

type StructFrozenBalanceByCycle struct {
	Cycle   int    `json:"cycle"`
	Deposit string `json:"deposit"`
	Fees    string `json:"fees"`
	Rewards string `json:"rewards"`
}

// type SnapShot struct {
//   Cycle int
//   RandomSeed string
//   Number string
//   Level int
// }

/*
Description: A structure to hold a snapshot query in.
Cycle (int): The cycle the snapshot was taken in
Number (int): Snap shot number decided. Empty if decided == false
Decided (bool): true or false if snap shot is decided
AssociatedBlock: The block number the snapshot reference, only available if snapshot is decided.
*/
type SnapShot struct {
	Cycle           int
	Number          int
	AssociatedBlock int
}
