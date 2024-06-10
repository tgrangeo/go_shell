package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for{
		fmt.Fprint(os.Stdout, "$ ")
		cmd , _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println(cmd[:len(cmd)-1] +": command not found")
	}
}
