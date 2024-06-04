/*
非常简便的 cli 命令的代码框架
参考 https://darkcoding.net/software/the-simplest-cli-that-could-possibly-work/
*/
package main

import (
	"fmt"
	"os"
)

var commands = map[string]func(){
	"greet": greet,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage())
		os.Exit(1)
	}
	cmd, ok := commands[os.Args[1]]
	if !ok {
		fmt.Println(usage())
		os.Exit(1)
	}
	cmd()
}
func usage() string {
	s := "Usage: sl [cmd] [options]\nCommands:\n"
	for k := range commands {
		s += fmt.Sprintf(" - %s\n", k)
	}
	return s
}

func greet() {
	fmt.Println("Hello World")
}
