package analyzer

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestSampleFiles(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %s", err)
	}

	testData := filepath.Join(filepath.Dir(wd), "..", "testdata")
	analysistest.Run(t, testData, Analyzer, "sample")
}
