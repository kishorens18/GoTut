package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

var pr = fmt.Println

func main() {
	// identify data types
	pr(reflect.TypeOf(25))
	pr(reflect.TypeOf(2.5))
	pr(reflect.TypeOf("25"))
	pr(reflect.TypeOf(true))

	//automatically generate type
	v1 := 25
	v2 := 2.5
	v3 := v1 - int(v2)
	pr(v3)

	// str to float manually
	c1 := "3.14"
	if c4, err := strconv.ParseFloat(c1, 64); err == nil {
		pr(c4, reflect.TypeOf(c4))
	} else {
		log.Fatal(err)
	}

	// str to int manually
	str := "123"
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Parsed integer: %d\n", num)

	//str to int using a to i
	str2 := "345"
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Converted integer: %d\n", num2)

	//int to string using i to a
	num3 := 547
	str3 := strconv.Itoa(num3)
	fmt.Printf("converted string: %s\n", str3)
}
