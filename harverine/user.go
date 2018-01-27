package harverine

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

func createUser() User {
	return User{}
}

func CreateUsers(data []string) ([]User, error) {
	var users []User
	for _, d := range data {
		u := User{}
		err := json.Unmarshal([]byte(d), &u)
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

func AmIInvited(fr *FileReader) ([]User, error) {
	//var users []Users
	//
	//err := json.Unmarshal(data, &users)
	//if err != nil {
	//	return nil, err
	//}

	return []User{}, nil
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
