package pkg

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	neturl "net/url"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// ScrapeURL makes an HTTP GET request to the specified URL, parses the HTML content,
// finds all anchor tags, and returns a map of href values that have numbers 1-10 right before ".html".
func ScrapeURL(url string) (map[string]string, error) {
	scraperAPIURL, err := neturl.Parse("http://localhost:8080")
	if err != nil {
		return nil, err
	}

	query := scraperAPIURL.Query()
	query.Set("url", "https://"+url)
	scraperAPIURL.RawQuery = query.Encode()

	log.Printf("Searching URL: %s", scraperAPIURL.String())

	res, err := http.Get(scraperAPIURL.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}

	pattern := "^.+[0-9]\\.html$"
	hrefs := make(map[string]string)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					r, err := regexp.Compile(pattern)
					if err != nil {
						fmt.Println(err)
						return
					}
					if r.MatchString(a.Val) {
						refID := strings.Split(a.Val, "/")[len(strings.Split(a.Val, "/"))-1]
						refID = string(refID)
						id := refID[:len(refID)-5]
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

	log.Printf("Found %v links", len(hrefs))
	return hrefs, nil
}

// PostListings scrapes the listings for the specified search, checks if the assets/objects are new,
// inserts them into the database, and returns a slice of the URLs of the new listings.
func PostListings(search string, assetList map[string]bool) []string {
	var links []string
	listings, err := ScrapeURL(search)
	randomSleep()
	if err != nil {
		fmt.Print(err)
	}

	repo, err := NewRepository()
	if err != nil {
		fmt.Println(err)
		return links
	}

	for id, url := range listings {
		if IsNewAsset(id, assetList) {
			err := repo.InsertAsset(id, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			links = append(links, url)
		}
	}

	return links
}


func randomSleep() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(40)
	fmt.Printf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("Done")
}

// IsNewAsset checks if the asset/object with the specified ID exists in the database (assetList).
// It returns true if it's a new asset/object or false if it already exists.
func IsNewAsset(id string, assetList map[string]bool) bool {
	if assetList[id] {
		log.Printf("%s already exists in the database", id)
		return false
	}
	return true
}
