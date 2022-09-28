package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now())
	current := time.Now()
	fmt.Println(current.Hour())
	fmt.Println(current.Minute())
	fmt.Println(current.Second())

	value, check := os.LookupEnv("GOPATH")
	if check {
		fmt.Fprintln("GOPATH:", value)
	} else {
		fmt.Appendln("GOPATH:", "Not set....")
	}

}
