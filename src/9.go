package main

import (
	"math/rand"
	"time"
)

func generateRandomHomework() {
	rand.Seed(time.Now().UnixNano())
	problems := []string{
		"2 + 2 =",
		"3 x 4 =",
		"5 - 2 =",
		"7 / 2 =",
	}
	correctAnswers := []int{
		4,
		12,
		3,
		3.5,
	}
	randIndex := rand.Intn(len(problems))
	fmt.Println(problems[randIndex])
	userAnswer := 0
	for userAnswer != correctAnswers[randIndex] {
		fmt.Scanf("%d", &userAnswer)
		if userAnswer == correctAnswers[randIndex] {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect, please try again.")
		}
	}
}
