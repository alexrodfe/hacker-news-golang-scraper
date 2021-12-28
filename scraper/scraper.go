package scraper

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	errorutility "github.com/alexrodfe/hacker-news-golang-scraper/error_utility"
)

// Entry will represent an article and their relevant information important to us
type Entry struct {
	Title     string
	NComments uint
	NPoints   uint
}

var EmptyEntry = Entry{}

// EntryCollection is a collection of entries,
// this type is declared so we can define some sorting needed functions
type EntryCollection []Entry

func (ec EntryCollection) Len() int { return len(ec) }

// Less will apply our sorting rules, which are the following:
// Titles with more than 5 words come first
// If both titles have more than 5 words, greater number of comments comes first
// If both titles are short, greater number of points comes first
func (ec EntryCollection) Less(i, j int) bool {
	entry1 := ec[i]
	entry2 := ec[j]
	title1InWords := strings.Split(entry1.Title, " ")
	title2InWords := strings.Split(entry2.Title, " ")
	isTitle1Short := len(title1InWords) <= 5
	isTitle2Short := len(title2InWords) <= 5

	if isTitle1Short && !isTitle2Short {
		return false
	} else if !isTitle1Short && isTitle2Short {
		return true
	} else if isTitle1Short && isTitle2Short { // both short
		return entry1.NPoints > entry2.NPoints
	} else { // both long
		return entry1.NComments > entry2.NComments
	}

}
func (ec EntryCollection) Swap(i, j int) { ec[i], ec[j] = ec[j], ec[i] }

type Scraper struct {
	captureTitleRegex    *regexp.Regexp
	captureNComentsRegex *regexp.Regexp
	captureNPointsRegex  *regexp.Regexp
}

func (s *Scraper) Init() error {
	var (
		wrapErr = errorutility.ErrorWrapper("scraper init: %w")
		err     error
	)

	// first ".*?" is because we may encounter a `rel="nofollow"` in between
	s.captureTitleRegex, err = regexp.Compile(`class="titlelink".*?>(.+?)<\/a>`)
	if err != nil {
		return wrapErr(err)
	}
	s.captureNComentsRegex, err = regexp.Compile(`(\d+)&nbsp;comment`)
	if err != nil {
		return wrapErr(err)
	}
	s.captureNPointsRegex, err = regexp.Compile(`(\d+)\spoints`)
	if err != nil {
		return wrapErr(err)
	}

	return nil
}

// ScrapWebpage will take as input a webpage as raw data and return an ordered entry collection
// For this we have overfited the function to Hacker News' HTML code, making precise cuts within their page structure
func (s Scraper) ScrapWebpage(webpage []byte) EntryCollection {
	result1 := strings.Split(string(webpage), `id="pagespace"`) // we want to start scraping from first entry

	result2 := strings.Split(result1[1], `</table>`) // our data end will be the end of the table

	results := strings.Split(result2[0], "\n") // now all entries will come in a 3 rows group

	resultCollection := make(EntryCollection, 0)
	for idx := 1; len(resultCollection) < 30 && idx < len(results)-2; { // we want 30 entries but what if there are less?
		newEntry := s.ExtractEntry(idx, results)
		if newEntry != EmptyEntry { // do not append empty entries
			resultCollection = append(resultCollection, newEntry)
		}
		idx = idx + 4
	}

	sort.Sort(resultCollection)
	return resultCollection
}

// ExtractEntry will take the next 3 rows from a given row collection and will try to extract data from said rows
// data will be extracted following declared scraper regex rules, if no data is extracted function will suppose value is 0
// except for title, title is mandatory, if no title is fetched function will stop and return empty Entry
func (s Scraper) ExtractEntry(idx int, entries []string) Entry {
	var (
		title       string
		nComments64 uint64
		nPoints64   uint64
	)
	firstRow := entries[idx+1]
	secondRow := entries[idx+2]

	titleResult := s.captureTitleRegex.FindStringSubmatch(firstRow)
	nCommentsResult := s.captureNComentsRegex.FindStringSubmatch(secondRow)
	nPointsResult := s.captureNPointsRegex.FindStringSubmatch(secondRow)

	if len(titleResult) > 0 {
		title = titleResult[1]
	} else {
		return EmptyEntry
	}
	if len(nCommentsResult) > 0 {
		nComments64, _ = strconv.ParseUint(nCommentsResult[1], 10, 32)
	}
	if len(nPointsResult) > 0 {
		nPoints64, _ = strconv.ParseUint(nPointsResult[1], 10, 32)
	}

	return Entry{
		Title:     title,
		NComments: uint(nComments64),
		NPoints:   uint(nPoints64),
	}
}
