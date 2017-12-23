package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const version = "0.1"
const debug = false

func divideCommandAndArgs(input string) (string, []string) {
	words := strings.Split(strings.TrimRight(input, "\n"), " ")
	command := strings.Trim(words[0], "\"")
	args := make([]string, 0)
	for _, word := range words[1:] {
		if len(word) > 0 {
			args = append(args, strings.Trim(word, "\""))
		}
	}

	return command, args
}

func extractExistingPaths(args []string) []string {
	paths := make([]string, 0)
	for _, arg := range args {
		if _, error := os.Stat(arg); error == nil {
			paths = append(paths, arg)
		}
	}

	return paths
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	prompt := fmt.Sprintf("vanish-%s$ ", version)
	for {
		fmt.Print(prompt)
		input, error := reader.ReadString('\n')
		if error != nil {
			panic(error)
		}

		command, args := divideCommandAndArgs(input)
		if command == "exit" {
			fmt.Println("exit")
			break
		}

		paths := extractExistingPaths(args)
		if len(paths) == 0 {
			continue
		}

		rm := exec.Command("rm", "-rf", strings.Join(paths, " "))
		if debug {
			fmt.Println("Would execute:", strings.Join(rm.Args, " "))
		} else {
			rm.Run()
			fmt.Println("😇  vanish 😇")
		}
	}
}
