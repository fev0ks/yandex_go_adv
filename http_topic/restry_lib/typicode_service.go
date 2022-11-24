package main

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"sort"
	"time"
)

type (
	User struct {
		ID       int     `json:"id"`
		Name     string  `json:"name"`
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Address  Address `json:"address"`
		Phone    string  `json:"phone"`
		Website  string  `json:"website"`
		Company  Company `json:"company"`
	}
	Address struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
	}

	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	}
)

func (u User) String() string {
	return fmt.Sprintf("{ id: %d, name: %s, username: %s, email: %s, phone: %s, website: %s, address: %v, company: %v }",
		u.ID, u.Name, u.Username, u.Email, u.Phone, u.Website, u.Address, u.Company)
}

//Приложение должно уметь:
//Выводить всех пользователей /users/ в виде JSON: GET https://jsonplaceholder.typicode.com/users.
//Выводить пользователей в отсортированном виде по полю name, а не по id. Для этого сделайте сортировку внутри приложения.
//В случае ошибки от сервера повторять запрос 5 раз с интервалом 10 секунд. После пятого неуспешного запроса выводить в консоль сообщение о невозможности сделать запрос.

func main() {
	var users []User
	//url := "https://jsonplaceholder.typicode.com/users"
	//url := "https://jsonplaceholder.typicode.comm/users"
	client := resty.New()

	client.
		// устанавливаем количество повторений
		SetRetryCount(5).
		// длительность ожидания между попытками
		SetRetryWaitTime(2 * time.Second).
		// длительность максимального ожидания
		SetRetryMaxWaitTime(10 * time.Second).
		SetBaseURL("https://jsonplaceholder.typicode.com")

	users, err := getUsers(client)
	if err != nil {
		fmt.Printf("failed to get users: %v\n", err)
	}
	sortUsers(users)
	printUsers(users)
}

func getUsers(client *resty.Client) ([]User, error) {
	var users []User
	var responseErr string

	resp, err := client.R().
		SetError(&responseErr).
		SetResult(&users).
		Get("/users")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("can't get users. Status code <> 200")
	}
	return users, nil
}

func sortUsers(users []User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})
}

func printUsers(users []User) {
	for i, user := range users {
		fmt.Printf("%d: %v\n", i, user)
	}
}
