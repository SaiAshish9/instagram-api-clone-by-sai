package home

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/saiashish9/instagram/backend/utils/models/home/categories"
	"github.com/saiashish9/instagram/backend/utils/models/home/posts"
	"github.com/saiashish9/instagram/backend/utils/models/home/status"
	"github.com/saiashish9/instagram/backend/utils/models/home/suggestions"
	"github.com/saiashish9/instagram/backend/utils/models/search"
)

type Home struct {
	db *sql.DB
}

func HomeController(db *sql.DB) *Home {
	return &Home{
		db: db,
	}
}

func (h Home) FetchStatusList(w http.ResponseWriter, r *http.Request) {

	// x := []status.Link{
	// 	{"Sai", "https://images.unsplash.com/photo-1553451310-1416336a3cca?ixid=MXwxMjA3fDB8MHxzZWFyY2h8OXx8dm9sbGV5fGVufDB8fDB8&ixlib=rb-1.2.1&auto=format&fit=crop&w=500&q=60"},
	// }
	// _ = x

	rows, err := h.db.Query("SELECT * FROM status ORDER BY id")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	links := make([]status.Link, 0)
	for rows.Next() {
		lk := status.Link{}
		err := rows.Scan(&lk.ID, &lk.Name, &lk.Image, &lk.Msg, &lk.StatusImg, &lk.StatusVideo, &lk.StatusImg, &lk.Time)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		links = append(links, lk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	m, _ := json.Marshal(links)
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}

func (h Home) FetchPosts(w http.ResponseWriter, r *http.Request) {

	rows, err := h.db.Query("SELECT * FROM posts")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	links := make([]posts.Post, 0)
	for rows.Next() {
		lk := posts.Post{}
		err := rows.Scan(&lk.ID, &lk.Name, &lk.URL, &lk.ProfileURL, &lk.Title, &lk.Description, &lk.CommentsCount, &lk.Time)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		links = append(links, lk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	m, _ := json.Marshal(links)
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)

}

func (h Home) FetchCategories(w http.ResponseWriter, r *http.Request) {

	rows, err := h.db.Query("SELECT * FROM categories")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	links := make([]categories.Category, 0)
	for rows.Next() {
		lk := categories.Category{}
		err := rows.Scan(&lk.ID, &lk.Title)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		links = append(links, lk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	m, _ := json.Marshal(links)
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}

func (h Home) FetchMedia(w http.ResponseWriter, r *http.Request) {

	rows, err := h.db.Query("SELECT * FROM media order by id")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	links := make([]search.Media, 0)
	for rows.Next() {
		lk := search.Media{}
		err := rows.Scan(&lk.ID, &lk.URL, &lk.IS_VIDEO, &lk.IS_GALLERY)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		links = append(links, lk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	m, _ := json.Marshal(links)
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}

func (h Home) FetchSuggestions(w http.ResponseWriter, r *http.Request) {

	rows, err := h.db.Query("SELECT * FROM suggestions order by id")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	links := make([]suggestions.Suggestion, 0)
	for rows.Next() {
		lk := suggestions.Suggestion{}
		err := rows.Scan(&lk.ID, &lk.Image, &lk.Title, &lk.Desc1, &lk.Desc2)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		links = append(links, lk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	m, _ := json.Marshal(links)
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)
}
