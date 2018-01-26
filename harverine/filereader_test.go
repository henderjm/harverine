package harverine_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "intercom/harverine"
	"path/filepath"
)

var _ = Describe("FileReader", func() {
	var fileReader FileReader
	Context("Opening", func() {
		It("should not error opening json file", func() {
			lines, err := fileReader.ReadLines(filepath.FromSlash("./fixtures/" + "simple_one_line.json"))
			Expect(err).ToNot(HaveOccurred())

			Expect(len(lines)).To(Equal(1))
			Expect(lines[0]).To(Equal("I'm not even valid json"))
		})
	})
})
