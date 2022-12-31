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
}
