package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func funcEcho(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func funcType(args []string) {
	builtins := map[string]struct{}{
		"echo": {}, "exit": {}, "type": {}, "cd": {},
	}
	if _, found := builtins[args[1]]; found {
		fmt.Println(args[1] + " is a shell builtin")
		return
	}

	pathDirs := strings.Split(os.Getenv("PATH"), ":")
	for _, dir := range pathDirs {
		if _, err := os.Stat(filepath.Join(dir, args[1])); err == nil {
			fmt.Println(args[1] + " is " + filepath.Join(dir, args[1]))
			return
		}
	}
	fmt.Println(args[1] + ": not found")
}

func funcExec(args []string) bool {
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	fmt.Printf("%s", output)
	return true
}

func funcCd(args []string) {
	if len(args) < 2 {
		fmt.Println("cd: missing argument")
		return
	}
	if args[1] == "~" {
		args[1] = os.Getenv("HOME")
	}
	err := os.Chdir(args[1])
	if err != nil {
		fmt.Println("cd: " + args[1] + ": No such file or directory")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		showPrefix();
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		cmd = strings.TrimSpace(cmd)
		if cmd == "exit 0" {
			return
		}

		args := strings.Split(cmd, " ")
		switch args[0] {
		case "echo":
			funcEcho(args)
		case "type":
			funcType(args)
		case "cd":
			funcCd(args)
		default:
			if !funcExec(args) {
				fmt.Println(args[0] + ": command not found")
			}
		}
	}
}
