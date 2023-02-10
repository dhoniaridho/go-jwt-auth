package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service struct{}

func (Service) GetAll() []User {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		print(err)
	}

	if err == nil {
		fmt.Print(resp.Body)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var users []User
	err = json.Unmarshal(body, &users)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	return users
}

func (Service) GetOne(id string) User {

	user := User{Name: "dhoniaridho", ID: 1}

	return user
}
