package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Print(time.Now().Add(24 * time.Hour).Unix())
}
