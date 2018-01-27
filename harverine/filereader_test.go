package harverine_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	. "intercom/harverine"
)

var _ = Describe("FileReader", func() {
	fileReader := NewFileReader()

	Context("Opening", func() {

		It("should return a slice of file lines", func() {
			lines, err := fileReader.ReadLines(buildFilePath("valid_json_one_line.json"))
			Expect(err).ToNot(HaveOccurred())

			Expect(len(lines)).To(Equal(1))
			Expect(lines[0]).To(Equal(fmt.Sprintf("{\"%s\" : \"%s\"}", "im", "valid json")))
		})

		Context("when file is invalid", func() {
			It("should return error containing invalid json", func() {
				_, err := fileReader.ReadLines(buildFilePath("invalid_json.json"))
				Expect(err).To(MatchError("Not valid json `I'm not json`"))
			})
		})
	})
})
