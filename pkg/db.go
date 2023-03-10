// The package name.
package pkg

// Importing the packages that are needed for the program to run.
import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/mattn/go-sqlite3"
)

const SQLITEDB = `D:\\GoCraigslist\\Craigslist.db`

// constant query string to select from searches table
const select_query string = `
	SELECT 
	  c.city || s.search_string 
	FROM
	searches s,
	cities c
`

// It opens a database connection, queries the database for all rows in the searches table, and returns
// a slice of Search structs.
func FetchSearches() []string {
	db, err := sql.Open("sqlite3", SQLITEDB)
	fmt.Print(sqlite3.SQLITE_TEXT)
	if err != nil {
		fmt.Print(err)
	}

	searches := []string{}

	defer db.Close()

	rows, err := db.Query(select_query)

	if err != nil {
		fmt.Print(err)
	}

	defer rows.Close()

	for rows.Next() {

		var search_string string

		err = rows.Scan(&search_string)

		searches = append(searches, search_string)

		if err != nil {
			fmt.Print(err)
		}
	}
	return searches
}

// FetchVehicleList() returns a map of strings to booleans. The map is populated with the id's of all
// the vehicles in the database
func FetchVehicleList() map[string]bool {
	db, err := sql.Open("sqlite3", SQLITEDB)
	ids := map[string]bool{}
	fmt.Print(sqlite3.SQLITE_TEXT)
	if err != nil {
		fmt.Print(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id FROM vehicle_list")

	if err != nil {
		fmt.Print(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id string

		err = rows.Scan(&id)

		ids[id] = true

		if err != nil {
			fmt.Print(err)
		}

	}
	return ids
}

// Inserts the id and url into the vehicle_list table
func Insert(id string, url string) {

	db, err := sql.Open("sqlite3", SQLITEDB)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer db.Close()
	_, err = db.Exec(`INSERT INTO vehicle_list(id,url,posting_time) VALUES(?,?,?)`, id, url, time.Now())
	if err != nil {
		fmt.Printf("%s", err)

	} else {
		log.Printf("%s placed into vehicle_list", url)
	}

}
