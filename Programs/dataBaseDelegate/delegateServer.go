package main

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: A Program to initialize the goTezosServer DB, and keep it updated
License: MIT
*/

import (
	"fmt"

	"github.com/DefinitelyNotAGoat/goTezosServer"
)

func main() {
	goTezosServer.GetDelegate("tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc")
	snap, err := goTezosServer.GetSnapShot(9)
	if err != nil {
		fmt.Println(err)
	}
	goTezosServer.PrettyReport(snap)
}
