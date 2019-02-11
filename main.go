package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Firefox")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Copy data from resp to stdout
	body, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Num:", body)
}
