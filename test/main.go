package main

import "fmt"

func main() {
	s := "hello world"
	bytes := []byte(s)

	new := bytes[:6]
	new = append(new, new...)
	fmt.Println(len(bytes), string(new), s)
}
