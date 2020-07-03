package main

import (
	"encoding/json"
	"os"
)

func main() {
	users := getUsers()
	outputJson, err := json.MarshalIndent(&users, "", "  ")
	if err != nil {
		return
	}

	file, err := os.OpenFile("output.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return
	}
	defer file.Close()

	file.Write(outputJson)
	return
}

type User struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Hobbies []Hobby `json:"hobbies"`
}

type Hobby struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}

func getUsers() []User {
	var users []User
	var hobbies1, hobbies2 []Hobby

	hobbies1 = append(hobbies1, *newHobby("sport", "volleyball"))
	hobbies1 = append(hobbies1, *newHobby("amusement", "movie"))
	user1 := newUser("1", "Alice", 20, hobbies1)
	users = append(users, *user1)

	hobbies2 = append(hobbies2, *newHobby("sport", "soccer"))
	hobbies2 = append(hobbies2, *newHobby("music", "guitar"))
	user2 := newUser("2", "Bob", 25, hobbies2)
	users = append(users, *user2)

	return users
}

func newUser(id, name string, age int, hobbies []Hobby) *User {
	return &User{
		Id:      id,
		Name:    name,
		Age:     age,
		Hobbies: hobbies,
	}
}

func newHobby(category, name string) *Hobby {
	return &Hobby{
		Category: category,
		Name:     name,
	}
}
