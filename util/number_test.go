package util

import (
	"log"
	"testing"
)

func TestNumber(t *testing.T) {

	log.Print(EqualFloatString("10.01","10.0",0))
	log.Print(EqualFloatString("10.01","10.011",2))
}

