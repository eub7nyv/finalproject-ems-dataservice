package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

// Payer type and contains an array of plan
type Payer struct {
	id        int
	payerId   int
	payerName string
	parentId  int
}

type Plan struct {
	id       int
	payerId  int
	planId   int
	planName string
	segment  string
	planType string
	ppo      string
	hmo      string
	pos      string
}

type PayerPlan struct {
	payer Payer
	plans []Plan
}

// HiString constant to say hi
const HiString string = "Hello "

func checkArgs(args []string, del string, fieldLen int) (err error) {
	fmt.Printf("[DEBUG] Checking paramters: '%s'\n", strings.Join(args, del))
	if fieldLen != len(args) {
		err = errors.New(" Invalid field field")
	}
	return
}

func buildSql(sqltype string, name string, id int) string {
	switch sqltype {
	case "PayerName":
		return fmt.Sprintf("SELECT id, payer_id, payer_name, parent_id FROM payer WHERE payer_name ='%s';", name)
	case "PayerId":
		return fmt.Sprintf("SELECT id, payer_id, payer_name, parent_id FROM payer WHERE payer_id = %d;", id)
	case "PlanName":
		return fmt.Sprintf("SELECT id, payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos FROM plan WHERE plan_name ='%s';", name)
	case "PlanId":
		return fmt.Sprintf("SELECT id, payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos FROM plan WHERE plan_id = %d;", id)
	case "PlanPayerId":
		return fmt.Sprintf("SELECT id, payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos FROM plan WHERE payer_id = %d;", id)
	default:
		return ""
	}
}

/*
func searchPayer(searchName string, name string, id int) (result []Payer, err error){
	//build query based on searchName, searchString
			sqlStatementPlan := buildSql(searchName, name, id)
			fmt.Println("sqlStatementPlan " + sqlStatementPlan)
			rows, err := db.Query(sqlStatementPlan)
			switch err {
			case sql.ErrNoRows:
				fmt.Println("No row was returned!")
				return "", err
			case nil:
				defer rows.Close()
				for rows.Next() {
					var i2 = 0
					var payer2 Payer
					err = rows.Scan(&payer2.id, &payer2.payerId, &payer2.payerName, &payer2.parentId)
					if err != nil {
						panic(err)
					}
					fmt.Println(payer2)
					result[i2].payer = payer2
					i2++
				}
				// get any error encountered during iteration
				err = rows.Err()
				if err != nil {
					panic(err)
				}
			default:
				panic(err)

			return result, err
}
*/

func searchResultPlan(searchName string, name string, id int) (resultSet []Plan, err error) {
	//setup DB connection

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//open a connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	result := make([]Plan, 1)

	//build query based on searchName, searchString
	// get plan
	sqlStatementPlan := buildSql(searchName, name, id)
	fmt.Println("sqlStatementPlan " + sqlStatementPlan)
	rows, err := db.Query(sqlStatementPlan)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No row was returned!")
		err = nil
	case nil:
		defer rows.Close()
		for rows.Next() {
			var i = 0
			var plan Plan
			err = rows.Scan(&plan.id, &plan.payerId, &plan.planId, &plan.planName, &plan.segment, &plan.planType, &plan.ppo, &plan.hmo, &plan.pos)
			if err != nil {
				panic(err)
			}
			fmt.Println(plan)
			result[0] = plan
			i++
		}
	default:
		panic(err)
	}
	return result, err
}

func searchResult(searchName string, name string, id int) (resultSet []Payer, err error) {
	//setup DB connection

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//open a connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	result := make([]Payer, 5)

	//build query based on searchName, searchString
	switch searchName {
	case "PayerId":
		{
			sqlStatement := buildSql(searchName, name, id)
			fmt.Println("payerInfo " + sqlStatement)

			var payer Payer
			row := db.QueryRow(sqlStatement)
			err := row.Scan(&payer.id, &payer.payerId, &payer.payerName, &payer.parentId)
			switch err {
			case sql.ErrNoRows:
				fmt.Println("No row was returned!")
				err = nil
			case nil:

				fmt.Println(payer)
				result[0] = payer
				/*
					// get plan
					sqlStatementPlan := buildSql("PlanPayerId", "", result[0].payer.payerId)
					fmt.Println("sqlStatementPlan " + sqlStatementPlan)
					rows, err := db.Query(sqlStatementPlan)
					if err != nil {
						// handle this error better than this
						panic(err)
					}
					defer rows.Close()
					for rows.Next() {
						var i = 0
						var plan Plan
						err = rows.Scan(&plan.id, &plan.payerId, &plan.planId, &plan.planName, &plan.segment, &plan.planType, &plan.ppo, &plan.hmo, &plan.pos)
						if err != nil {
							panic(err)
						}
						fmt.Println(plan)
						result[0].plans[i] = plan
						i++
					}
					// get any error encountered during iteration
					err = rows.Err()
					if err != nil {
						panic(err)
					}
				*/
			default:
				panic(err)

			}
			//resultSet = result
			return result, err
		}
	case "PayerName":
		{
			// get plan
			sqlStatementPlan := buildSql("PayerName", name, id)
			fmt.Println("sqlStatementPlan " + sqlStatementPlan)
			rows, err := db.Query(sqlStatementPlan)
			if err != nil {
				// handle this error better than this
				panic(err)
			}
			defer rows.Close()
			for rows.Next() {
				var i2 = 0
				var payer2 Payer
				err = rows.Scan(&payer2.id, &payer2.payerId, &payer2.payerName, &payer2.parentId)
				if err != nil {
					panic(err)
				}
				fmt.Println(payer2)
				//result[i2].payer = payer2
				result[i2] = payer2
				i2++
			}
			// get any error encountered during iteration
			err = rows.Err()
			if err != nil {
				panic(err)
			}
		}
	default:
		return result, err
	}
	return result, err
}
