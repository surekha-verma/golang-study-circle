package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	subCommand := os.Args[1]
	if subCommand == "add" {
		if os.Args[2] != "" {
			addTodo()
		}
	}

	if subCommand == "print" {
		printAllTodo()
	}

	if subCommand == "mark-done" {
		markDone()
	}

}

func addTodo() {
	todoItem := fmt.Sprintf("%s#1\n", os.Args[2])
	file, err := os.OpenFile("todo.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("not able to open file todo.txt: " + err.Error())
	}

	defer file.Close()

	_, err = file.WriteString(todoItem)
	if err != nil {
		log.Fatal("error at the time of writing to file: " + err.Error())
	}
}

func printAllTodo() {
	file, err := os.OpenFile("todo.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("not able to open file todo.txt: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		msg := scanner.Text()
		todo := strings.Split(msg, "#")
		if todo[1] == "1" {
			msg = fmt.Sprintf("[✓] %s", todo[0])
		} else {
			msg = fmt.Sprintf("[ ] %s", todo[0])
		}
		fmt.Println(msg)
	}
}

func markDone() {
	todoSeq, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("invalid todo sequence argument: " + err.Error())
	}
	bytesFromFile, err := ioutil.ReadFile("todo.txt")
	if err != nil {
		log.Fatal("not able to open file todo.txt: " + err.Error())
	}
	fileContent := string(bytesFromFile)
	lines := strings.Split(fileContent, "\n")
	line := strings.Split(lines[todoSeq-1], "#")
	lines[todoSeq-1] = line[0] + "#1"

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("todo.txt", []byte(output), 0644)
	if err != nil {
		log.Fatal("error while overwriting the file " + err.Error())
	}
}


