package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: Unit tests for the tezQuery Package
License: MIT
*/

import (
  "github.com/DefinitelyNotAGoat/goTezosServer"
  "github.com/DefinitelyNotAGoat/goTezosServer/tezQuery"
  "fmt"
)


func main(){
  var block goTezosServer.Block
  var err error
  block, err = tezQuery.GetBlock(10)
  if (err != nil){
    fmt.Println(err)
  }
  fmt.Println(block)
  block, err = tezQuery.GetBlock("BLq3jarZuxz4F7pamd9xvNmrSstMqAS2JMP5SB2VMqpdWpN88eJ")
  if (err != nil){
    fmt.Println(err)
  }
  fmt.Println(block)
}
