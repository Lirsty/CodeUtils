package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Lirsty/CodeUtils/codereader"
)

func main() {
	f, err := ioutil.ReadFile("test/main.go")
	if err != nil {
		panic(err)
	}
	c := codereader.New(f)

	c.Func["main"].AddLine("    //ABCDEFG")
	c.Func["main"].Remove("Event.printFuncArgs()")
	c.Func["Event.interface2Func"].AddLine("	//TestTestTest")
	fmt.Println(string(c.Write()))
}
