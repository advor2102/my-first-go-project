package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Stat(`C:\Users\AlexanderDvornikov\golang_course\my-first-go-project\lesson12.2\log.txt`)
	if os.IsNotExist(err) {
		logFile, err := os.Create(`C:\Users\AlexanderDvornikov\golang_course\my-first-go-project\lesson12.2\log.txt`)
		if err != nil {
			log.Println("Creating file error: ", err)
			return
		}
		logFile.Close()
	} else if err != nil {
		fmt.Println(err)
		return
	}

	logFile, err := os.OpenFile(`C:\Users\AlexanderDvornikov\golang_course\my-first-go-project\lesson12.2\log.txt`, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Opening file error:", err)
		return
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags)

	scanner := bufio.NewScanner(os.Stdin)

	log.Println("Scanning started. Type 'exit' if you want to stop")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()

		if text == "exit" {
			log.Println("Logging process was stopped")
			break
		}
		log.Println(text)
	}

	if err = scanner.Err(); err != nil {
		log.Println("Scanning error", err)
	}
}
