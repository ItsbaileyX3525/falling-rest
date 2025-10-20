package api

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
}

type Fail struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Season() []byte {
	testJson := Test{FirstName: "Bailey", LastName: "Miles"}

	jsonStr, err := json.Marshal(testJson)

	if err != nil {
		fmt.Println("Error occured!")

		jsonErr := Fail{Success: false, Message: "Something went wrong"}
		jsonStr, _ := json.Marshal(jsonErr)
		return jsonStr
	}

	return jsonStr
}
