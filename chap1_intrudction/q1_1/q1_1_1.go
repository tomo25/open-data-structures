package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	fmt.Println(string(b))
}
