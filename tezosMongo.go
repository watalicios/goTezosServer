package goTezosServer

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: The Tezos API written in GO, for easy development.
License: MIT
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	reGetBlockLevelHead  = regexp.MustCompile(`"level": ([0-9]+), "proto"`)
	reGetHash            = regexp.MustCompile(`"hash": "([0-9a-zA-Z]+)",`)
	reGetRandomSeed      = regexp.MustCompile(`"random_seed":\n[ ]+ "([0-9a-zA-Z]+)"`)
	reGetRollSnapShot    = regexp.MustCompile(`"roll_snapshot": ([0-9]+)`)
	reGetBalance         = regexp.MustCompile(`([0-9.]+)`)
	reDelegatedContracts = regexp.MustCompile(`"([A-Z0-9a-z]+)"`)
)

var TezosPath string
var Session *mgo.Session
var Collection *mgo.Collection

/*
Description: This library needs the TEZOSPATH enviroment variable to function
*/
func init() {
	var ok bool
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
}

func SetDatabaseConnection(connection, database, collection string) {
	Session, err := mgo.Dial(connection)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Collection = Session.DB(database).C(collection)
}

func InitSynchronizeTezosMongo() {
	err := MongoGetAllBlocks()
	if err != nil {
		fmt.Println(err)
	}
}

func SynchronizeTezosMongo() {
	blockDb, _ := GetBlockHead() //Get last block in db
	level := blockDb.Header.Level
	nextLevel := level + 1
	run := true

	for run == true {
		tmpBlock, _ := GetBlockRPCHead() //Get current head in rpc
		blockHead := ConvertBlockToJson(tmpBlock)

		if nextLevel <= blockHead.Header.Level {
			block, err := GetBlockRPC(nextLevel, blockHead.Hash, blockHead.Header.Level)
			if err != nil {
				fmt.Println(err)
			}
			err = Collection.Insert(block)
			if err != nil {
				fmt.Println(err)
			}
			nextLevel = nextLevel + 1
		}
		time.Sleep(1 * time.Second)
	}
}

