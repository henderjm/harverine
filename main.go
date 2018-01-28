package main

import (
	"fmt"
	"intercom/greatcircledistance"
	"log"
	"sort"
)

/*
We have some customer records in a text file (customers.json) -- one customer per line, JSON-encoded. We want to invite any customer within 100km of our Dublin office for some food and drinks on us. Write a program that will read the full list of customers and output the names and user ids of matching customers (within 100km), sorted by User ID (ascending).

You can use the first formula from this Wikipedia article to calculate distance. Don't forget, you'll need to convert degrees to radians.
The GPS coordinates for our Dublin office are 53.339428, -6.257664.
You can find the Customer list here.

⭑ Please don’t forget, your code should be production ready, clean and tested!
*/
const MAXDISTANCE = 100

func main() {
	fr := greatcircledistance.NewFileReader()

	data, err := fr.ReadLines("./greatcircledistance/fixtures/intercom_official_user_db.json")
	if err != nil {
		log.Fatal("Could not read file: %v", err)
	}

	users, err := greatcircledistance.CreateUsers(data)
	if err != nil {
		log.Fatal("Could not create users: %v", err)
	}

	invitedUsers, err := greatcircledistance.InviteUsers(users, MAXDISTANCE)
	if err != nil {
		log.Fatal("Could not create users: %v", err)
	}

	sort.Sort(greatcircledistance.ByID(invitedUsers))
	for _, user := range invitedUsers {
		fmt.Println(fmt.Sprintf("ID: %d, Name: %s", user.ID, user.Name))
	}
}
