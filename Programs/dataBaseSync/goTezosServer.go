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
	init := flag.Bool("init", false, "Start synchronization of the database from cycle 0")
  connection := flag.String("connection", "127.0.0.1", "URL or IP to the MongoDB Database")
  db := flag.Bool("db", "TEZOS", "Use the TEZOS Database")
  collection := flag.String("collection", "blocks", "Use the blocks collection")
  user := flag.String("user", "", "If using authentication, set the user")
  pass := flag.String("pass", "", "If using authentication, set the password")

  flag.Parse()

  var dbCon string

  if (*user != "" && *pass != ""){
    dbCon = "mongodb://" + *user + ":" + *pass + "@" + connection
  } else{
    dbCon = "mongodb://"+ connection
  }
  
  SetDatabaseConnection(dbCon, *database, *collection)

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
