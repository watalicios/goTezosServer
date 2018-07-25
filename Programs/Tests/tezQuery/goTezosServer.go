package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: Unit tests for the tezQuery Package
License: MIT
*/

import (
  "github.com/DefinitelyNotAGoat/goTezosServer"
  "fmt"
)


func main(){
  var block goTezosServer.Block
  var err error
  block, err = goTezosServer.GetBlock(10)
  if (err != nil){
    fmt.Println(err)
  }
  fmt.Println(block)
  block, err = goTezosServer.GetBlock("BLq3jarZuxz4F7pamd9xvNmrSstMqAS2JMP5SB2VMqpdWpN88eJ")
  if (err != nil){
    fmt.Println(err)
  }
  fmt.Println(block)
}
