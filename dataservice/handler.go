package main

import (
	"encoding/json"
	"fmt"

	//"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	//_ "github.com/gobuffalo/packr"

	_ "github.com/lib/pq"
)

type Page struct {
	Title string
	Body  []byte
}

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

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()         // Parses the request body
	x := r.Form.Get("id") // x will be "" if parameter is not set
	fmt.Println("test parameter" + x)
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

	// req.ParseForm()          // Parses the request body
	// id := req.Form.Get("id") // x will be "" if parameter is not set
	// name := req.Form.Get("name")
	// fmt.Println("test parameter" + id)
	// fmt.Println("test parameter" + name)

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

func pharmacyHandler(w http.ResponseWriter, req *http.Request) {
	params, err := parseParams(req, "/pharmacy/", 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	i, err := strconv.Atoi(params[0])
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	var searchType = "Id"
	data, err := searchResultPharmacy(searchType, "", i)
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

//UI section
var tmpl = template.Must(template.ParseGlob("*.tmpl"))

// func run() error {
// 	info, err := pkger.Stat("./template/*")
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func payerIndexHandler(w http.ResponseWriter, r *http.Request) {
	var searchType = "PayerId"
	var name = "Anthem Empire BlueCross BlueShield"
	var i2 = 140
	payer, err := searchResultPayer(searchType, name, i2)
	if err != nil {
		//panic(err.Error())
		tmpl.ExecuteTemplate(w, "Error", "")
	}
	tmpl.ExecuteTemplate(w, "Index", payer)
}

// func payerViewHandler(w http.ResponseWriter, r *http.Request) {
// 	nId := r.URL.Query().Get("Id")
// 	input, _ := strconv.Atoi(nId)
// 	var searchType = "PayerId"
// 	payer, err := searchResultPayer(searchType, "", input)
// 	if err != nil {
// 		//panic(err.Error())
// 		tmpl.ExecuteTemplate(w, "Error", err.Error())
// 	}
// 	tmpl.ExecuteTemplate(w, "Index", payer)
// }

func payerShowHandler(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("Id")
	input, _ := strconv.Atoi(nId)
	var searchType = "PayerId"
	payer, err := searchResultPayer(searchType, "", input)
	if err != nil {
		//panic(err.Error())
		tmpl.ExecuteTemplate(w, "Error", err.Error())
	}
	tmpl.ExecuteTemplate(w, "Show", payer)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Error", "")
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Header", "")
}

func payerEditHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Edit", "")
}

var validPath = regexp.MustCompile("^/(edit|save|view|error|payer|plan|pharmacy)/([a-zA-Z0-9]+)$")

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

func renderTemplate(w http.ResponseWriter, tmplName string, p *Page) {
	err := tmpl.ExecuteTemplate(w, tmplName+".tmpl", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "Edit", p)
}
