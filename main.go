package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValidPrefix(line string) bool {
	validPrefixes := []string{"feat: ", "fix: ", "chore: ", "doc: ", "refactor: "}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No commit message file provided")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: Unable to open the commit message file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		if !isValidPrefix(firstLine) {
			fmt.Println("Commit message does not start with a valid prefix (feat: , fix: , chore: , doc: , refactor: ).")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading commit message: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Commit message is valid.")
}


