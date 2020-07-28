package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Payer type and contains an array of plan
type Pharmacy struct {
	Id           int       `json:"id"`
	PharmacyId   int       `json:"pharmacyId"`
	PharmacyName string    `json:"pharmacyName "`
	Ncpdp        string    `json:"ncpdp"`
	Npi          string    `json:"npi"`
	Phone        string    `json:"phone"`
	Fax          string    `json:"fax"`
	CreateDate   time.Time `json:"createdDate"`
	CreatedBy    string    `json:"createdBy"`
	//LastUpdatedDate time.Time `json:"lastUpdatedDate"`
	//LastUpdatedBy   string    `json:"lastUpdatedBy"`
}

func buildSqlPharmacy(sqltype string, str string, id int) string {
	switch sqltype {
	case "Id":
		return fmt.Sprintf("SELECT id, pharmacy_id, pharmacy_name, ncpdp, npi_nbr, phone_nbr, fax_nbr, created_date, created_by FROM T_NCPDP_DATA WHERE pharmacy_id = %d;", id)
	case "NCPDP":
		return fmt.Sprintf("SELECT id, pharmacy_id, pharmacy_name, ncpdp, npi_nbr, phone_nbr, fax_nbr, created_date, created_by FROM T_NCPDP_DATA WHERE ncpdp ='%s';", str)
	case "NPI":
		return fmt.Sprintf("SELECT id, pharmacy_id, pharmacy_name, ncpdp, npi_nbr, phone_nbr, fax_nbr, created_date, created_by  FROM T_NCPDP_DATA WHERE npi = '%s';", str)
	case "Name":
		return fmt.Sprintf("SELECT id, pharmacy_id, pharmacy_name, ncpdp, npi_nbr, phone_nbr, fax_nbr, created_date, created_by  FROM T_NCPDP_DATA WHERE pharmacy_name ='%s';", str)
	default:
		return ""
	}
}

func searchResultPharmacy(searchName string, str string, id int) (resultSet []Pharmacy, err error) {
	result := make([]Pharmacy, 0)
	sqlStatementPlan := buildSqlPharmacy(searchName, str, id)
	fmt.Println("sqlStatementPlan " + sqlStatementPlan)
	rows, err := db.Query(sqlStatementPlan)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No row was returned!")
		//err = nil
	case nil:
		defer rows.Close()
		//var i = 0
		for rows.Next() {
			var pharmacy Pharmacy
			err = rows.Scan(&pharmacy.Id, &pharmacy.PharmacyId, &pharmacy.PharmacyName, &pharmacy.Ncpdp, &pharmacy.Npi, &pharmacy.Phone, &pharmacy.Fax, &pharmacy.CreateDate, &pharmacy.CreatedBy)
			if err != nil {
				panic(err)
			}
			//fmt.Println(plan)
			result = append(result, pharmacy)
			//	i++
		}
	default:
		panic(err)
	}
	return result, err
}
