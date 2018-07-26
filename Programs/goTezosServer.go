package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: The Tezos API written in GO, for easy development.
License: MIT
*/

import (
  "sync"
  "flag"
  "fmt"
  // "log"
  // "html"
  // "net/http"
  // "github.com/gorilla/mux"
  "github.com/DefinitelyNotAGoat/goTezosServer"
)

var wg sync.WaitGroup

func main(){
	init := flag.Bool("init", true, "Start synchronization of the database from cycle 0")
  //database := flag.String("database", "TEZOS", "Database name for Tezos data")

  //goTezosServer.SetDataBase(*database)

  if (*init){
    fmt.Println("Initializing the server from cycle 0.")
    goTezosServer.InitSynchronizeTezosMongo()
    fmt.Println("Done Initializing.")
  }
  wg.Add(1)
  go goTezosServer.SynchronizeTezosMongo()

  wg.Wait()
  wg.Done()

}

// func getBlockTest() (bool, error){
//   pass := false
//   block goTezosServer.Block
//   //Get block by level
//   block, err := tezQuery.GetBlock(10)
//   if (err == nil){
//     pass = true
//   }
//
//   //Get block by Hash
//   block, err := tezQuery.GetBlock("BLq3jarZuxz4F7pamd9xvNmrSstMqAS2JMP5SB2VMqpdWpN88eJ")
//   if (err == nil){
//     pass = true
//   }
//
//   //Get block by bad type
//   tezQuery.GetBlock(30.22)
//   if (err != nil){
//     return true, err
//   }
//
//   //Get block by right type, but bad reference
//   tezQuery.GetBlock("lkjs")
//
// }
