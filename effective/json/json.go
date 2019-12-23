package main

import (
	"encoding/json"
	"fmt"
)

/* omitempty的用法*/
type Tips struct {
	Name string `json:"name"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Tips     *Tips  `json:"tips,omitempty"`
}

func main()  {

	tips := Tips{Name:"hello"}
	user := User{
		Email:    "liumingming",
		Password: "123456",
		Tips:     &tips,
	}

	jsonStr, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonStr))

	var user2 User
	err = json.Unmarshal(jsonStr, &user2)
	if err != nil {
		fmt.Println("unmarsh failed")
	}
	fmt.Println(user2)
}



