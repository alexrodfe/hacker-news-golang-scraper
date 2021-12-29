package scraper

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
	Scraper
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

func (suite *MainTestSuite) SetupSuite() {
	err := suite.Scraper.Init()
	if err != nil {
		log.Fatalf("could not init scrapper: %v", err)
	}
}

func (suite *MainTestSuite) TestSortEntryCollection() {
	var orderedCollection = EntryCollection{
		{Title: "I've seen footage, I stay noided", NComments: 322, NPoints: 6},
		{Title: "My starter is a shiny? More likely than you'd think", NComments: 252, NPoints: 350},
		{Title: "Don't you even know dignity when you see it?", NComments: 120, NPoints: 550},
		{Title: "Ratio: A modern take on social media", NComments: 115, NPoints: 265},
		{Title: "You expected a reference? Too bad! It was me, a test!", NComments: 42, NPoints: 375},
		{Title: "Like tears in the rain", NComments: 42, NPoints: 935},
		{Title: "Speedrunning: a dream's nightmare", NComments: 858, NPoints: 640},
		{Title: "Rivers in the desert", NComments: 5, NPoints: 515},
		{Title: "Cognitive Behavioral Therapy", NComments: 404, NPoints: 321},
		{Title: "The bionic Brooks", NComments: 547, NPoints: 11},
	}

	copyOfCollection := make(EntryCollection, len(orderedCollection))
	copy(copyOfCollection, orderedCollection)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(copyOfCollection), func(i, j int) {
		copyOfCollection[i], copyOfCollection[j] = copyOfCollection[j], copyOfCollection[i]
	})
	suite.NotEqual(orderedCollection, copyOfCollection)

	sort.Sort(copyOfCollection)
	suite.Equal(orderedCollection, copyOfCollection)
}

func (suite *MainTestSuite) TestExample1() {
	f, err := os.Open("./test_examples/example1.html")
	require.NoError(suite.T(), err)
	data, err := ioutil.ReadAll(f)
	require.NoError(suite.T(), err)

	resultCollection := suite.Scraper.ScrapWebpage(data)

	suite.Len(resultCollection, 30)
	suite.Equal(example1Result, resultCollection)
}

var example1Result = EntryCollection{
	{Title: "Ask HN: What is your spiritual practice?",
		NComments: 621, // number of comments
		NPoints:   289},
	{Title: `"Widevine Dump":Leaked Code Downloads HD Video from Disney+, Amazon, and Netflix`,
		NComments: 214,
		NPoints:   307}, // over number of points
	{Title: "A Review of the Zig Programming Language (Using Advent of Code 2021)",
		NComments: 147,
		NPoints:   211},
	{Title: "Maybe we could tone down the JavaScript (2016)",
		NComments: 114,
		NPoints:   126},
	{Title: "Italian Courts Find Open Source Software Terms Enforceable",
		NComments: 77,
		NPoints:   406},
	{Title: "Ask HN: How did my LastPass master password get leaked?",
		NComments: 72,
		NPoints:   110},
	{Title: "A stochastic method to generate the Sierpinski triangle",
		NComments: 57,
		NPoints:   96},
	{Title: "The Swiss wanderer who found the soul of 1950s Japan",
		NComments: 44,
		NPoints:   142},
	{Title: "rC3 Fahrplan – The Chaos Communication Congress 2021 schedule",
		NComments: 43,
		NPoints:   99},
	{Title: "immudb – world’s fastest immutable database, built on a zero trust model",
		NComments: 37,
		NPoints:   73},
	{Title: "Recording 660FPS Video on a $6 Raspberry Pi Camera (2019)",
		NComments: 32,
		NPoints:   155},
	{Title: "New muscle layer discovered on the jaw",
		NComments: 13,
		NPoints:   50},
	{Title: "The Best Things and Stuff of 2021",
		NComments: 10,
		NPoints:   129},
	{Title: "Luerl – An Implementation of Lua in Erlang",
		NComments: 7,
		NPoints:   56},
	{Title: "Show HN: Drovp – Convenient UI for any drag and drop operations",
		NComments: 3,
		NPoints:   9},
	{Title: "How to Open a Door (1979) [video]",
		NComments: 0,
		NPoints:   4},
	{Title: "Quick and dirty way to rip an eBook from Android",
		NComments: 0,
		NPoints:   27},
	{Title: "Compose.ai (YC W21) Is Hiring Engineers and Designers",
		NComments: 0,
		NPoints:   0},
	{Title: "Five takeaways from looking for a new senior role in tech",
		NComments: 0,
		NPoints:   6},
	{Title: "Read J.D. Salinger’s first short story to feature Holden Caufield",
		NComments: 0,
		NPoints:   13},
	{Title: "Windows 2000 Modernization Guide", // entries with 5 or less words are second in list
		NComments: 149,
		NPoints:   159},
	{Title: "Plant Root System Drawings",
		NComments: 12,
		NPoints:   145}, // number of points
	{Title: "Decoding James Webb Space Telescope",
		NComments: 41, // over number of comments
		NPoints:   127},
	{Title: "Rack 2 (Virtual Eurorack)",
		NComments: 59,
		NPoints:   120},
	{Title: "Practical Transformer Winding (2010)",
		NComments: 30,
		NPoints:   104},
	{Title: "Metrics-driven product development is hard",
		NComments: 19,
		NPoints:   42},
	{Title: "Speedcabling",
		NComments: 16,
		NPoints:   37},
	{Title: "The Economist tracks excess deaths",
		NComments: 0,
		NPoints:   29},
	{Title: "Regression with the C64",
		NComments: 20,
		NPoints:   27},
	{Title: "Self-made EL segment displays",
		NComments: 2,
		NPoints:   26},
}
