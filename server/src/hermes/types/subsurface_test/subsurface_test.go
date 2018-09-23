package subsurface_test

import (
	"gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

type SubsurfaceTest struct{}

var _ = check.Suite(&SubsurfaceTest{})
