package macdive_test

import (
	"gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type MacDiveTest struct {
}

var _ = check.Suite(&MacDiveTest{})
