package home

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/saiashish9/instagram/backend/utils/models/home/status"
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

	rows, err := h.db.Query("SELECT * FROM status")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	links := make([]status.Link, 0)
	for rows.Next() {
		lk := status.Link{}
		err := rows.Scan(&lk.ID, &lk.Name, &lk.Image)
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