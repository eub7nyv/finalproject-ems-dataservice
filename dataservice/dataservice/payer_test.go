package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSearchResult(t *testing.T) {

	var searchType = "PayerId"
	var name = ""
	var i2 = 140
	var payer []Payer
	payer, err := searchResult(searchType, name, i2)
	if nil != err {
		t.Error(err)
	}
	if len(payer) < 0 {
		t.Errorf("TestSearchResult is off")
	}

	searchType = "PayerName"
	name = "Anthem Empire BlueCross BlueShield"
	i2 = 140
	//payerInfo := buildSql("Anthem Empire BlueCross BlueShield")
	//payer []Payer
	payer2, err := searchResult(searchType, name, i2)
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
	if planId == plan[0].planId {
		fmt.Println("id " + strconv.Itoa(planId) + " matched expection id from DB " + strconv.Itoa(plan[0].planId))
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
	if planPayerId == plan2[0].payerId {
		fmt.Println("input " + strconv.Itoa(planPayerId) + " matches expected from DB " + strconv.Itoa(plan[0].payerId))
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
	if planName == plan3[0].planName {
		fmt.Println("input " + planName + " matches expected from DB " + plan3[0].planName)
	}
}
