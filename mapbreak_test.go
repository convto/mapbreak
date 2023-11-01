package mapbreak_test

import (
	"testing"

	"github.com/convto/mapbreak"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	patterns := []string{
		"singlefile",
		"nodetection",
		"pointer",
		"subpkg",
	}
	for _, pattern := range patterns {
		pattern := pattern
		t.Run(pattern, func(t *testing.T) {
			t.Parallel()
			analysistest.Run(t, testdata, mapbreak.Analyzer, pattern)
		})
	}
}
