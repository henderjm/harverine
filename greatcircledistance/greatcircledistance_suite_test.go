package greatcircledistance_test

import (
	"testing"

	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHarverine(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GreatCircleDistance Suite")
}

func buildFilePath(filename string) string {
	return filepath.FromSlash("./fixtures/" + filename)
}
