package goTezosServer

/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: The Tezos API written in GO, for easy development.
License: MIT
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

var (
	reRewards = regexp.MustCompile(`"rewards": "([0-9]+)"`)
	reStake   = regexp.MustCompile(`"([0-9]+)"`)
)
var snapShot []SnapShot

func GetSnapShot(cycle int) (SnapShot, error) {
	var snapShot SnapShot
	snapShot.Cycle = cycle
	strCycle := strconv.Itoa(cycle)

	snapshotStr := "/chains/main/blocks/head/context/raw/json/cycle/" + strCycle

	s, err := TezosRPCGet(snapshotStr)
	if err != nil {
		return snapShot, errors.New("func GetSnapShot(cycle int) (SnapShot, error) failed: " + err.Error())
	}

	regRandomSeed := reGetRandomSeed.FindStringSubmatch(string(s[:]))
	if regRandomSeed == nil {
		return snapShot, errors.New("No random seed: func GetSnapShot(cycle int) (SnapShot, error) failed.")
	}

	regRollSnapShot := reGetRollSnapShot.FindStringSubmatch(string(s[:]))
	if regRollSnapShot == nil {
		return snapShot, errors.New("Could not parse snapshot: func GetSnapShot(cycle int) (SnapShot, error) failed.")
	}
	number, _ := strconv.Atoi(regRollSnapShot[1])
	snapShot.Number = number
	snapShot.AssociatedBlock = ((cycle - 7) * 4096) + (number+1)*256
	snapShot.AssociatedHash, _ = GetBlockLevelHash(snapShot.AssociatedBlock)

	return snapShot, nil
}

