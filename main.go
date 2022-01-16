package main

import (
	"fmt"

	"github.com/hlmerscher/go-playground/set"
)

func main() {
	words := set.New[string]()
	words.Add("hello")
	words.Add("world")
	words.Add("world")
	words.Add("world")
	words.Add("golang")
	words.Add("oops")
	words.Add("oops")
	words.Del("world")
	fmt.Printf("words:\t\t%s\n", words)
	fmt.Printf("size:\t\t%d\n", words.Len())

	numbers := set.New[int]()
	numbers.Add(1)
	numbers.Add(1)
	numbers.Add(7)
	numbers.Add(666)
	numbers.Del(1)
	fmt.Printf("numbers:\t%s\n", numbers)
	fmt.Printf("size:\t\t%d\n", numbers.Len())

	fmt.Println("immutable set")
	immuset := set.NewImmutable[string]()
	fmt.Printf("immuset:\t%q\n", immuset)
	immuset = immuset.Add("hello")
	immuset = immuset.Add("world")
	immuset = immuset.Add("hercules")
	fmt.Printf("immuset:\t%q\n", immuset)
	immuset = immuset.Del("world")
	fmt.Printf("words:\t\t%s\n", immuset)
	fmt.Printf("size:\t\t%d\n", immuset.Len())
}
