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
	"time"

	"gopkg.in/mgo.v2"
)

var (
	reGetBlockLevelHead = regexp.MustCompile(`"level": ([0-9]+), "proto"`)
	reGetHash           = regexp.MustCompile(`"hash": "([0-9a-zA-Z]+)",`)
	reGetRandomSeed     = regexp.MustCompile(`"random_seed":\n[ ]+ "([0-9a-zA-Z]+)"`)
	reGetRollSnapShot   = regexp.MustCompile(`"roll_snapshot": ([0-9]+)`)
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
