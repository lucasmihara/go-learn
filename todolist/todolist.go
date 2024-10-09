package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	fileName := getFileName()
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error on loading file")
		fmt.Println(err)
		return
	}
	fileContent := string(fileData)

	entries := strings.Split(fileContent, "\n")

	for index, entry := range entries {
		fmt.Printf("%d - %s\n", index + 1, entry)
	}

	fmt.Println("Choose a operation:")
	fmt.Println("1 - [C]reate a new entry")
	fmt.Println("2 - [R]emove a entry")

	var option string
	fmt.Scan(&option)

	if option == "1" || option == "C" || option == "c" {
		entry := readLine()
		entries = append(entries, entry)
	} else if option == "2" || option == "R" || option == "r" {

	} else {
		
	}
}

func getFileName() string {
	var fileName string
	if len(os.Args) > 1 {
		fileName = os.Args[1]
		return fileName
	}

	fmt.Println("Enter the file name")
	fmt.Scan(&fileName)
	
	return fileName
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\r')
	expression := line[:len(line) - 1]

	return expression
}
