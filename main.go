package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yhkaplan/scraper/scraper"
)

func main() {
	lambda.Start(HandleReq) //TODO: move this to cmd dir
}

//TODO: move all this to cmd dir
func HandleReq() error {
	resp, err := http.Get("https://www.tokyometro.jp/unkou/history/hanzoumon.html")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create goquery doc
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	// Wrap this in errorable way?
	doc.Find("tr").Each(scraper.ProcessTableRow)

	return nil
}
