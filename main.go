package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	errorutility "github.com/alexrodfe/hacker-news-golang-scraper/error_utility"
	"github.com/alexrodfe/hacker-news-golang-scraper/scraper"
)

func main() {
	wrapErr := errorutility.ErrorWrapper("error in main: %w")

	// Make request
	response, err := http.Get("https://news.ycombinator.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	webpage, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(wrapErr(err))
		os.Exit(1)
	}

	s := scraper.Scraper{}
	s.Init()

	resultCollection, err := s.ScrapWebpage(webpage)
	if err != nil {
		fmt.Println(wrapErr(err))
		os.Exit(1)
	}

	for _, result := range resultCollection {
		fmt.Println(result)
	}
}
