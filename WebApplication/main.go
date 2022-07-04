package main

import (
	"fmt"

	"os"

	"net/http"

	"log"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("it is handled")
	fmt.Fprintf(w, "Hi there, you f**cking %s, :)", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	wikiPage, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, "<h1>Not Found</h1><div>Your requested file does not exist</div>")
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", wikiPage.Title, wikiPage.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1><form action=\"/save/%s\" method=\"POST\"><textarea name=\"body\">%s</textarea><br/><input type=\"submit\" value=\"Save\"></form>", page.Title, page.Title, page.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	page.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
