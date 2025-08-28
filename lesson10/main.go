package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Book struct {
	Title string `xml:"title"`
	Year  int    `xml:"year"`
}
type Library struct {
	Books []Book `xml:"book"`
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

func loadConfigXML(path string) (Library, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Library{}, err
	}
	var b Library
	err = xml.Unmarshal(data, &b)
	return b, err
}

func main() {
	fileJSON := `C:\Users\AlexanderDvornikov\golang_course\my-first-go-project\lesson10\users.json`
	data, err := loadConfigJSON(fileJSON)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("First version JSON: %+v\n", data)
	for i, _ := range data {
		data[i].Age += 1
	}
	fmt.Printf("Version after age increasing JSON: %+v\n", data)
	raw, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(raw))

	fileXML := `C:\Users\AlexanderDvornikov\golang_course\my-first-go-project\lesson10\books.xml`
	library, err := loadConfigXML(fileXML)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("First version XML: %+v\n", library)
	for i := range library.Books {
		library.Books[i].Year++
	}
	fmt.Printf("Version after age increasing XML: %+v\n", library)
	rawXML, err := xml.MarshalIndent(library, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rawXML))
}
