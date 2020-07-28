package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSearchResultPayer(t *testing.T) {

	var searchType = "PayerId"
	var name = ""
	var i2 = 140
	var payer []Payer
	payer, err := searchResultPayer(searchType, name, i2)
	if nil != err {
		t.Error(err)
	}
	if len(payer) < 0 {
		t.Errorf("TestSearchResult is off")
	}

	searchType = "PayerName"
	name = "Anthem Empire BlueCross BlueShield"
	i2 = 0
	payer2, err := searchResultPayer(searchType, name, i2)
	if nil != err {
		t.Error(err)
	}
	if len(payer2) < 0 {
		t.Errorf("TestSearchResult is off")
	}

}

func TestSearchResultPlan(t *testing.T) {
	var searchType = "PlanId"
	var planId = 17629
	plan, err := searchResultPlan(searchType, "", planId)
	if nil == err {
		//fmt.Println(plan)
		for i := range plan {
			fmt.Println(plan[i])
		}
	}
	if nil != err {
		t.Error(err)
	}
	if planId == plan[0].PlanId {
		fmt.Println("id " + strconv.Itoa(planId) + " matched expection id from DB " + strconv.Itoa(plan[0].PlanId))
	}

	searchType = "PlanPayerId"
	var planPayerId = 140
	plan2, err2 := searchResultPlan(searchType, "", planPayerId)
	if nil == err {
		//fmt.Println(plan)
		for i := range plan2 {
			fmt.Println(plan2[i])
		}
	}
	if nil != err2 {
		fmt.Println(err2)
	}
	if planPayerId == plan2[0].PayerId {
		fmt.Println("input " + strconv.Itoa(planPayerId) + " matches expected from DB " + strconv.Itoa(plan2[0].PayerId))
	}

	searchType = "PlanName"
	var planName = "Blue Access"
	plan3, err3 := searchResultPlan(searchType, planName, 0)
	if nil == err {
		//fmt.Println(plan)
		for i := range plan3 {
			fmt.Println(plan3[i])
		}
	}
	if nil != err3 {
		fmt.Println(err)
	}
	if planName == plan3[0].PlanName {
		fmt.Println("input " + planName + " matches expected from DB " + plan3[0].PlanName)
	}
}
