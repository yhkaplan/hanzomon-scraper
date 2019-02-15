package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/yhkaplan/scraper/scraper"
)

func main() {
	resp, err := http.Get("https://www.tokyometro.jp/unkou/history/hanzoumon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create goquery doc
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tr").Each(scraper.ProcessTableRow)
}
