package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	//_ "github.com/gobuffalo/packr"

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
func coalmineJasonHandler(w http.ResponseWriter, r *http.Request) {
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
*/

//var templates = template.Must(template.ParseFiles("Index.tmpl", "Show.tmpl", "New.tmpl", "Edit.tmpl"))
var (
	templates = template.Must(template.ParseFiles("Index.tmpl"))
)

//var templates = template.Must(template.ParseFiles("editTemplate.html", "viewTemplate.html"))

//var templates = template.Must(template.ParseFiles("template/viewTemplate.tmpl"))
/*

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//var tmpl1 = tmpl + ".tmpl"
	err := templates.ExecuteTemplate(w, tmpl, p)
	//tmpl.ExecuteTemplate(w, tmpl+".tmpl", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
*/

var tmpl = template.Must(template.ParseGlob("*.tmpl"))

func Index(w http.ResponseWriter, r *http.Request) {
	var searchType = "PayerId"
	var name = "Anthem Empire BlueCross BlueShield"
	var i2 = 140
	payer, err := searchResultPayer(searchType, name, i2)
	if err != nil {
		panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", payer)
}

func Show(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("Id")
	input, _ := strconv.Atoi(nId)
	var searchType = "PayerId"
	payer, err := searchResultPayer(searchType, "", input)
	if err != nil {
		panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "show", payer)
	//defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "new", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(nId)
	var searchType = "PayerId"
	payer, err := searchResultPayer(searchType, "", id)
	if err != nil {
		panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "edit", payer)
	//defer db.Close()
}

var validPath = regexp.MustCompile("^/(edit|view|error|payer|coal-mine)/([a-zA-Z0-9]+)$")

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

// Error handling types

type errRepoNotInitialized string

func (e errRepoNotInitialized) Error() string {
	return string(e)
}

type errRepoNotFound string

func (e errRepoNotFound) Error() string {
	return string(e)
}

func parseParams(req *http.Request, prefix string, num int) ([]string, error) {
	url := strings.TrimPrefix(req.URL.Path, prefix)
	params := strings.Split(url, "/")

	fmt.Println(url)
	for i := range params {
		fmt.Println(params[i])
	}

	if len(params) != num {
		return nil, fmt.Errorf("Bad format. %d Expecting exactly %d params", len(params), num)
	}
	return params, nil
}

// repoHandler processes the response by parsing the params, then calling
func payerHandler(w http.ResponseWriter, req *http.Request) {
	params, err := parseParams(req, "/payer/", 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	i, err := strconv.Atoi(params[0])
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	var searchType = "PayerId"
	data, err := searchResultPayer(searchType, "", i)
	if err != nil {
		switch err.(type) {
		case errRepoNotFound:
			http.Error(w, err.Error(), 404)
		case errRepoNotInitialized:
			http.Error(w, err.Error(), 401)
		default:
			http.Error(w, err.Error(), 500)
		}
		return
	}
	out, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))

	// var planBool = false //as default
	// if len(data) == 1 && planBool {
	// 	plan, err := searchResultPlan("PlanPayerId", "", data[0].PayerId)
	// 	if err != nil {
	// 		switch err.(type) {
	// 		case errRepoNotFound:
	// 			http.Error(w, err.Error(), 404)
	// 		case errRepoNotInitialized:
	// 			http.Error(w, err.Error(), 401)
	// 		default:
	// 			http.Error(w, err.Error(), 500)
	// 		}
	// 		return
	// 	}
	// 	outPlan, err := json.Marshal(plan)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), 500)
	// 		return
	// 	}
	// 	fmt.Fprintf(w, string(outPlan))
	// }
}

func planHandler(w http.ResponseWriter, req *http.Request) {
	params, err := parseParams(req, "/plan/", 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	i, err := strconv.Atoi(params[0])
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	var searchType = "PlanId"
	data, err := searchResultPlan(searchType, "", i)
	if err != nil {
		switch err.(type) {
		case errRepoNotFound:
			http.Error(w, err.Error(), 404)
		case errRepoNotInitialized:
			http.Error(w, err.Error(), 401)
		default:
			http.Error(w, err.Error(), 500)
		}
		return
	}
	out, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))
}

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

	//http.HandleFunc("/api/index", indexHandler)
	http.HandleFunc("/coal-mine/", coalmineHandler)
	http.HandleFunc("/payer/", payerHandler)
	http.HandleFunc("/payerplan/", payerHandler)
	http.HandleFunc("/plan/", planHandler)
	//http.HandleFunc("/payer/name", payerHandler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))

	// log.Println("Server started on: http://localhost:8081")
	// http.HandleFunc("/coal-mine/", coalmineHandler)
	// http.HandleFunc("/error/", errorHandler)
	// http.HandleFunc("/", Index)
	// http.HandleFunc("/show", Show)
	// http.HandleFunc("/new", New)
	// http.HandleFunc("/edit", Edit)

	//http.HandleFunc("/insert", Insert)
	//http.HandleFunc("/update", Update)
	//http.HandleFunc("/delete", Delete)
	//http.ListenAndServe(":8080", nil)
	//log.Fatal(http.ListenAndServe(":8081", nil))

}
