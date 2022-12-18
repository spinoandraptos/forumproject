package handlers

import (
	"html/template"
	"net/http"

	"github.com/spinoandraptos/forumproject/Server/models"
)

// the mainpage handler is called on entry to the starting page
// it is supposed to display all the categories available on the forum
// it parses needed HTML files and create a set of templates using ParseFiles
// after parsing, Must catches any potential errors by panicking
// we then retrieve all the categories in the database and implemenet them in the template "layout"
// this formatted html file is then written to the client and displayed

func MainPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./frontend/src/components/layout.html",
		"./frontend/src/components/navbar.html",
		"./frontend/src/components/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	categories, err := models.RetrieveAllCategories()
	if err != nil {
		catch(err)
	}
	templates.ExecuteTemplate(w, "layout", categories)
}

func ThreadsPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO!"))
}

func CommmentsPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO!"))
}
