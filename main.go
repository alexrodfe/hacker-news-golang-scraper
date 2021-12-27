package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type scrapper struct {
	captureTitleRegex    *regexp.Regexp
	captureNComentsRegex *regexp.Regexp
	captureNPointsRegex  *regexp.Regexp
}

func (s scrapper) init() error {
	var (
		wrapErr = errorWrapper("init: %w")
		err     error
	)

	s.captureTitleRegex, err = regexp.Compile(`class="titlelink">(.*)</a><span`)
	if err != nil {
		return wrapErr(err)
	}
	s.captureNComentsRegex, err = regexp.Compile(`(\d+)&nbsp;comments`)
	if err != nil {
		return wrapErr(err)
	}
	s.captureNPointsRegex, err = regexp.Compile(`(\d+)\spoints`)
	if err != nil {
		return wrapErr(err)
	}

	return nil
}

func main() {
	// Make request
	response, err := http.Get("https://news.ycombinator.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	readText, _ := ioutil.ReadAll(response.Body)

	result1 := strings.Split(string(readText), `id="pagespace"`)
	result2 := strings.Split(result1[1], `</table>`)

	results := strings.Split(result2[0], "\n")

	for idx, result := range results {
		if idx == 0 {
			continue
		}
		fmt.Println(result)
	}

}

// errorWrapper will wrap error messages inside a desired format for proper error management and tracking
func errorWrapper(message string) func(params ...interface{}) error {
	return func(params ...interface{}) error {
		return fmt.Errorf(message, params...)
	}
}
