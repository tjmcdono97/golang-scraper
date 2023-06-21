package pkg

import (
	"database/sql"
	"fmt"
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

// FetchData opens a database connection, queries the database for all rows, and returns
// a slice of strings.
func FetchData() []string {
	db, err := sql.Open("sqlite3", sqliteDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data := make([]string, 0)

	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}
		data = append(data, value)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

// FetchIDs returns a map of IDs. The map is populated with the IDs of all
// the records in the database.
func FetchIDs() map[string]bool {
	db, err := sql.Open("sqlite3", sqliteDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ids := make(map[string]bool)

	rows, err := db.Query("SELECT id FROM vehicle_list")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		ids[id] = true
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return ids
}

// Insert inserts the ID and URL into the vehicle_list table.
func Insert(id, url string) {
	db, err := sql.Open("sqlite3", sqliteDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO vehicle_list (id, url, posting_time) VALUES (?, ?, ?)", id, url, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s placed into vehicle_list", url)
}
