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
		if (cmd == "exit 0\n"){
			return
		}
		fmt.Println(cmd[:len(cmd)-1] +": command not found")
	}
}
