package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Payer type and contains an array of plan
type Payer struct {
	Id        int    `json:"id"`
	PayerId   int    `json:"payerId"`
	PayerName string `json:"payerName"`
	ParentId  int    `json:"parentId"`
}

type Plan struct {
	Id       int    `json:"id"`
	PayerId  int    `json:"payerId"`
	PlanId   int    `json:"planId"`
	PlanName string `json:"planName"`
	Segment  string `json:"segment"`
	PlanType string `json:"planType"`
	Ppo      string `json:"ppo"`
	Hmo      string `json:"hmo"`
	Pos      string `json:"pos"`
}

// type PayerPlan struct {
// 	payer Payer
// 	plans []Plan
// }

// // check string if needed
// func checkArgs(args []string, del string, fieldLen int) (err error) {
// 	fmt.Printf("[DEBUG] Checking paramters: '%s'\n", strings.Join(args, del))
// 	if fieldLen != len(args) {
// 		err = errors.New(" Invalid field field")
// 	}
// 	return
// }

func buildSql(sqltype string, name string, id int) string {
	switch sqltype {
	case "PayerName":
		return fmt.Sprintf("SELECT id, payer_id, payer_name, parent_id FROM T_DRG_PAYER_DATA WHERE payer_name ='%s';", name)
	case "PayerId":
		return fmt.Sprintf("SELECT id, payer_id, payer_name, parent_id FROM T_DRG_PAYER_DATA WHERE payer_id = %d;", id)
	case "PlanName":
		return fmt.Sprintf("SELECT id, payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos FROM T_DRG_PLAN_DATA WHERE plan_name ='%s';", name)
		//return fmt.Sprintf("SELECT id, payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos FROM plan WHERE plan_name =CONCAT(''%'',CONCAT('%s', '%'));", name)
	case "PlanId":
		return fmt.Sprintf("SELECT id, payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos FROM T_DRG_PLAN_DATA WHERE plan_id = %d;", id)
	case "PlanPayerId":
		return fmt.Sprintf("SELECT id, payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos FROM T_DRG_PLAN_DATA WHERE payer_id = %d;", id)
	default:
		return ""
	}
}

func searchResultPlan(searchName string, name string, id int) (resultSet []Plan, err error) {
	result := make([]Plan, 0)
	sqlStatementPlan := buildSql(searchName, name, id)
	fmt.Println("sqlStatementPlan " + sqlStatementPlan)
	rows, err := db.Query(sqlStatementPlan)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No row was returned!")
		//err = nil
	case nil:
		defer rows.Close()
		for rows.Next() {
			var plan Plan
			err = rows.Scan(&plan.Id, &plan.PayerId, &plan.PlanId, &plan.PlanName, &plan.Segment, &plan.PlanType, &plan.Ppo, &plan.Hmo, &plan.Pos)
			if err != nil {
				panic(err)
			}
			result = append(result, plan)
		}
	default:
		panic(err)
	}
	return result, err
}

func searchResultPayer(searchName string, name string, id int) (resultSet []Payer, err error) {
	result := make([]Payer, 0)
	sqlStatementPlan := buildSql(searchName, name, id)
	fmt.Println("sqlStatementPlan " + sqlStatementPlan)
	rows, err := db.Query(sqlStatementPlan)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No row was returned!")
		//err = nil
	case nil:
		defer rows.Close()
		for rows.Next() {
			var payer Payer
			err = rows.Scan(&payer.Id, &payer.PayerId, &payer.PayerName, &payer.ParentId)
			if err != nil {
				panic(err)
			}
			result = append(result, payer)
		}
		//for debug , remove later
		if nil == err {
			//fmt.Println(plan)
			for i := range result {
				fmt.Println(result[i])
			}
		}
	default:
		panic(err)
	}
	return result, err
}
