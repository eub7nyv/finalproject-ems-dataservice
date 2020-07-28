package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	//_ "github.com/gobuffalo/packr"

	_ "github.com/lib/pq"
)

func main() {
	//var tmpl = template.Must(template.ParseGlob("*.tmpl"))
	// func run() error {
	// 	info, err := pkger.Stat("./template/*")
	// 	if err != nil {
	// 		return err
	// 	}

	// fmt.Println("Name: ", info.Name())
	// fmt.Println("Size: ", info.Size())

	// Tell pkger that it has to package that directory.
	//dir := pkger.Include("/templates")
	// Only compile templates on startup.
	//tpl, _ := compileTemplates(dir)
	//fmt.Println(tpl)

	// //box := packr.NewBox("/templates")
	// t, err := template.New("hello").Parse(box.String("hello.txt"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = t.Execute(os.Stdout, "world")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	initDb()
	defer db.Close()

	var searchType = "PayerId"
	var name = ""
	var i2 = 140
	var payer []Payer
	payer, err := searchResultPayer(searchType, name, i2)
	if nil == err {
		//fmt.Println(plan)
		for i := range payer {
			fmt.Println(payer[i])
		}
	}
	if nil != err {
		fmt.Println(err)
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

	http.HandleFunc("/coal-mine/", coalmineHandler)
	http.HandleFunc("/payer/", payerHandler)
	http.HandleFunc("/plan/", planHandler)
	http.HandleFunc("/pharmacy/", pharmacyHandler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))

	// log.Println("Server started on: http://localhost:8081")
	// http.HandleFunc("/error/", errorHandler)
	// http.HandleFunc("/", Index)
	// http.HandleFunc("/show", Show)
	// http.HandleFunc("/new", New)
	// http.HandleFunc("/edit", Edit)
	//http.ListenAndServe(":8080", nil)
	//log.Fatal(http.ListenAndServe(":8081", nil))

}