func InitDelegateDB() error {
	list, err := GetAllDelegates()
	if err != nil {
		fmt.Println(err)
	}

	curCycle, err := GetCurrentCycle()
	if err != nil {
		return errors.New("Could not get contracts for cycles: " + string(curCycle))
	}

	lCycle := curCycle + 5
	for i := 0; i <= lCycle; i++ {
		tmpSnapShot, err := GetSnapShot(i)
		if err == nil {
			snapShot = append(snapShot, tmpSnapShot)
		}
	}
	count := len(list)
	bar := pb.StartNew(count)
	for _, del := range list {
		GetDelegate(del, lCycle)
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("The End!")
	return nil
}

func GetDelegate(delegateAddr string, cycle int) error {
	var delegate StructDelegate
	get := "chains/main/blocks/head/context/delegates/" + delegateAddr
	s, err := TezosRPCGet(get)
	if err != nil {
		return errors.New("Could not get delegate " + delegateAddr)
	}

	delegate = ConvertBytestoDelegate(s)
	delegate.Address = delegateAddr

	bal, _ := strconv.Atoi(delegate.DelegatedBalance)
	if bal != 0 {
		contractsByCycle, err := GetContractsBySnapshot(delegateAddr, cycle)
		if err != nil {
			return errors.New("Could not get delegate " + delegateAddr)
		}
		delegate.ContractsBySnapShot = contractsByCycle
	}

	err = Collection.Insert(delegate)

	return nil
}

func GetContractsBySnapshot(delegateAddr string, cycle int) ([]StructContractsBySnapShot, error) {
	var contractsBySnapShot []StructContractsBySnapShot
	for _, shot := range snapShot {

		contracts, err := GetDelegatedContractsForCycle(shot.AssociatedHash, delegateAddr)
		if err == nil {
			strStake := strings.TrimSpace(GetStakingBalanceAtCycle(shot.Cycle, delegateAddr))
			stake, err := strconv.Atoi(strStake)
			if err != nil {
				return contractsBySnapShot, errors.New("Could not get contracts by cycle!")
			}
			var delegateContracts []StructDelegateContracts
			if len(contracts) > 0 {
				for _, contract := range contracts {
					balance, err := GetAccountBalanceAtBlock(contract, shot.AssociatedHash)
					if err != nil {
						return contractsBySnapShot, errors.New("Could not get contracts by cycle!")
					}

					share := float64(balance) / float64(stake)
					delegateContracts = append(delegateContracts, StructDelegateContracts{ContractAddress: contract, Balance: balance, Share: share})
				}
				rewards := GetRewardsForCycle(shot.Cycle, delegateAddr)
				contractsBySnapShot = append(contractsBySnapShot, StructContractsBySnapShot{Cycle: shot.Cycle, Rewards: rewards, DelegateContracts: delegateContracts})
			}
		}
	}
	return contractsBySnapShot, nil
}

func GetStakingBalanceAtCycle(cycle int, delegateAddr string) string {

	var hash string
	for _, shot := range snapShot {
		if shot.Cycle == cycle {
			hash = shot.AssociatedHash
		}
	}
	get := "/chains/main/blocks/" + hash + "/context/delegates/" + delegateAddr + "/staking_balance"
	s, err := TezosRPCGet(get)
	if err != nil {
		return ""
	}

	match := reStake.FindStringSubmatch(string(s[:]))
	if match == nil {
		return ""
	}
	return match[1]
}

func GetRewardsForCycle(cycle int, delegateAddr string) string {

	get := "/chains/main/blocks/head/context/raw/json/contracts/index/" + delegateAddr + "/frozen_balance/" + strconv.Itoa(cycle) + "/"
	s, err := TezosRPCGet(get)
	if err != nil {
		return ""
	}
	res := string(s[:])
	match := reRewards.FindStringSubmatch(res)
	if match == nil {
		return ""
	}

	return match[1]
}

func GetCurrentCycle() (int, error) {
	s, err := GetBlockRPCHead()
	if err != nil {
		return 0, errors.New("Could not get Cycle: " + err.Error())
	}
	block := ConvertBlockToJson(s)
	var cycle int
	cycle = block.Header.Level / 4096

	return cycle, nil
}

/*
Description: Takes an  array of interface (struct in our case), jsonifies it, and allows a much neater print.
Param v (interface{}): Array of an interface
*/
func ConvertBytestoDelegate(v []byte) StructDelegate {
	var delegate StructDelegate
	err := json.Unmarshal(v, &delegate)
	if err != nil {
		panic(err)
	}

	return delegate
}

/*
Description: Will get the balance of an account at a specific snapshot
Param tezosAddr (string): Takes a string representation of the address querying
Param cycle (int): The cycle we are getting the snapshot for
Returns (float64): Returns a float64 representation of the balance for the account
*/
func GetAccountBalanceAtBlock(tezosAddr string, hash string) (int, error) {

	balanceCmdStr := "/chains/main/blocks/" + hash + "/context/contracts/" + tezosAddr + "/balance"

	s, err := TezosRPCGet(balanceCmdStr)
	if err != nil {
		return 0, errors.New("Could not get balance at snapshot for " + tezosAddr + ": TezosRPCGet(arg string) failed: " + err.Error())
	}

	var returnBalance int
	var regGetBalance []string

	if strings.Contains(string(s[:]), "No service found at this URL") {
		returnBalance = 0
	} else {
		regGetBalance = reGetBalance.FindStringSubmatch(string(s[:]))
		if regGetBalance == nil {
			return 0, errors.New("Could not parse balance for " + string(s[:]))
		}
	}

	if len(regGetBalance) < 1 {
		returnBalance = 0
	} else {
		floatBalance, _ := strconv.Atoi(regGetBalance[1]) //TODO error checking
		returnBalance = int(floatBalance)
	}

	return returnBalance, nil
}

/*
Description: Will retreive the current block level as an integer
Returns (int): Returns integer representation of block level
*/
func GetBlockLevelHead() (int, string, error) {
	s, err := TezosRPCGet("chains/main/blocks/head")
	if err != nil {
		return 0, "", errors.New("func GetBlockLevelHead() failed: " + err.Error())
	}

	regHeadLevelResult := reGetBlockLevelHead.FindStringSubmatch(string(s[:]))
	if regHeadLevelResult == nil {
		return 0, "", errors.New("Could not parse head level: func GetBlockLevelHead() failed.")
	}
	regHash := reGetHash.FindStringSubmatch(string(s[:]))
	if regHash == nil {
		return 0, "", errors.New("Could not parse head hash: func GetBlockLevelHead() failed.")
	}
	headlevel, _ := strconv.Atoi(regHeadLevelResult[1]) //TODO Error Checking

	return headlevel, regHash[1], nil
}

/*
Description: Takes a block level, and returns the hash for that specific level
Param level (int): An integer representation of the block level to query
Returns (string): A string representation of the hash for the block level queried.
*/
func GetBlockLevelHash(level int) (string, error) {
	head, headHash, err := GetBlockLevelHead()
	if err != nil {
		return "", errors.New("func GetBlockLevelHash(level int) failed: " + err.Error())
	}
	diff := head - level

	diffStr := strconv.Itoa(diff)
	getBlockByLevel := "chains/main/blocks/" + headHash + "~" + diffStr

	s, err := TezosRPCGet(getBlockByLevel)
	if err != nil {
		return "", errors.New("func GetBlockLevelHash(level int) failed: " + err.Error())
	}

	hash := reGetHash.FindStringSubmatch(string(s[:])) //TODO Error check the regex
	if hash == nil {
		return "", errors.New("Could not parse hash: func GetBlockLevelHash(level int) failed.")
	}

	return hash[1], nil
}

/*
Description: Retrieves the list of addresses delegated to a delegate
Param SnapShot: A SnapShot object describing the desired snap shot.
Param delegateAddr: A string that represents a delegators tz address.
Returns []string: An array of contracts delegated to the delegator during the snap shot
*/
func GetDelegatedContractsForCycle(hash string, delegateAddr string) ([]string, error) {
	var rtnString []string
	getDelegatedContracts := "/chains/main/blocks/" + hash + "/context/delegates/" + delegateAddr + "/delegated_contracts"

	s, err := TezosRPCGet(getDelegatedContracts)
	if err != nil {
		return rtnString, errors.New("Could not get delegated contracts!")
	}

	DelegatedContracts := reDelegatedContracts.FindAllStringSubmatch(string(s[:]), -1)
	if DelegatedContracts == nil {
		return rtnString, errors.New("Could not get delegated contracts!")
	}
	rtnString = addressesToArray(DelegatedContracts)
	return rtnString, nil
}

/*
Description: Retrieves the list of addresses delegated to a delegate
Param SnapShot: A SnapShot object describing the desired snap shot.
Param delegateAddr: A string that represents a delegators tz address.
Returns []string: An array of contracts delegated to the delegator during the snap shot
*/
func GetAllDelegates() ([]string, error) {
	var rtnString []string
	getDelegatedContracts := "/chains/main/blocks/head/context/delegates?active"

	s, err := TezosRPCGet(getDelegatedContracts)
	if err != nil {
		return rtnString, errors.New("Could not get delegates!")
	}

	DelegatedContracts := reDelegatedContracts.FindAllStringSubmatch(string(s[:]), -1)
	if DelegatedContracts == nil {
		return rtnString, errors.New("Could not get delegates!")
	}
	rtnString = addressesToArray(DelegatedContracts)
	return rtnString, nil
}

/*
Description: Takes a multi-dimmensional array of addresses from a regex parse, and converts them into a single index(able) array.
Param matches ([][]string): All the addresses found and parsed by regex (ex. DelegatedContracts := reDelegatedContracts.FindAllStringSubmatch(s, -1) returns a multi dimmensional array)
Returns ([]string): Returns an index(able) string array of the matches input.
*/
func addressesToArray(matches [][]string) []string {
	var addresses []string
	for _, x := range matches {
		addresses = append(addresses, x[1])
	}

	return addresses
}
