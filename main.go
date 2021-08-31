package main

import (
	"fmt"
	"time"
)

func main() {
	s := time.Now().UnixNano()
	fmt.Println(s)
}
