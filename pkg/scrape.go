// This is the package declaration. It is the first line of code in a Go source file.
package pkg

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// The function takes a URL as an argument, makes an HTTP GET request to the URL, parses the HTML
// content, finds all the anchor tags, and returns a map of the href values that have numbers 1-10
// right before ".html"
func ScrapeCraigslist(url string) (map[string]string, error) {
	// Make an HTTP GET request to the URL
	scraperApiUrl := fmt.Sprintf("http://localhost:8080/?url=https://%s", url)
	log.Printf("searching url: %s", scraperApiUrl)
	res, err := http.Get(scraperApiUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Load the HTML document
	// Parse the HTML content
	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Find all the anchor tags
	pattern := "^.+[0-9]\\.html$"
	var hrefs = make(map[string]string)
	var f func(*html.Node)
	f = func(n *html.Node) {
		// Checking if the node is an element node and if the data is an anchor tag.
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				// fmt.Print("\n",a.Val,"\n")
				if a.Key == "href" {
					// Use a regular expression to extract the href values that have numbers 1-10 right before ".html"
					r, err := regexp.Compile(pattern)
					if err != nil {
						fmt.Println(err)
						return
					}
					// Checking if the href value matches the regular expression pattern. If it does, it extracts the
					// number from the href value and adds it to the map.
					if r.MatchString(a.Val) {
						ref_id := strings.Split(a.Val, "/")[len(strings.Split(a.Val, "/"))-1]
						ref_id = string(ref_id)
						id := ref_id[:len(ref_id)-5]
						hrefs[id] = a.Val
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {

			f(c)
		}
	}
	f(doc)

	// Print the href values
	log.Printf("found %v links", len(hrefs))
	return hrefs, err

}

func PostListings(search string, vehicleList map[string]bool) []string {

	var links []string
	listings, err := ScrapeCraigslist(search)
	randomSleep()
	if err != nil {
		fmt.Print(err)
	}
	// This is a for loop that is iterating through the listings and printing them out.
	for id, url := range listings {
		// This is checking to see if the vehicle is new.
		if IsNewVehicle(id, vehicleList) {
			Insert(id, url)
			links = append(links, url)
		}
	}
	return links
}

func randomSleep() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(40) // n will be between 0 and 10
	fmt.Printf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("Done")
}

// If the vehicle ID is in the database, return false, otherwise return true
func IsNewVehicle(id string, vehicleList map[string]bool) bool {
	if vehicleList[id] {
		log.Printf("%s already exists in database", id)
		return false
	} else {
		return true
	}
}
