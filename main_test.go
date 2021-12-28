package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MainTestSuite struct {
	suite.Suite
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
