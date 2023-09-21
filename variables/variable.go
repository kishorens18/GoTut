package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

var pr = fmt.Println

func main() {
	pr(reflect.TypeOf(25))
	pr(reflect.TypeOf(2.5))
	pr(reflect.TypeOf("25"))
	pr(reflect.TypeOf(true))

	v1 := 25
	v2 := 2.5
	v3 := v1 - int(v2)
	pr(v3)

	c1 := "3.14"
	if c4, err := strconv.ParseFloat(c1, 64); err == nil {
		pr(c4, reflect.TypeOf(c4))
	} else {
		log.Fatal(err)
	}

	str := "123"
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Parsed integer: %d\n", num)
}
