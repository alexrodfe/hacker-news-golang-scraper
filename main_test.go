package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
	scrapper
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}

func (suite *MainTestSuite) SetupSuite() {
	err := suite.scrapper.init()
	if err != nil {
		log.Fatalf("could not init scrapper: %v", err)
	}
}

func (suite *MainTestSuite) TestExample1() {
	f, err := os.Open("./test_examples/example1.html")
	require.NoError(suite.T(), err)
	data, err := ioutil.ReadAll(f)
	require.NoError(suite.T(), err)

	resultCollection, err := suite.scrapper.scrapWebpage(data)
	require.NoError(suite.T(), err)

	suite.Len(resultCollection, 30)
	suite.Equal(example1Result, resultCollection)
}

var example1Result = entryCollection{
	{title: "Ask HN: What is your spiritual practice?",
		nComments: 621, // number of comments
		nPoints:   289},
	{title: `"Widevine Dump":Leaked Code Downloads HD Video from Disney+, Amazon, and Netflix`,
		nComments: 214,
		nPoints:   307}, // over number of points
	{title: "A Review of the Zig Programming Language (Using Advent of Code 2021)",
		nComments: 147,
		nPoints:   211},
	{title: "Maybe we could tone down the JavaScript (2016)",
		nComments: 114,
		nPoints:   126},
	{title: "Italian Courts Find Open Source Software Terms Enforceable",
		nComments: 77,
		nPoints:   406},
	{title: "Ask HN: How did my LastPass master password get leaked?",
		nComments: 72,
		nPoints:   110},
	{title: "A stochastic method to generate the Sierpinski triangle",
		nComments: 57,
		nPoints:   96},
	{title: "The Swiss wanderer who found the soul of 1950s Japan",
		nComments: 44,
		nPoints:   142},
	{title: "rC3 Fahrplan – The Chaos Communication Congress 2021 schedule",
		nComments: 43,
		nPoints:   99},
	{title: "immudb – world’s fastest immutable database, built on a zero trust model",
		nComments: 37,
		nPoints:   73},
	{title: "Recording 660FPS Video on a $6 Raspberry Pi Camera (2019)",
		nComments: 32,
		nPoints:   155},
	{title: "New muscle layer discovered on the jaw",
		nComments: 13,
		nPoints:   50},
	{title: "The Best Things and Stuff of 2021",
		nComments: 10,
		nPoints:   129},
	{title: "Luerl – An Implementation of Lua in Erlang",
		nComments: 7,
		nPoints:   56},
	{title: "Show HN: Drovp – Convenient UI for any drag and drop operations",
		nComments: 3,
		nPoints:   9},
	{title: "How to Open a Door (1979) [video]",
		nComments: 0,
		nPoints:   4},
	{title: "Quick and dirty way to rip an eBook from Android",
		nComments: 0,
		nPoints:   27},
	{title: "Compose.ai (YC W21) Is Hiring Engineers and Designers",
		nComments: 0,
		nPoints:   0},
	{title: "Five takeaways from looking for a new senior role in tech",
		nComments: 0,
		nPoints:   6},
	{title: "Read J.D. Salinger’s first short story to feature Holden Caufield",
		nComments: 0,
		nPoints:   13},
	{title: "Windows 2000 Modernization Guide", // entries with 5 or less words are second in list
		nComments: 149,
		nPoints:   159},
	{title: "Plant Root System Drawings",
		nComments: 12,
		nPoints:   145}, // number of points
	{title: "Decoding James Webb Space Telescope",
		nComments: 41, // over number of comments
		nPoints:   127},
	{title: "Rack 2 (Virtual Eurorack)",
		nComments: 59,
		nPoints:   120},
	{title: "Practical Transformer Winding (2010)",
		nComments: 30,
		nPoints:   104},
	{title: "Metrics-driven product development is hard",
		nComments: 19,
		nPoints:   42},
	{title: "Speedcabling",
		nComments: 16,
		nPoints:   37},
	{title: "The Economist tracks excess deaths",
		nComments: 0,
		nPoints:   29},
	{title: "Regression with the C64",
		nComments: 20,
		nPoints:   27},
	{title: "Self-made EL segment displays",
		nComments: 2,
		nPoints:   26},
}
