package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSearchResultPharmacy(t *testing.T) {

	var searchType = "Id"
	var name = ""
	var pharmacyId = 14650
	//var pharmacy []Pharmacy
	pharmacy, err := searchResultPharmacy(searchType, name, pharmacyId)
	if nil == err {
		//fmt.Println(plan)
		for i := range pharmacy {
			fmt.Println(pharmacy[i])
		}
	}
	if nil != err {
		t.Error(err)
	}
	if pharmacyId == pharmacy[0].PharmacyId {
		fmt.Println("Pharmacy id " + strconv.Itoa(pharmacyId) + " matched expection id from DB " + strconv.Itoa(pharmacy[0].PharmacyId))
	}

	searchType = "Name"
	name = "ACCREDO"
	pharmacyName, err2 := searchResultPharmacy(searchType, name, 0)
	if nil == err2 {
		for i := range pharmacyName {
			fmt.Println(pharmacyName[i])
			if name == pharmacyName[i].PharmacyName {
				fmt.Println("input " + name + " matches expected from DB " + pharmacyName[i].PharmacyName)
			}
		}
	}
	if nil != err2 {
		fmt.Println(err)
	}
}
