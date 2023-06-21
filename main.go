package main

import (
	"craigslist.com/scraper/pkg"
	"fmt"
	"log"
	"time"
)

// setupLogs creates a log file with a timestamp in the name and sets the log output to that file.
func setupLogs() {
	const layout = "2006_01_02_15_04_05"
	t := time.Now().Format(layout)
	logFileName := fmt.Sprintf("D:\\GoCraigslist\\logs\\%s.log", t)
	logFile, err := pkg.OpenLogFile(logFileName)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
}

func main() {
	setupLogs()

	searchList := pkg.FetchData()
	vehicleList := pkg.FetchIDs()
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
