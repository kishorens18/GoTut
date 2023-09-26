package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pr = fmt.Println

func main() {
	pr("What is your name?")

	// var name string

	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err == nil {
		pr("Hello", name)
	} else {
		log.Fatal(err)
	}

	// fmt.Fscanf(os.Stdin, "%v", &name)
	// pr("hello", name)

}
