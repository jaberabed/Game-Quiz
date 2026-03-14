package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to my quiz game!!!")

	name :=  ""
	age := 0
	fmt.Printf("Please enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello %v, let's start the quiz!\n", name)
	fmt.Printf("Please enter your age: ")
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Printf("Sorry %v, you are not old enough to play the quiz.\n", name)
		return
	}
	fmt.Printf("Great! Let's get started.\n")
	questions := []string{
		"What is the capital of France?",
		"What is the largest planet in our solar system?",
		"What is the smallest country in the world?",	
		"What is the chemical symbol for water?",
	}
	answers := []string{
		"Paris",
		"Jupiter",
		"City",
		"H2O",
	}
	score := 0
	for i, question := range questions {
		var answer string
		fmt.Printf("%v: %v\n", i+1, question)
		fmt.Scanln(&answer)
		if strings.EqualFold(strings.TrimSpace(answer), strings.TrimSpace(answers[i])) {
			fmt.Printf("Correct!\n")
			score++
		} else {
			fmt.Printf("Wrong! The correct answer is %v.\n", answers[i])
		}
	}
	fmt.Printf("Your final score is %v out of %v.\n", score, len(questions))
  }
