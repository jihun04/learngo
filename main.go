package main

import (
	"fmt"
	"learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "Hello"
	dictionary.Add(baseWord, "It might not be a greeting")
	word, _ := dictionary.Search("Hello")
	fmt.Println(word)
	dictionary.Delete(baseWord)
	err := dictionary.Update(baseWord, "It should be a greeting")
	if err != nil {
		fmt.Println(err)
	}
	word, err = dictionary.Search("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
