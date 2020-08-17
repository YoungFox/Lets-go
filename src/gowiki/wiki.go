package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

// Page struct
// 123
type Page struct {
	Title string
	Body  []byte
	Links []string
}

var tmplDir = "tmpl/"
var dataDir = "data/"

var templates = template.Must(template.ParseFiles(tmplDir+"edit.html", tmplDir+"view.html"))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (p *Page) save() error {
	filename := dataDir + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := dataDir + title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}

	return m[2], nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "view/FrontPage", http.StatusFound)
}

func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/FrontPage.html")
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	linkList := loadList()
	p.Links = linkList

	renderTemplate(w, "view", p)
}

func loadList() []string {
	var s []string
	files, _ := ioutil.ReadDir("data/")

	for _, file := range files {
		s = append(s, file.Name())
	}
	replaceNameToLinkElement(s)
	return s
}

func replaceNameToLinkElement(x []string) {
	re := regexp.MustCompile(`^([a-zA-Z0-9]+)\.txt$`)
	for k, v := range x {
		// a := re.ReplaceAllStringFunc(v, replFunc)
		b := re.FindStringSubmatch(v)
		x[k] = replFunc(b[1])
	}
}
func replFunc(s string) string {
	return `<a href="/view/` + s + `">` + s + `</a>`
}
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	// The value returned by FormValue is of type string.
	// We must convert that value to []byte before it will fit into the Page struct.
	//  We use []byte(body) to perform the conversion.
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

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

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/FrontPage", frontPageHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
