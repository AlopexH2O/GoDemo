package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	//	fmt.Println(s)
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], "-"))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
