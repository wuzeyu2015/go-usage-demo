package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	foostr:=flag.NewFlagSet("str",flag.ExitOnError)
	strValue:=foostr.String("a","string","打印字符串")
	intValue:=foostr.Int("b",1,"打印数值")

	if len(os.Args)<1{
		fmt.Println("expected'str'subcommands")
		os.Exit(1)
		}
	switch os.Args[1]{
		case"str":
			foostr.Parse(os.Args[2:])
			fmt.Println("a",*strValue)
			fmt.Println("b",*intValue)
		default:
			fmt.Println("expected'str'subcommands")
			os.Exit(1)
	}
}