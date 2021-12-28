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
		wrapErr = errorutility.ErrorWrapper("init: %w")
		err     error
	)

	// may encounter a `rel="nofollow"` in between
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

func (s Scraper) ScrapWebpage(webpage []byte) (EntryCollection, error) {
	result1 := strings.Split(string(webpage), `id="pagespace"`)

	result2 := strings.Split(result1[1], `</table>`)

	results := strings.Split(result2[0], "\n")

	resultCollection := make(EntryCollection, 0)
	for idx := 1; len(resultCollection) < 30 && idx < len(results)-2; {
		newEntry := s.ExtractEntry(idx, results)
		if newEntry != EmptyEntry { // do not append empty entries
			resultCollection = append(resultCollection, newEntry)
		}
		idx = idx + 4
	}

	sort.Sort(resultCollection)
	return resultCollection, nil
}

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
	} // return err
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
