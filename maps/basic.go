package maps

import "fmt"

type User struct {
	Id    string
	Email string
	Name  string
}

func MapsBasic() {
	users := map[string]User{
		"john":  {"aa1", "john@test.com", "John Doe"},
		"sarah": {"aa2", "sara@miller.ai", "Sarah Miller"},
	}

	for key, el := range users {
		fmt.Println("Key: ", key, ", User: ", el.Name, ", Email: ", el.Email)
	}
}
