package greatcircledistance_test

import (
	"intercom/greatcircledistance"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Harverine Invitations", func() {

	fileReader := greatcircledistance.NewFileReader()
	harversine := greatcircledistance.Harverine{}
	var lines []string
	var users []greatcircledistance.User
	var err error

	const (
		OFFICE_LONGITUDE = "-6.257664"
		OFFICE_LATITUDE  = "53.339428"
	)

	BeforeEach(func() {
		lines, _ = fileReader.ReadLines(buildFilePath("invitable_user.json"))
		users, err = greatcircledistance.CreateUsers(lines)
		Expect(err).ToNot(HaveOccurred())
	})

	Context("When converting degrees to radian", func() {
		It("Should successfully apply the math", func() {
			radian := greatcircledistance.ConvertToRadians(53.339428)
			Expect(radian).To(Equal(0.9309486397304539))
		})
	})

	Context("Calculating distance from the office", func() {

		It("should return 44.XXXX kilometers for coordinates (52.986375, -6.043701)", func() {
			km, err := harversine.CalculateGreatCircleDistance(OFFICE_LATITUDE, "52.986375", OFFICE_LONGITUDE, "-6.043701")
			Expect(err).ToNot(HaveOccurred())
			Expect(km).To(Equal(float64(41.76872550083617)))
		})

		Context("Failure cases", func() {
			It("Should raise error on invalid #Latitude", func() {
				_, err := harversine.CalculateGreatCircleDistance(OFFICE_LATITUDE, "52.986375", OFFICE_LONGITUDE, "abc")
				Expect(err).To(MatchError(ContainSubstring("invalid syntax")))
				_, err = harversine.CalculateGreatCircleDistance("abc", "52.986375", OFFICE_LONGITUDE, "-6.043701")
				Expect(err).To(MatchError(ContainSubstring("invalid syntax")))
			})

			It("Should raise error on invalid #Longitude", func() {
				_, err := harversine.CalculateGreatCircleDistance(OFFICE_LATITUDE, "abc", OFFICE_LONGITUDE, "-6.043701")
				Expect(err).To(MatchError(ContainSubstring("invalid syntax")))
				_, err = harversine.CalculateGreatCircleDistance(OFFICE_LATITUDE, "52.986375", "abc", "-6.043701")
				Expect(err).To(MatchError(ContainSubstring("invalid syntax")))
			})

		})

	})

})
