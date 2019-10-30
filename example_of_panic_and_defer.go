package main

import (
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ShoppingCart struct {
	gorm.Model
	Item string
}

var (
	tpl    *template.Template
	tplTwo *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("main.html"))
	tplTwo = template.Must(template.ParseGlob("searchsite.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/adding", processor)
	http.HandleFunc("/adding/searchsite", parseThenQuery)
	http.HandleFunc("/searchsite", parseThenQuery)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "main.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "shoppinglist.sqlite")
	// panics if there's an error open database opening 
	if err != nil {
		panic("failed to connect to db")
	}
	// close the db after the function returns
	defer db.Close()
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	item := r.FormValue("item")

	db.AutoMigrate(&ShoppingCart{})
	db.Create(&ShoppingCart{Item: item})
	db.Model(&ShoppingCart{Item: item})

	tpl.ExecuteTemplate(w, "main.html", item)
}

//commented-out code doesn't perform a desired task; leaving it for now
func parseThenQuery(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "availableItems.sqlite")
	// panics if there's an error open database opening 
	if err != nil {
		panic("failed to connect to db")
	}
	// close the db after the function returns
	defer db.Close()

	// searchedItem := r.FormValue("item")
	// db.Find(&item, searchedItem)

	tplTwo.ExecuteTemplate(w, "searchsite.html", nil)
}
