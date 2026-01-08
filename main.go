package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numeroSorteado := rand.Intn(101)
	nTentativas := 10
	palpitesFeitos := []int{}

	fmt.Printf("Guess the number, you %d tries.\n", nTentativas)
	contTentativa := 0
	for contTentativa < nTentativas {
		var guess int
		fmt.Printf("Enter your guess (0, 100) - [%d tries remaining]: ", nTentativas-contTentativa)
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}
		for _, palpite := range palpitesFeitos {
			if guess == palpite {
				fmt.Println("You already guessed that number. Try again.")
				continue
			}
		}

		palpitesFeitos = append(palpitesFeitos, guess)

		if guess < 0 || guess > 100 {
			fmt.Println("Please guess a number between 0 and 100.")
			continue
		}

		contTentativa++

		if guess < numeroSorteado {
			fmt.Println("Too low! Guess higher.")
		} else if guess > numeroSorteado {
			fmt.Println("Too high! Guess lower.")
		} else {
			fmt.Println("Congratulations! You've guessed the number! The number was ", numeroSorteado)
			return
		}
	}
	if contTentativa >= nTentativas {
		fmt.Println("Sorry, you've used all your tries. The number was: ", numeroSorteado)
		return
	}
}
