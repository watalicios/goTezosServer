package goTezosServer

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

type ServiceReport struct {
	Delegate         string           `json:"delegate"`
	Rate             float64          `json:"rate"`
	ContractsByCycle []StructCycle    `json:"contracts_by_cycle"`
	ContractTotals   []ContractReport `json:"contracts_totals"`
	TotalPayouts     float64          `json:"total_payouts"`
	TotalFee         float64          `json:"total_fee"`
	TotalRewards     float64          `json:"total_rewards"`
}

type StructCycle struct {
	Cycle     int              `json:"cycle"`
	Contracts []ContractReport `json:"contracts"`
}

type ContractReport struct {
	Contract     string  `json:"contract"`
	Share        float64 `json:"share"`
	GrossRewards float64 `json:"gross_rewards"`
	Fee          float64 `json:"fee"`
	NetPayout    float64 `json:"net_payout"`
}

var (
	reCycles = regexp.MustCompile(`([0-9]+)`)
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

func ComputeDelegateServiceReport(rate float64, delegateAddr string, cycles string) (ServiceReport, error) {
	var report ServiceReport
	delegate, err := GetDelegateInfo(delegateAddr)
	if err != nil {
		return report, errors.New("Could not find delegate in database.")
	}
	report.Delegate = delegate.Address
	report.Rate = rate

	cycRange := parseCyclesInput(cycles)

	var contractsByCycle []StructCycle
	var rewards float64
	rewards = 0

	for _, contractsBySnapshot := range delegate.ContractsBySnapShot {
		if contractsBySnapshot.Cycle >= cycRange[0] && contractsBySnapshot.Cycle <= cycRange[1] {

			var cycle StructCycle

			fRewards, err := strconv.ParseFloat(contractsBySnapshot.Rewards, 64)
			if err != nil {
				return report, errors.New("Could not find delegate in database.")
			}
			rewards += (fRewards / 1000000)
			fRewards = fRewards / 1000000
			cycle.Cycle = contractsBySnapshot.Cycle
			for _, contracts := range contractsBySnapshot.DelegateContracts {
				var contractReport ContractReport
				contractReport.Share = contracts.Share
				contractReport.Contract = contracts.ContractAddress
				contractReport.Fee = (fRewards * contractReport.Share) * rate
				contractReport.GrossRewards = fRewards * contractReport.Share
				contractReport.NetPayout = contractReport.GrossRewards - contractReport.Fee
				cycle.Contracts = append(cycle.Contracts, contractReport)

			}
			contractsByCycle = append(contractsByCycle, cycle)
		}

	}
	report.ContractsByCycle = contractsByCycle
	report.ContractTotals = calculateContractTotals(contractsByCycle, delegate.DelegateContracts)
	report.TotalPayouts, report.TotalFee = CalculateTotals(report.ContractTotals)
	report.TotalRewards = rewards - report.TotalPayouts
	return report, nil
}

func CalculateTotals(totals []ContractReport) (float64, float64) {
	var payout float64
	var fee float64
	payout = 0
	fee = 0

	for _, total := range totals {
		payout += total.NetPayout
		fee += total.Fee
	}

	return payout, fee
}

func calculateContractTotals(cycles []StructCycle, contracts []string) []ContractReport {
	var contractTotals []ContractReport

	for _, contract := range contracts {
		contractTotals = append(contractTotals, ContractReport{Contract: contract, Share: 0, GrossRewards: 0, Fee: 0, NetPayout: 0})
	}
	for _, cycle := range cycles {
		for _, contract := range cycle.Contracts {
			for index, total := range contractTotals {
				if total.Contract == contract.Contract {
					contractTotals[index].GrossRewards += contract.GrossRewards
					contractTotals[index].NetPayout += contract.NetPayout
					contractTotals[index].Fee += contract.Fee
				}
			}

		}
	}
	return contractTotals
}

func parseCyclesInput(cycles string) [2]int {
	arrayCycles := reCycles.FindAllStringSubmatch(cycles, -1)
	if arrayCycles == nil || len(arrayCycles) > 2 {
		fmt.Println("Unable to parse cycles flag. Example format -cycles=8 or -cycles=8-12.")
	}
	var cycleRange [2]int

	if len(arrayCycles) == 1 {
		cycleRange[0], _ = strconv.Atoi(arrayCycles[0][1])
		cycleRange[1], _ = strconv.Atoi(arrayCycles[0][1])
	} else {
		cycleRange[0], _ = strconv.Atoi(arrayCycles[0][1])
		cycleRange[1], _ = strconv.Atoi(arrayCycles[1][1])
	}

	return cycleRange
}
