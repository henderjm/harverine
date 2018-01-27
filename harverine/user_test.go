package harverine_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "intercom/harverine"
)

var _ = Describe("User", func() {

	fileReader := NewFileReader()
	var lines []string

	BeforeEach(func() {
		lines, _ = fileReader.ReadLines(buildFilePath("invitable_user.json"))
	})

	Context("when unmarshalling json input", func() {

		It("Should unmarshal a valid json user string", func() {
			users, err := CreateUsers(lines)
			Expect(err).ToNot(HaveOccurred())

			Expect(len(users)).To(Equal(1))
		})

		It("all user fields should be populated", func() {
			users, err := CreateUsers(lines)
			Expect(err).ToNot(HaveOccurred())

			firstUser := users[0]
			Expect(firstUser.Name).To(Equal("Mark Hender"))
			Expect(firstUser.ID).To(Equal(1))
			Expect(firstUser.Latitude).To(Equal("52.986375"))
			Expect(firstUser.Longitude).To(Equal("-6.043701"))
		})

		Context("When user properties are missing", func() {
			It("Should raise error if #Name is not found", func() {
				invalidLines, _ := fileReader.ReadLines(buildFilePath("invalid_user_missing_name.json"))
				_, err := CreateUsers(invalidLines)
				Expect(err).To(MatchError(ContainSubstring("Missing required field(s) `[Name]`")))
			})

			It("Should raise error if #ID is not found", func() {
				invalidLines, _ := fileReader.ReadLines(buildFilePath("invalid_user_missing_id.json"))
				_, err := CreateUsers(invalidLines)
				Expect(err).To(MatchError(ContainSubstring("Missing required field(s) `[ID]`")))
			})

			It("Should raise error if #Latitude is not found", func() {
				invalidLines, _ := fileReader.ReadLines(buildFilePath("invalid_user_missing_latitude.json"))
				_, err := CreateUsers(invalidLines)
				Expect(err).To(MatchError(ContainSubstring("Missing required field(s) `[Latitude]`")))
			})

			It("Should raise error if #Longitude is not found", func() {
				invalidLines, _ := fileReader.ReadLines(buildFilePath("invalid_user_missing_longitude.json"))
				_, err := CreateUsers(invalidLines)
				Expect(err).To(MatchError(ContainSubstring("Missing required field(s) `[Longitude]`")))
			})

			It("Should raise error if #Name and #ID are not found", func() {
				invalidLines, _ := fileReader.ReadLines(buildFilePath("invalid_user_missing_name_and_id.json"))
				_, err := CreateUsers(invalidLines)
				Expect(err).To(MatchError(ContainSubstring("Missing required field(s) `[Name, ID]`")))
			})
		})

	})

	Describe("Invited", func() {
		Context("When inside 100 kilometers of Dublin Office", func() {

		})

	})
})
