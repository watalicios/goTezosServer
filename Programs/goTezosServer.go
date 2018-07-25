package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: The Tezos API written in GO, for easy development.
License: MIT
*/

import (
  "github.com/DefinitelyNotAGoat/goTezosServer"
)


func main(){
  goTezosServer.SynchronizeTezosMongo()
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
