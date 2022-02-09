package main

import (
	"fmt"
	"learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	err := dictionary.Add("food", "kimchi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary)
	}
}
