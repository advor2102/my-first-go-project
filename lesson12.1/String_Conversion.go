package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	url := `C:\Users\AlexanderDvornikov\golang_course\my-first-go-project\lesson12\input.txt`
	data, err := os.ReadFile(url)
	if err != nil {
		fmt.Println("Reading file error:", err)
	}

	rawText := strings.ToLower(string(data))
	cleanText := strings.Join(strings.Fields(rawText), " ")
	words := strings.Fields(cleanText)
	wordsCount := make(map[string]int)
	for _, v := range words {
		cleanWord := strings.TrimFunc(v, func(r rune) bool {
			return unicode.IsPunct(r)
		})
		if cleanWord == "" {
			continue
		}
		wordsCount[cleanWord]++
	}
	file, err := os.Create(`C:\Users\AlexanderDvornikov\golang_course\my-first-go-project\lesson12\output.txt`)
	if err != nil {
		fmt.Println("Creating file error:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err = writer.Write([]string{"слово", "частота"}); err != nil {
		fmt.Println("Writing file error:", err)
	}

	for word, freq := range wordsCount {
		record := []string{word, fmt.Sprintf("%d", freq)}
		if err = writer.Write(record); err != nil {
			fmt.Println("Writing file error:", err)
		}
	}
}
