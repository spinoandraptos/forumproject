package models

import (
	"github.com/spinoandraptos/forumproject/Server/database"
	"github.com/spinoandraptos/forumproject/Server/helper"
)

// a forum category has an unique ID, title and description
type Category struct {
	ID          uint32 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// to retrieve all created threads from the database
// we select every row from the threads table and sort by most recent thread on top
// then we store every row in a thread and append the thread to a slice of threads
// finally we and close the database connection
func RetrieveAllCategories() (categorySlice []Category, err error) {
	categories, err := database.DB.Query("SELECT * FROM categories ORDER BY CreatedAt DESC")
	if err != nil {
		helper.Catch(err)
	}
	for categories.Next() {
		var category Category
		err = categories.Scan(&category.ID, &category.Title, &category.Description)
		if err != nil {
			helper.Catch(err)
		}
		categorySlice = append(categorySlice, category)
	}
	categories.Close()
	return
}
