package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	Text   string
	Answer string
}

func readLine(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}

func main() {

	quiz := []Question{
		{"What is the capital of France?", "Paris"},
		{"What is the largest planet in our solar system?", "Jupiter"},
		{"What is the smallest country in the world?", "Vatican City"},
		{"What is the chemical symbol for water?", "H2O"},
	}

	fmt.Println("Welcome to my quiz game!!!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your name: ")
	name := readLine(reader)

	fmt.Printf("Hello %s, let's start the quiz!\n", name)

	fmt.Print("Please enter your age: ")
	var age int
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Printf("Sorry %s, you are not old enough to play the quiz.\n", name)
		return
	}

	fmt.Println("Great! Let's get started.")

	score := 0

	for i, q := range quiz {

		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Text)
		fmt.Print("Your answer: ")

		answer := readLine(reader)

		if strings.EqualFold(answer, q.Answer) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Wrong! The correct answer is %s.\n", q.Answer)
		}
	}

	fmt.Printf("\nYour final score is %d out of %d.\n", score, len(quiz))
}