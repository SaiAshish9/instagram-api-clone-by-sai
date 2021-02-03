package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/saiashish9/instagram/backend/utils/controllers/home"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://saiashish:saiashish@localhost/instagram?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

}

func main() {

	x := home.HomeController(db)

	http.HandleFunc("/", test)
	http.HandleFunc("/status-list", x.FetchStatusList)
	http.HandleFunc("/posts", x.FetchPosts)
	http.HandleFunc("/categories", x.FetchCategories)
	http.HandleFunc("/media", x.FetchMedia)
	http.HandleFunc("/suggestions", x.FetchSuggestions)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func test(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/status-list", http.StatusMovedPermanently)
}
