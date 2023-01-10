package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/spinoandraptos/forumproject/Server/database"
	"github.com/spinoandraptos/forumproject/Server/helper"
	"github.com/spinoandraptos/forumproject/Server/models"
)

// to view a category, retrieve the category id from the URL sent in request
// and look up the row in category table with a matching category id
// then we scan the row's info into a category struct and send this struct back to client-side

func ViewCategory(w http.ResponseWriter, r *http.Request) {

	categoryid, err := strconv.Atoi(chi.URLParam(r, "categoryid"))
	helper.Catch(err)
	var category models.Category
	response := database.DB.QueryRow("SELECT * FROM categories WHERE ID = $1", categoryid)
	err = response.Scan(&category.ID, &category.Title, &category.Description)
	helper.Catch(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

// to view all categories, we just retrieve all categories from category table, arranging them by their id in ascending order
// then we scan the info for each catageory into a category struct within a slice of category structs
// finally we send the entire slice back to client-side and close query connection

func ViewCategories(w http.ResponseWriter, r *http.Request) {

	var allcategories []models.Category
	categories, err := database.DB.Query("SELECT * FROM categories ORDER BY ID ASC")
	helper.Catch(err)
	for categories.Next() {
		var category models.Category
		err = categories.Scan(&category.ID, &category.Title, &category.Description)
		helper.Catch(err)
		allcategories = append(allcategories, category)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allcategories)
	categories.Close()
}
