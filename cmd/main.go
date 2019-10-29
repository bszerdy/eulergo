package main

import (
	"fmt"
	"os"
	"reflect"

	_ "github.com/bszerdy/eulergo/solutions"
)

func main() {
	solution := os.Args

	if len(solution) < 2 {
		fmt.Println("Nothing to run")
		os.Exit(0)
	}

	m := map[string]string{"problem001": "problem001"}

	in := make([]reflect.Value, len(m))

	f := reflect.ValueOf(solution[1])
	a := f.Call(in)

	fmt.Printf("answer: %+v", a[0])
}
