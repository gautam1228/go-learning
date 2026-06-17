package main

import (
	"fmt"
	"time"
)

func main() {

	num := 5

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("other")
	}

	// multi switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("yayy it's weekend :)")
	default:
		fmt.Println("it's a workday :(")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer")
		case bool:
			fmt.Println("Boolean")
		case string:
			fmt.Println("String")
		default:
			fmt.Println("Other", t)
		}
	}

	// whoAmI(1)
	// whoAmI(true)
	whoAmI("hello")
}
