package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: The Tezos API written in GO, for easy development.
License: MIT
*/

import (
  "fmt"
  "os"
  "os/exec"
  "errors"
  "regexp"
  "strconv"
  "time"
  "encoding/json"
  "gopkg.in/mgo.v2"
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

type BlockByte struct{
  Bytes []byte
}

var (
  reGetBlockLevelHead = regexp.MustCompile(`"level": ([0-9]+), "proto"`)
  reGetHash = regexp.MustCompile(`"hash": "([0-9a-zA-Z]+)",`)
)

var TezosPath string
var Session *mgo.Session
var Collection *mgo.Collection

type Person struct {
        Name string
        Phone string
}

/*
Description: This library needs the TEZOSPATH enviroment variable to function
*/
func init() {
  var ok bool
  var errs error
  TezosPath, ok = os.LookupEnv("TEZOSPATH")
  if !ok {
	   fmt.Println("TEZOSPATH not set. Please 'export TEZOSPATH=<path_to_tezos>'.")
	   os.Exit(1)
  }
  TezosPath = TezosPath + "tezos-client"
  if _, err := os.Stat(TezosPath); os.IsNotExist(err) {
    fmt.Println("Could not find tezos-client in TEZOSPATH: " + err.Error())
    os.Exit(1)
  }

  Session, errs = mgo.Dial("127.0.0.1")
  if (errs != nil){
    fmt.Println(errs)
  }
  Collection = Session.DB("TEZOS").C("blocks")

}

func SynchronizeTezosMongo(){
  err := MongoGetAllBlocks()
  if (err != nil){
    fmt.Println(err)
  }
}

func MongoGetAllBlocks() error{
  head, err := GetBlockHead()
  if (err != nil){
    return err
  }

  regHeadLevelResult := reGetBlockLevelHead.FindStringSubmatch(string(head[:]))
  if (regHeadLevelResult == nil){
    return errors.New("Could not get block level for head")
  }
  headLevel, _ := strconv.Atoi(regHeadLevelResult[1]) //TODO Error Checking
  headHash := reGetHash.FindStringSubmatch(string(head[:])) //TODO Error check the regex
  if (headHash == nil){
    return errors.New("Could not get hash for block head")
  }

  for i := headLevel-headLevel; i < headLevel; i ++{
    block, err := GetBlock(i, headHash[1], headLevel)
    if (err != nil){
      return err
    }
    err = Collection.Insert(block)
    if (err != nil){
      fmt.Println(err)
    }
  }
  return nil
}

func GetBlock(level int, headHash string, headLevel int) (Block, error){
  diff := headLevel - level
  diffStr := strconv.Itoa(diff)
  getBlockByLevel := "chains/main/blocks/" + headHash + "~" + diffStr
  var block Block

  s, err := TezosRPCGet(getBlockByLevel)
  if (err != nil){
    return block, err
  }

  block = ConvertBlockToJson(s)
  return block, nil
}

/*
Description: Takes a block level, and returns the hash for that specific level
Param level (int): An integer representation of the block level to query
Returns (string): A string representation of the hash for the block level queried.
*/
// func GetBlockLevelHash(level int, headHash string, headLevel int) (string, error){
//   diff :=  headLevel - level
//   diffStr := strconv.Itoa(diff)
//   getBlockByLevel := "chains/main/blocks/" + headHash + "~" + diffStr
//
//   s, err := TezosRPCGet(getBlockByLevel)
//   if (err != nil){
//     return "", errors.New("Could not get hash for block " +  strconv.Itoa(level) + ": TezosRPCGet(arg string) failed: " + err.Error())
//   }
//
//   hash := reGetHash.FindStringSubmatch(s) //TODO Error check the regex
//   if (hash == nil){
//     return "", errors.New("Could not get hash for block " + strconv.Itoa(level))
//   }
//
//   return hash[1], nil
// }

/*
Description: Will retreive the current block level as an integer
Returns (int): Returns integer representation of block level
*/
func GetBlockHead() ([]byte, error){
  s, err := TezosRPCGet("chains/main/blocks/head")
  if (err != nil){
    return s, errors.New("Could not get block level for head: TezosRPCGet(arg string) failed: " + err.Error())
  }
  return s, nil
}





/*
Description: Takes an  array of interface (struct in our case), jsonifies it, and allows a much neater print.
Param v (interface{}): Array of an interface
*/
func ConvertBlockToJson(v []byte) Block {
  var block Block
  err := json.Unmarshal(v, &block)
  if err != nil {
        panic(err)
  }

  return block
}


/*
Description: A function that executes a command to the tezos-client
Param args ([]string): Arguments to be executed
Returns (string): Returns the output of the executed command as a string
*/
func MongoAddBlock(json string) (error){

  json = "[" + json + "]"
  fmt.Println(json)
  out, err := exec.Command("mongoimport", "--db", "TEZOS", "--collection", "blocks", "--type", "json", json).Output()
  if err != nil {
    fmt.Println(err)
    return  err
  }

  fmt.Println(out)

  return nil
}

/*
Description: A function that executes a command to the tezos-client
Param args ([]string): Arguments to be executed
Returns (string): Returns the output of the executed command as a string
*/
func TezosDo(args ...string) ([]byte, error){
  out, err := exec.Command(TezosPath, args...).Output()
  if err != nil {
    return out, err
  }

  return out, nil
}

/*
Description: A function that executes an rpc get arg
Param args ([]string): Arguments to be executed
Returns (string): Returns the output of the executed command as a string
*/
func TezosRPCGet(arg string) ([]byte, error){
  output, err := TezosDo("rpc", "get", arg)
  if (err != nil){
    return output, errors.New("Could not rpc get " + arg + " : tezosDo(args ...string) failed: " + err.Error())
  }
  return output, nil
}
