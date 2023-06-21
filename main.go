package main

import (
	"craigslist.com/scraper/pkg"
	"fmt"
	"log"
	"os"
	"time"
)

// setupLogs creates a log file with a timestamp in the name and sets the log output to that file.
func setupLogs() {
	const layout = "2006_01_02_15_04_05"
	t := time.Now().Format(layout)
	logFilePath := os.Getenv("LOG_FILE_PATH")
	logFileName := fmt.Sprintf("%s\\logs\\%s.log", logFilePath, t)
	logFile, err := pkg.OpenLogFile(logFileName)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
}

func main() {
	setupLogs()

	// Creating a new repository
	repo, err := pkg.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	// Fetching data
	searchList, err := repo.FetchData()
	if err != nil {
		log.Fatal(err)
	}

	// Fetching IDs
	vehicleList, err := repo.FetchIDs()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Checking %v links", len(searchList))

	var links []string

	for _, search := range searchList {
		log.Printf("Checking URL: %s", search)
		searchLinks := pkg.PostListings(search, vehicleList)
		links = append(links, searchLinks...)
	}

	message := ""
	for _, link := range links {
		if len(message)+len(link) >= 1300 && len(message)+len(link) <= 1600 {
			pkg.SendMessage(message)
			message = link + "\n"
		} else {
			message += link + "\n"
		}
	}

	if len(message) > 3 {
		pkg.SendMessage(message)
		pkg.Alert()
	}

	log.Println("Exiting...")
}
