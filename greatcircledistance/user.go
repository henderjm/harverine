package greatcircledistance

import (
	"encoding/json"
	"fmt"

	"strings"

	"github.com/pkg/errors"
)

type User struct {
	ID        int    `json:"user_id, required"`
	Name      string `json:"name, required"`
	Latitude  string `json:"latitude, required"`
	Longitude string `json:"longitude, required"`
}

type ByID []User

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func CreateUsers(data []string) ([]User, error) {
	var users []User
	for _, d := range data {
		u := User{}
		err := json.Unmarshal([]byte(d), &u) // This should never fail, caught in filereader if invalid.
		if err != nil {
			return nil, err
		}

		if props, ok := hasRequiredProperties(u); !ok {
			err := errors.New(fmt.Sprintf("Missing required field(s) `[%s]` in User: %#v", props, u))
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func InviteUsers(users []User, maxDistance float64) ([]User, error) {
	harversine := Harverine{}
	var invitedUsers []User
	for _, user := range users {
		distanceFromOffice, err := harversine.CalculateGreatCircleDistance(OFFICE_LATITUDE, user.Latitude, OFFICE_LONGITUDE, user.Longitude)
		if err != nil {
			return nil, err
		}

		if distanceFromOffice <= maxDistance {
			invitedUsers = append(invitedUsers, user)
		}
	}
	return invitedUsers, nil
}

func hasRequiredProperties(user User) (string, bool) {
	var missingProps []string
	isValid := true
	if user.Name == "" {
		missingProps = append(missingProps, "Name")
		isValid = false
	}
	if user.ID == 0 {
		missingProps = append(missingProps, "ID")
		isValid = false
	}
	if user.Latitude == "" {
		missingProps = append(missingProps, "Latitude")
		isValid = false
	}
	if user.Longitude == "" {
		missingProps = append(missingProps, "Longitude")
		isValid = false
	}

	return strings.Join(missingProps[:], ", "), isValid
}
