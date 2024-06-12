package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func funcEcho(array []string){
	arg := array[1:]
	for e := range arg {
		fmt.Print(arg[e])
		if (e < len(arg) - 1){
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}

func funcType(array []string){ 
	builtin := []string{"echo", "exit", "type"}
	for i := range builtin{
		if (builtin[i] == array[1]){
			fmt.Println(array[1] + " is a shell builtin")
			return
		}
	}
	path := strings.Split(os.Getenv("PATH"), ":")
	for _ , e := range path {
		_ , err :=  os.Stat( e + "/" + array[1])
		if (err == nil){
			fmt.Println(array[1] + " is " + e + "/" + array[1])
			return
		}
	}
	fmt.Println(array[1] + ": not found")
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
			funcEcho(array)
		} else if (array[0] == "type"){
			funcType(array)
		} else {
			fmt.Println(cmd[:len(cmd)-1] +": command not found")
		}
	}
}
