package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	fileName := getFileName()
	tasks, err := getTasksFromFile(fileName)
	if err != nil {
		fmt.Println("Error on loading file")
		fmt.Println(err)
		return
	}
	for ;; {
		printTasks(tasks)
		option := getOption()
		
		if option == "1" || option == "C" || option == "c" {
			addTask(&tasks)
		} else if option == "2" || option == "R" || option == "r" {
			removeTask(&tasks)
		} else if option == "3" || option == "E" || option == "e" {
			saveTasks(tasks, fileName)
			return
		} else {
			fmt.Println("Invalid option")
		}
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

func getTasksFromFile(fileName string) ([]string, error) {
	fileContent, err := getFileContent(fileName)
	if err != nil {
		return nil, err
	}
	var tasks []string
	if len(fileContent) > 0 {
		tasks = strings.Split(fileContent, "\n")
	}
	return tasks, nil
}

func getFileContent(fileName string) (string, error) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(fileData), nil
}

func printTasks(tasks []string) {
	if len(tasks) > 0 {	
		fmt.Println("Tasks:")
		for index, task := range tasks {
			fmt.Printf("%d - %s\n", index + 1, task)
		}
	} else {
		fmt.Println("No tasks available")
	}
}

func getOption() string {
	fmt.Println("\nChoose an option:")
	fmt.Println("1 - [C]reate a new task")
	fmt.Println("2 - [R]emove a task")
	fmt.Println("3 - [E]xit and save")

	var option string
	fmt.Scanln(&option)

	return option
}

func addTask(tasks *[]string) {
	fmt.Println("Enter the new task")
	task := readLine()
	*tasks = append(*tasks, task)
}

func removeTask(tasks *[]string) {
	fmt.Println("Select the number of the task to remove")
	var taskIndex int32
	fmt.Scanln(&taskIndex)
	*tasks = append((*tasks)[:taskIndex-1], (*tasks)[taskIndex:]...)
}

func saveTasks(tasks []string, fileName string) {
	fileData := []byte(strings.Join(tasks, "\n"))
	err := os.WriteFile(fileName, fileData, 0777)
	if err != nil {
		fmt.Println("Error on writing file")
		fmt.Println(err)
	}
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\r')
	expression := line[:len(line) - 1]

	return expression
}
