package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Number of args: " + strconv.Itoa(len(os.Args)-1))
	inefficientEcho()
	joinEcho()
}

func inefficientEcho() {
	start := time.Now()
	args := os.Args
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Printf("FOR LOOP:\n%s\n%.5fs elapsed\n\n", s, time.Since(start).Seconds())
}

func joinEcho() {
	start := time.Now()
	args := os.Args[1:]
	fmt.Printf("JOIN:\n%s\n%.5fs elapsed\n\n", strings.Join(args, " "), time.Since(start).Seconds())
}
