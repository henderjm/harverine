package greatcircledistance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "intercom/greatcircledistance"
	"sort"
	"strconv"
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

			Expect(len(users)).To(Equal(2))
		})

		It("all user fields should be populated", func() {
			users, err := CreateUsers(lines)
			Expect(err).ToNot(HaveOccurred())

			firstUser := users[0]
			Expect(firstUser.Name).To(Equal("James Hender"))
			Expect(firstUser.ID).To(Equal(24))
			Expect(firstUser.Latitude).To(Equal("52.986375"))
			Expect(firstUser.Longitude).To(Equal("-6.043701"))
		})

		Context("Failure cases", func() {
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

	Describe("Invitations", func() {
		var users []User
		var err error

		Context("When within 100 kilometers of Dublin Office", func() {
			JustBeforeEach(func() {
				users, err = CreateUsers(lines)
				Expect(err).ToNot(HaveOccurred())
			})

			It("You should be invited", func() {
				invitedUsers, err := InviteUsers(users, 100)
				Expect(err).ToNot(HaveOccurred())
				Expect(len(invitedUsers)).To(Equal(2))
			})

			It("Should return a sorted slice of invitees", func() {
				invitedUsers, err := InviteUsers(users, 100)
				Expect(err).ToNot(HaveOccurred())

				unsortedUsers := invitedUsers
				sort.Sort(ByID(invitedUsers))

				Expect(unsortedUsers[0].ID).To(Equal(1))
				Expect(unsortedUsers[1].ID).To(Equal(24))
			})
		})

		Context("When outside of 100 kilometers of Dublin Office", func() {
			JustBeforeEach(func() {
				lines, _ = fileReader.ReadLines(buildFilePath("not_invitable_users.json"))
				users, err = CreateUsers(lines)
				Expect(err).ToNot(HaveOccurred())
			})

			It("You sadly miss out on the invitation", func() {
				invitedUsers, err := InviteUsers(users, 10)
				Expect(err).ToNot(HaveOccurred())
				Expect(len(invitedUsers)).To(Equal(0))
			})
		})

		Context("Failure cases", func() {
			JustBeforeEach(func() {
				lines, _ = fileReader.ReadLines(buildFilePath("not_valid_latitude_user.json"))
				users, err = CreateUsers(lines)
				Expect(err).ToNot(HaveOccurred())
			})

			It("Should raise error when user latitude is invalid", func() {
				_, err := InviteUsers(users, 100)

				expectedErr := &strconv.NumError{
					Num:  "incorrectvalue",
					Func: "ParseFloat",
					Err:  strconv.ErrSyntax,
				}
				Expect(err).To(Equal(expectedErr))
			})
		})

	})
})
