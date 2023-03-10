package main

import (
	"craigslist.com/scraper/pkg"
	"fmt"
	"log"
	"time"
)

// It creates a log file with a timestamp in the name, and then sets the log output to that file
func setupLogs() {
	const layout = "2006_01_02_15_04_05"
	t := time.Now().Format(layout)
	logFileName := fmt.Sprintf("D:\\GoCraigslist\\logs\\%s.log", t)
	logFile, err := pkg.OpenLogFile(logFileName)
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	if err != nil {
		log.Print(err)
	}
}

// This function is calling the FetchSearches and FetchVehicleList functions from the pkg package
func main() {
	// This is calling the FetchSearches and FetchVehicleList functions from the pkg package.
	var links []string
	setupLogs()
	searchList := pkg.FetchSearches()
	vehicleList := pkg.FetchVehicleList()
	log.Printf("checking %v links", len(searchList))

	// This is a function that is checking to see if the vehicle is new.
	for _, search := range searchList {
		log.Printf("Checking URL: %s", search)
		searchLinks := pkg.PostListings(search, vehicleList)
		for _, link := range searchLinks {
			links = append(links, link)
		}
	}
// This is checking to see if the message is greater than 3 characters. If it is, it will send the
// // message and alert the user.
	message := " "
	for _, link := range links {
		if len(message)+len(link) >= 1300 && len(message)+len(link) <= 1600 {
			pkg.SendMessage(message)
			message = link + "\n"
		} else {
			message += link + "\n"
		}
	}
// This is checking to see if the message is greater than 3 characters. If it is, it will send the
// message and alert the user.
	if len(message) > 3 {
		pkg.SendMessage(message)
		pkg.Alert()
	}
	
	log.Printf("Exiting...")
}