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

// entry will represent an article and their relevant information important to us
type entry struct {
	title     string
	nComments uint
	nPoints   uint
}

// entryCollection is a collection of entries,
// this type is declared so we can define some sorting needed functions
type entryCollection []entry

func (ec entryCollection) Len() int { return len(ec) }
func (ec entryCollection) Less(i, j int) bool {
	entry1 := ec[i]
	entry2 := ec[j]
	isTitle1Short := len(entry1.title) <= 5
	isTitle2Short := len(entry2.title) <= 5

	if isTitle1Short && !isTitle2Short {
		return true
	} else if !isTitle1Short && isTitle2Short {
		return false
	} else if isTitle1Short && isTitle2Short { // both short
		return entry1.nComments < entry2.nComments
	} else { // both long
		return entry1.nPoints < entry2.nPoints
	}

}
func (ec entryCollection) Swap(i, j int) { ec[i], ec[j] = ec[j], ec[i] }

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
