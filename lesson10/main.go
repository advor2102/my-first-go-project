package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func loadConfigJSON(path string) ([]User, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg []User
	if err = json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func main() {
	fileJSON := `C:\Users\AlexanderDvornikov\golang_course\my-first-go-project\lesson10\users.json`
	data, err := loadConfigJSON(fileJSON)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("First version: %+v\n", data)
	for i, _ := range data {
		data[i].Age += 1
	}
	fmt.Printf("Version after age increasing: %+v\n", data)
}
