package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: A Program to initialize the goTezosServer DB, and keep it updated
License: MIT
*/

import (
  "sync"
  "flag"
  "fmt"
  "github.com/DefinitelyNotAGoat/goTezosServer"
)

var wg sync.WaitGroup

func main(){
	init := flag.Bool("init", true, "Start synchronization of the database from cycle 0")

  flag.Parse()
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
