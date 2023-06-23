package pkg

import (
	"database/sql"
	"log"
	"os"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

var sqliteDB = os.Getenv("SQLITE_DB")

const selectQuery = `
	SELECT 
		c.city || s.search_string 
	FROM
		searches s,
		cities c
`

// Repository is a type that holds a pointer to a SQL database.
type Repository struct {
	DB *sql.DB
}

// NewRepository creates and returns a new instance of Repository,
// which holds a pointer to a new SQL database connection.
func NewRepository() (*Repository, error) {
	db, err := sql.Open("sqlite3", sqliteDB)
	if err != nil {
		return nil, err
	}
	
	return &Repository{DB: db}, nil
}

// FetchData retrieves data from a SQL database, and returns it as a slice of strings.
func (r *Repository) FetchData() ([]string, error) {
	data := make([]string, 0)

	rows, err := r.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, err
		}
		data = append(data, value)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// FetchIDs retrieves IDs from a SQL database, and returns them as a map of strings to booleans.
func (r *Repository) FetchIDs() (map[string]bool, error) {
	ids := make(map[string]bool)

	rows, err := r.DB.Query("SELECT id FROM vehicle_list")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids[id] = true
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

// Insert adds a new record to the vehicle_list table in the SQL database.
func (r *Repository) Insert(id, url string) error {
	_, err := r.DB.Exec("INSERT INTO vehicle_list (id, url, posting_time) VALUES (?, ?, ?)", id, url, time.Now())
	if err != nil {
		return err
	}

	log.Printf("%s placed into vehicle_list", url)
	return nil
}
