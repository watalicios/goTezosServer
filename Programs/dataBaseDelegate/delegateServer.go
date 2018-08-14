package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: A Program to initialize the goTezosServer DB, and keep it updated
License: MIT
*/

import (
	"github.com/DefinitelyNotAGoat/goTezosServer"
)

func main() {
	goTezosServer.SetDatabaseConnection("127.0.0.1", "TEZOS", "delegates")
	goTezosServer.InitDelegateDB()
}
