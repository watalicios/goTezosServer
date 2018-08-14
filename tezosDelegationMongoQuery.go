package goTezosServer

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func GetDelegateInfo(delegateAddr string) (StructDelegate, error) {
	var result StructDelegate
	err := Collection.Find(bson.M{"address": delegateAddr}).One(&result)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	return result, nil
}
