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
  "github.com/mongodb/mongo-go-driver/mongo"
  "github.com/mongodb/mongo-tools/common/bsonutil"
  "context"
  "regexp"
  "strconv"
  //"bytes"
  "encoding/json"
)

type BlockByte struct{
  Bytes []byte
}

var (
  reGetBlockLevelHead = regexp.MustCompile(`"level": ([0-9]+), "proto"`)
  reGetHash = regexp.MustCompile(`"hash": "([0-9a-zA-Z]+)",`)
)

var TezosPath string

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

func SynchronizeTezosMongo(){
  blocks, err := GetAllBlocks()
  if (err != nil){
    fmt.Println(err)
  }

  client, err := mongo.NewClient("mongodb://127.0.0.1:27017")
  if err != nil {
     fmt.Println(err)
   }
  err = client.Connect(context.TODO())
  if err != nil {
     fmt.Println(err)
  }

  collection := client.Database("TEZOS").Collection("blocks")

  for _, block := range blocks{
    //fmt.Println(block)
    _, err := collection.InsertOne(context.Background(), block)
    if err != nil { fmt.Println(err) }
  }
}

func GetAllBlocks() ([]interface{}, error){
  var blocks []interface{}
  head, err := GetBlockHead()
  if (err != nil){
    return blocks, err
  }

  regHeadLevelResult := reGetBlockLevelHead.FindStringSubmatch(string(head[:]))
  if (regHeadLevelResult == nil){
    return blocks, errors.New("Could not get block level for head")
  }
  headLevel, _ := strconv.Atoi(regHeadLevelResult[1]) //TODO Error Checking
  headHash := reGetHash.FindStringSubmatch(string(head[:])) //TODO Error check the regex
  if (headHash == nil){
    return blocks, errors.New("Could not get hash for block head")
  }

  for i := headLevel-10; i < headLevel; i ++{
    block, err := GetBlock(i, headHash[1], headLevel)
    if (err != nil){
      return blocks, err
    }
    blocks = append(blocks, block)
  }
  return blocks, nil
}

func GetBlock(level int, headHash string, headLevel int) (interface{}, error){
  diff := headLevel - level
  diffStr := strconv.Itoa(diff)
  getBlockByLevel := "chains/main/blocks/" + headHash + "~" + diffStr
  var blockByte interface{}

  s, err := TezosRPCGet(getBlockByLevel)
  if (err != nil){
    return blockByte, err
  }

  blockByte = ConvertToBson(s)
  return blockByte, nil
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
func ConvertToBson(v interface{}) interface{} {
  b, _ := json.MarshalIndent(v, "", "  ")
  //fmt.Println(b)
  block, err := bsonutil.ConvertJSONValueToBSON(string(b))
  if (err != nil){
    fmt.Println("ERROR: " + err.Error())
  }
//  fmt.Println(doc)
  return block
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
