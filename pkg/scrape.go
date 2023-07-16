package pkg

import (
	"golang.org/x/net/html"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"go.uber.org/zap"
)

func ScrapeURL(urlStr string, logger *zap.Logger) (map[string]string, error) {
	scraperAPIURL, err := url.Parse("http://localhost:8000/")
	if err != nil {
		return nil, err
	}

	query := scraperAPIURL.Query()
	query.Set("url", "https://"+urlStr)
	scraperAPIURL.RawQuery = query.Encode()

	logger.Info("Searching urlStr", zap.String("URL", scraperAPIURL.String()))

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
						logger.Error("Error compiling regex", zap.Error(err))
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

	logger.Info("Found links", zap.Int("Count", len(hrefs)))
	return hrefs, nil
}

func PostListings(search string, assetList map[string]bool, repo *Repository, logger *zap.Logger) []string {
	var links []string
	listings, err := ScrapeURL(search, logger)
	randomSleep(logger)
	if err != nil {
		logger.Error("Error scraping URL", zap.Error(err))
	}

	for id, urlStr := range listings {
		if IsNewAsset(id, assetList, repo, logger) {
			repo.Insert(id, urlStr)
			links = append(links, urlStr)
		}
	}

	return links
}

func randomSleep(logger *zap.Logger) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(40)
	logger.Info("Sleeping", zap.Int("Seconds", n))
	time.Sleep(time.Duration(n) * time.Second)
	logger.Info("Done sleeping")
}

func IsNewAsset(id string, assetList map[string]bool, repo *Repository, logger *zap.Logger) bool {
	if assetList[id] {
		logger.Info("Asset already exists in the database", zap.String("ID", id))
		return false
	}
	return true
}
