package util

import (
	"testing"
	"github.com/deepzz0/go-com/log"
)

func TestVersionCompare(t *testing.T) {
	log.Print(VersionCompareBigger("4.1.6566","4.1.3"))
}

