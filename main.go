package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spainshBot struct{}

func main() {
	eb := englishBot{}
	sb := spainshBot{}

	printGreeting(eb)
	printGreeting(sb)

}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func printGreeting(sb spainshBot) {
//	fmt.Println(sb.getGreeting()

func (eb englishBot) getGreeting() string {
	//very custom logic for generating aneng greeting
	return "hi there!"
}
func (sb spainshBot) getGreeting() string {
	return "Hola"
}
