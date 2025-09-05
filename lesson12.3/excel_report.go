package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	file, err := os.ReadFile(`C:\Users\AlexanderDvornikov\golang_course\my-first-go-project\lesson12.3\users.json`)
	if err != nil {
		fmt.Println("Reading file error:", err)
		return
	}

	var users []User
	err = json.Unmarshal(file, &users)
	if err != nil {
		fmt.Println("Parsing json error:", err)
	}

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetCellValue(sheet, "A1", "Name")
	f.SetCellValue(sheet, "B1", "Age")

	for i, user := range users {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), user.Name)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), user.Age)
	}

	if err = f.SaveAs(`C:\Users\AlexanderDvornikov\golang_course\my-first-go-project\lesson12.3\report.xlsx`); err != nil {
		fmt.Println("Saving file error:", err)
		return
	}

	fmt.Println("File saved succesful")
}
