package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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

	doc.Find("tr").Each(processTableRow)
}

func processTableRow(index int, tableRow *goquery.Selection) {
	txt := tableRow.Text()
	announce := announcementMessage(txt)

	date, err := currentLocalizedDate()
	if err != nil {
		fmt.Printf("Fatal error: %s", err)
		os.Exit(1)
	}

	isToday := strings.Contains(txt, date)
	if isToday {
		fmt.Printf("%s\n\n", announce)
	}
}

// print text after 分, stripping newlines and whitespace
func announcementMessage(txt string) string {
	messageStartIdx := strings.Index(txt, "分") + 4
	message := string([]byte(txt[messageStartIdx:]))

	return strings.TrimSpace(message)
}

// Check date
func currentLocalizedDate() (string, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return "", err
	}
	now := time.Now().In(loc)

	localizedDate := fmt.Sprintf("%d月%d日", now.Month(), now.Day()-3) //TODO: minus 3 for debug purposes
	return localizedDate, nil
}
