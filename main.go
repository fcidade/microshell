package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		printShellPointer()

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stderr, "Error: ", err)
		}
		
		if err := execInput(input); err != nil {
			fmt.Println("Error: ", err)
		}
		
	}
}

func printShellPointer() {
	dir, _ := os.Getwd()
	folders := strings.Split(dir, "/")
	dir = folders[len(folders) - 1]
	fmt.Printf("\033[1;36m%s -> ", dir)
}

func execInput(input string) error {
	input = strings.Trim(input, "\n")

	args := strings.Fields(input)

	switch args[0] {
	case "cd":

		var targetDir string

		if len(args) < 2 {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			targetDir = home
		} else {
			targetDir = args[1]
		}

		if err := os.Chdir(targetDir); err != nil {
			return err
		}

		return nil
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	return cmd.Run()
}