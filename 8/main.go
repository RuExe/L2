package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Гони команды")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Fields(line)[0]
		args := strings.Fields(line)

		switch command {
		case "cd":
			cd(args)
		case "pwd":
			path, _ := os.Getwd()
			fmt.Println(path)
		case "echo":
			echo(args)
		case "kill":

		case "ps":

		default:
			fmt.Println("Данной команды нет")
		}
	}
}

func cd(args []string) {
	if len(args) < 2 {
		fmt.Println("cd: Аргументов не хватает")
		return
	}
	err := os.Chdir(args[1])
	if err != nil {
		fmt.Println(err)
	}
}

func echo(args []string) {
	if len(args) < 2 {
		fmt.Println("echo: Аргументов не хватает")
		return
	}
	fmt.Println(args[1])
}
