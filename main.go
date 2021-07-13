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

	c.Func["Event.printFuncArgs"].AddLine("    //TestTestTest")
	fmt.Println(string(c.Write()))
}
