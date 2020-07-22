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
}
