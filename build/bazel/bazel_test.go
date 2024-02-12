package bazel_test

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/build/bazel"
)

func TestBuildWithBazel(t *testing.T) {
	if !bazel.BuiltWithBazel() {
		t.Error("not built with Bazel")
	}
}
