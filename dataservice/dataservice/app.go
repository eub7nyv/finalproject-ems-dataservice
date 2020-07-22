package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	_ "github.com/lib/pq"
)

type Page struct {
	Title string
	Body  []byte
}

/*

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
*/

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func coalmineHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Error Page!")
}

/*
func coalmineJasonHandler(w http.ResponseWriter, r *http.Request, user User) {
	fmt.Fprintf(w, "Tweet! %s\n")
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err) // should redirect error page
		//http.Redirect(w, r, "/error/"+title, http.StatusFound)
		return

	}
	//fmt.Fprintf(w, os.Stdout.Write(b), "")
	fmt.Fprintf(w, string(b), "")
}
*/

/*
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		//http.Redirect(w, r, "/error/"+title, http.StatusFound)
		errorHandler(w, r)
		return
	}
	renderTemplate(w, "viewTemplate", p)
}
*

var templates = template.Must(template.ParseFiles("viewTemplate.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
*/

var validPath = regexp.MustCompile("^/(edit|view|error|coal-mine)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

/*
func main() {
	http.HandleFunc("/coal-mine/", coalmineHandler)
	http.HandleFunc("/error/", errorHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
*/

func getPayer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

const (
	//host = "localhost"
	host     = "localhost"
	port     = 5432
	user     = "newuser"
	password = "password"
	dbname   = "payer"
)

func main() {
	//connStr := "postgres://newuser:password@localhost:5432/payer?sslmode=disable"
	/*
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	*/
	//db, err := sql.Open("postgres", psqlInfo)
	/*
		var searchType = "PayerId"
		var name = "Anthem Empire BlueCross BlueShield"
		var i2 = 140
		//payerInfo := buildSql("Anthem Empire BlueCross BlueShield")
		var payer []Payer
		payer, err := searchResult(searchType, name, i2)

		if nil == err {
			fmt.Println(payer[0].payerId)
			for i := range payer[1:] {
				fmt.Println(payer[i].payerId)
			}
		}

		searchType = "PayerName"
		name = "Anthem Empire BlueCross BlueShield"
		i2 = 140
		//payerInfo := buildSql("Anthem Empire BlueCross BlueShield")
		//payer []Payer
		payer2, err2 := searchResult(searchType, name, i2)

		if nil == err2 {
			fmt.Println(payer2[0].payerId)
			for i := range payer2[1:] {
				fmt.Println(payer2[i].payerId)
			}
		}
	*/
	var searchType3 = "PlanId"
	var planId = 17629
	plan, err := searchResultPlan(searchType3, "", planId)
	if nil == err {
		//fmt.Println(plan)
		for i := range plan {
			fmt.Println(plan[i])
		}
	}
	if nil != err {
		fmt.Println("error")
		fmt.Println(err)
	}
	if planId == plan[0].planId {
		fmt.Println("id " + strconv.Itoa(planId) + " matched expection id from DB " + strconv.Itoa(plan[0].planId))
	}

	var searchType = "PlanPayerId"
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

	fmt.Println("Successfully done!")
}
