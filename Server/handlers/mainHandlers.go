package handlers

import (
	"html/template"
	"net/http"

	"github.com/spinoandraptos/forumproject/Server/helper"
	"github.com/spinoandraptos/forumproject/Server/models"
)

// the mainpage handler is called on entry to the starting page
// it is supposed to display all the categories available on the forum
// it parses needed HTML files and create a set of templates using ParseFiles
// after parsing, Must helper.Catches any potential errors by panicking
// we then retrieve all the categories in the database and implemenet them in the template "layout"
// this formatted html file is then written to the client and displayed

//The three files are HTML files with certain embedded commands, called actions
//Actions are annotations added to the HTML between {{ and }}

func MainPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./frontend/src/components/layoutcategories.html",
		"./frontend/src/components/navbar.html",
		"./frontend/src/components/content.html"}
	templates := template.Must(template.ParseFiles(files...))
	categories, err := models.RetrieveAllCategories()
	if err != nil {
		helper.Catch(err)
	}
	templates.ExecuteTemplate(w, "layoutcategories", &categories)
}

func ThreadsPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO!"))
}

func CommmentsPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO!"))
}
