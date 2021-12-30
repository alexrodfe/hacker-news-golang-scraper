# Hacker News Scraper (Golang)

A golang web scraper designed for the Hacker News webpage. This scraper will take the first 30 entries from the [Hacker News frontpage](https://news.ycombinator.com/) and then it will sort them in the following manner:

- First come entries with more than 5 words in their title
- Long title entries will then be sorted by number of comments
- Short title entries will be sorted by number of points

Then sorted entries will be printed out by console output.

## Installation

You just need [Go](https://go.dev/) !
Alongside the code in this repository:

```console
git clone https://github.com/alexrodfe/hacker-news-golang-scraper.git
```

This scraper has been designed to work with golang's native libraries, so no need for dependencies!

The repo also includes some unitary testing, if you wish to run them as well on your own then you will need to install [testify](https://github.com/stretchr/testify).
For this you can use the `go.mod` file attached or run:

```console
go get github.com/stretchr/testify
```

## Running

To see the scraper in action you can do so by running the following command inside the proyect's root folder:

```console
go run main.go
```

To run the tests you may use the conveniently placed makefile:

```console
make test
```
