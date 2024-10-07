package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Filename argument required")
		return
	}
	fileName := os.Args[1]
	file, err := os.OpenFile(fileName, fds, 0777)
	if err != nil {
		fmt.Println("Error on loading file")
		fmt.Println(err)
	}
	defer file.Close()
	
	fmt.Println(string(data))
}