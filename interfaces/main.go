package main

import (
	"fmt"
)
type bot interface {
	greeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main(){

	en := englishBot{}
	sp := spanishBot{}

	PrintGreetng(en)
	PrintGreetng(sp)
}

func PrintGreetng(b bot){
	fmt.Println(b.greeting())
}

func (englishBot) greeting() string {
	return "Hello!!!.."
}

func (spanishBot) greeting() string {
	return "Holla!!!.."
}