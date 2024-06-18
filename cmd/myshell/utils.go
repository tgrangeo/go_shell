package main

import (
	"fmt"
	"os"
	"strings"
)

func showPrefix(){
	dir, _ := os.Getwd()
	array := strings.Split(dir, "/")
	fmt.Print("\033[32m ➜ \033[35m" + array[len(array) - 1] + "\033[33m ✗ \033[0m")
}