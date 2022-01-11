package main

import (
	"context"
	"fmt"
)

type paramKey struct{}

func main() {
	c := context.WithValue(context.Background(), paramKey{}, "abc")
	mainTask(c)
}

func mainTask(c context.Context) {
	fmt.Printf("main task started with param %q\n", c.Value(paramKey{}))
	smallTask(c, "task1")
	smallTask(c, "task2")
}

func smallTask(c context.Context, name string) {
	fmt.Printf("%s task started with param %q\n", name, c.Value(paramKey{}))
	fmt.Printf("%s started\n", name)
}
