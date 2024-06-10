package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	cmd , _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(cmd[:len(cmd)-1] +": command not found")
}
