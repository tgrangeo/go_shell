package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func echo(array []string){
	arg := array[1:]
	for e := range arg {
		fmt.Print(arg[e])
		if (e < len(arg) - 1){
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}

func main() {
	for{
		fmt.Fprint(os.Stdout, "$ ")
		cmd , _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if (cmd == "exit 0\n"){
			return
		}
		str := strings.TrimSuffix(cmd, "\n")
		array := strings.Split(str, " ")
		if (array[0] == "echo"){
			echo(array)
		} else {
			fmt.Println(cmd[:len(cmd)-1] +": command not found")
		}
	}
}
