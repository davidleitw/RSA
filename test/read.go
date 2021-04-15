package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("README.md")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			fmt.Println("0..0")
		} else {
			fmt.Println(scanner.Text(), "\\n")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
