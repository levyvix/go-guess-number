package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := rand.Intn(101)
	nTries := 10
	palpites := []int{}

	fmt.Printf("Guess the number, you %d tries.\n", nTries)
	tryCount := 0
	for tryCount < nTries {
		var guess int
		fmt.Printf("Enter your guess (0, 100) - [%d tries remaining]: ", nTries-tryCount)
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}
		alreadyGuessed := false
		for i := range palpites {
			if guess == palpites[i] {
				fmt.Println("You already guessed that number. Try again.")
				alreadyGuessed = true
			}
		}
		if alreadyGuessed {
			continue
		}
		palpites = append(palpites, guess)

		if guess < 0 || guess > 100 {
			fmt.Println("Please guess a number between 0 and 100.")
			continue
		}

		tryCount++

		if guess < number {
			fmt.Println("Too low!")
		} else if guess > number {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Congratulations! You've guessed the number! The number was ", number)
			return
		}
	}
	if tryCount >= nTries {
		fmt.Println("Sorry, you've used all your tries. The number was", number)
		return
	}

}
