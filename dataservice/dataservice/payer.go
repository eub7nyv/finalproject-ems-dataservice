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

// check string if needed
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
		//return fmt.Sprintf("SELECT id, payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos FROM plan WHERE plan_name =CONCAT(''%'',CONCAT('%s', '%'));", name)
	case "PlanId":
		return fmt.Sprintf("SELECT id, payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos FROM plan WHERE plan_id = %d;", id)
	case "PlanPayerId":
		return fmt.Sprintf("SELECT id, payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos FROM plan WHERE payer_id = %d;", id)
	default:
		return ""
	}
}

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

	sqlStatementPlan := buildSql(searchName, name, id)
	fmt.Println("sqlStatementPlan " + sqlStatementPlan)
	rows, err := db.Query(sqlStatementPlan)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No row was returned!")
		err = nil
	case nil:
		defer rows.Close()
		var i = 0
		for rows.Next() {
			var plan Plan
			err = rows.Scan(&plan.id, &plan.payerId, &plan.planId, &plan.planName, &plan.segment, &plan.planType, &plan.ppo, &plan.hmo, &plan.pos)
			if err != nil {
				panic(err)
			}
			//fmt.Println(plan)
			result[0] = plan
			i++
		}
	default:
		panic(err)
	}
	return result, err
}

func searchResultPayer(searchName string, name string, id int) (resultSet []Payer, err error) {
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

	result := make([]Payer, 2)

	sqlStatementPlan := buildSql(searchName, name, id)
	fmt.Println("sqlStatementPlan " + sqlStatementPlan)
	rows, err := db.Query(sqlStatementPlan)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No row was returned!")
		err = nil
	case nil:
		defer rows.Close()
		var i = 0
		for rows.Next() {
			var payer Payer
			err = rows.Scan(&payer.id, &payer.payerId, &payer.payerName, &payer.parentId)
			if err != nil {
				panic(err)
			}
			//fmt.Println(payer)
			result[i] = payer
			i++
		}

	default:
		panic(err)
	}
	return result, err
}
