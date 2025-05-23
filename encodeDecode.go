package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}

func main() {
	useall := UseAll{Name: "Vasiliy", Surname: "Konstantinov", Year: 2021}

	t, err := json.Marshal(&useall)
	fmt.Println("t->",t)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t)
	}

	str := `{"username": "V.", "surname": "Ko", "created":2020}`
	jsonRecord := []byte(str)
	temp := UseAll{}
	err = json.Unmarshal(jsonRecord, &temp)
	fmt.Println("jsonRecord->", jsonRecord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type: %T with value %v\n", temp, temp)
	}
}