func MongoGetAllBlocks() error {
	head, err := GetBlockRPCHead()
	if err != nil {
		return err
	}

	regHeadLevelResult := reGetBlockLevelHead.FindStringSubmatch(string(head[:]))
	if regHeadLevelResult == nil {
		return errors.New("Could not get block level for head")
	}
	headLevel, _ := strconv.Atoi(regHeadLevelResult[1])       //TODO Error Checking
	headHash := reGetHash.FindStringSubmatch(string(head[:])) //TODO Error check the regex
	if headHash == nil {
		return errors.New("Could not get hash for block head")
	}

	for i := headLevel - headLevel; i < headLevel; i++ {
		block, err := GetBlockRPC(i, headHash[1], headLevel)
		if err != nil {
			return err
		}

		// for _, op := range block.Operations{
		//   for _, operation := range op{
		//     for _, content := range op.Contents{
		//       if (content.Kind == "delegation"){
		//         content.Delegate
		//       }
		//     }
		//   }
		// }
		err = Collection.Insert(block)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func GetBlockRPC(level int, headHash string, headLevel int) (Block, error) {
	diff := headLevel - level
	diffStr := strconv.Itoa(diff)
	getBlockByLevel := "chains/main/blocks/" + headHash + "~" + diffStr
	var block Block

	s, err := TezosRPCGet(getBlockByLevel)
	if err != nil {
		return block, err
	}

	block = ConvertBlockToJson(s)
	return block, nil
}

func GetBlockRPCHead() ([]byte, error) {
	s, err := TezosRPCGet("chains/main/blocks/head")
	if err != nil {
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

func GetSnapShot(cycle int) (SnapShot, error) {
	var snapShot SnapShot
	snapShot.Cycle = cycle
	strCycle := strconv.Itoa(cycle)

	snapshotStr := "/chains/main/blocks/head/context/raw/json/cycle/" + strCycle

	s, err := TezosRPCGet(snapshotStr)
	if err != nil {
		return snapShot, errors.New("func GetSnapShot(cycle int) (SnapShot, error) failed: " + err.Error())
	}

	regRandomSeed := reGetRandomSeed.FindStringSubmatch(string(s[:]))
	if regRandomSeed == nil {
		return snapShot, errors.New("No random seed: func GetSnapShot(cycle int) (SnapShot, error) failed.")
	}

	regRollSnapShot := reGetRollSnapShot.FindStringSubmatch(string(s[:]))
	if regRollSnapShot == nil {
		return snapShot, errors.New("Could not parse snapshot: func GetSnapShot(cycle int) (SnapShot, error) failed.")
	}
	number, _ := strconv.Atoi(regRollSnapShot[1])
	snapShot.Number = number
	snapShot.AssociatedBlock = ((cycle - 7) * 4096) + (number+1)*256

	return snapShot, nil
}

/*
Description: Will get the balance of an account at a specific snapshot
Param tezosAddr (string): Takes a string representation of the address querying
Param cycle (int): The cycle we are getting the snapshot for
Returns (float64): Returns a float64 representation of the balance for the account
*/
func GetAccountBalanceAtSnapshot(tezosAddr string, cycle int) (float64, error) {
	snapShot, err := GetSnapShot(cycle)
	if err != nil {
		return 0, errors.New("Could not get balance at snapshot for " + tezosAddr + ": GetSnapShot(cycle int) failed: " + err.Error())
	}

	hash, err := GetBlockLevelHash(snapShot.AssociatedBlock)
	//fmt.Println(hash)
	if err != nil {
		return 0, errors.New("Could not get hash for block " + strconv.Itoa(snapShot.AssociatedBlock) + ": GetBlockLevelHead() failed: " + err.Error())
	}

	balanceCmdStr := "/chains/main/blocks/" + hash + "/context/contracts/" + tezosAddr + "/balance"

	s, err := TezosRPCGet(balanceCmdStr)
	if err != nil {
		return 0, errors.New("Could not get balance at snapshot for " + tezosAddr + ": TezosRPCGet(arg string) failed: " + err.Error())
	}

	var returnBalance float64
	var regGetBalance []string

	if strings.Contains(string(s[:]), "No service found at this URL") {
		returnBalance = 0
	} else {
		regGetBalance = reGetBalance.FindStringSubmatch(string(s[:]))
		if regGetBalance == nil {
			return 0, errors.New("Could not parse balance for " + string(s[:]))
		}
	}

	if len(regGetBalance) < 1 {
		returnBalance = 0
	} else {
		floatBalance, _ := strconv.ParseFloat(regGetBalance[1], 64) //TODO error checking
		returnBalance = floatBalance
	}

	return returnBalance / 1000000, nil
}

/*
Description: Will retreive the current block level as an integer
Returns (int): Returns integer representation of block level
*/
func GetBlockLevelHead() (int, string, error) {
	s, err := TezosRPCGet("chains/main/blocks/head")
	if err != nil {
		return 0, "", errors.New("func GetBlockLevelHead() failed: " + err.Error())
	}

	regHeadLevelResult := reGetBlockLevelHead.FindStringSubmatch(string(s[:]))
	if regHeadLevelResult == nil {
		return 0, "", errors.New("Could not parse head level: func GetBlockLevelHead() failed.")
	}
	regHash := reGetHash.FindStringSubmatch(string(s[:]))
	if regHash == nil {
		return 0, "", errors.New("Could not parse head hash: func GetBlockLevelHead() failed.")
	}
	headlevel, _ := strconv.Atoi(regHeadLevelResult[1]) //TODO Error Checking

	return headlevel, regHash[1], nil
}

/*
Description: Takes a block level, and returns the hash for that specific level
Param level (int): An integer representation of the block level to query
Returns (string): A string representation of the hash for the block level queried.
*/
func GetBlockLevelHash(level int) (string, error) {
	head, headHash, err := GetBlockLevelHead()
	if err != nil {
		return "", errors.New("func GetBlockLevelHash(level int) failed: " + err.Error())
	}
	diff := head - level

	diffStr := strconv.Itoa(diff)
	getBlockByLevel := "chains/main/blocks/" + headHash + "~" + diffStr

	s, err := TezosRPCGet(getBlockByLevel)
	if err != nil {
		return "", errors.New("func GetBlockLevelHash(level int) failed: " + err.Error())
	}

	hash := reGetHash.FindStringSubmatch(string(s[:])) //TODO Error check the regex
	if hash == nil {
		return "", errors.New("Could not parse hash: func GetBlockLevelHash(level int) failed.")
	}

	return hash[1], nil
}

/*
Description: Retrieves the list of addresses delegated to a delegate
Param SnapShot: A SnapShot object describing the desired snap shot.
Param delegateAddr: A string that represents a delegators tz address.
Returns []string: An array of contracts delegated to the delegator during the snap shot
*/
func GetDelegatedContractsForCycle(cycle int, delegateAddr string) ([]string, error) {
	var rtnString []string
	snapShot, err := GetSnapShot(cycle)
	// fmt.Println(snapShot)
	if err != nil {
		return rtnString, errors.New("Could not get delegated contracts for cycle " + strconv.Itoa(cycle) + ": GetSnapShot(cycle int) failed: " + err.Error())
	}
	hash, err := GetBlockLevelHash(snapShot.AssociatedBlock)
	if err != nil {
		return rtnString, errors.New("Could not get delegated contracts for cycle " + strconv.Itoa(cycle) + ": GetBlockLevelHash(level int) failed: " + err.Error())
	}
	// fmt.Println(hash)
	getDelegatedContracts := "/chains/main/blocks/" + hash + "/context/delegates/" + delegateAddr + "/delegated_contracts"

	s, err := TezosRPCGet(getDelegatedContracts)
	if err != nil {
		return rtnString, errors.New("Could not get delegated contracts for cycle " + strconv.Itoa(cycle) + ": TezosRPCGet(arg string) failed: " + err.Error())
	}

	DelegatedContracts := reDelegatedContracts.FindAllStringSubmatch(string(s[:]), -1)
	if DelegatedContracts == nil {
		return rtnString, errors.New("Could not get delegated contracts for cycle " + strconv.Itoa(cycle) + ": You have no contracts.")
	}
	rtnString = addressesToArray(DelegatedContracts)
	return rtnString, nil
}

/*
Description: Takes a multi-dimmensional array of addresses from a regex parse, and converts them into a single index(able) array.
Param matches ([][]string): All the addresses found and parsed by regex (ex. DelegatedContracts := reDelegatedContracts.FindAllStringSubmatch(s, -1) returns a multi dimmensional array)
Returns ([]string): Returns an index(able) string array of the matches input.
*/
func addressesToArray(matches [][]string) []string {
	var addresses []string
	for _, x := range matches {
		addresses = append(addresses, x[1])
	}

	return addresses
}

/*
Description: A function that executes a command to the tezos-client
Param args ([]string): Arguments to be executed
Returns (string): Returns the output of the executed command as a string
*/
func TezosDo(args ...string) ([]byte, error) {
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
func TezosRPCGet(arg string) ([]byte, error) {
	output, err := TezosDo("rpc", "get", arg)
	if err != nil {
		return output, errors.New("Could not rpc get " + arg + " : tezosDo(args ...string) failed: " + err.Error())
	}
	return output, nil
}

func PrettyReport(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		return string(b)
	}
	return ""
}
