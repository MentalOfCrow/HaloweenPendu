package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Bienvenue au jeu des allumettes!")

	var n int
	fmt.Print("Veuillez entrer le nombre d'allumettes initial (minimum 4) : ")
	fmt.Scan(&n)

	if n < 4 {
		fmt.Println("Le nombre d'allumettes initial doit être d'au moins 4.")
		return
	}

	playerTurn := true

	for n > 0 {
		fmt.Printf("\nIl reste %d allumettes.\n", n)

		if playerTurn {
			var playerChoice int
			fmt.Print("Combien d'allumettes souhaitez-vous prendre (1-3) ? ")
			fmt.Scan(&playerChoice)

			if playerChoice < 1 || playerChoice > 3 || playerChoice > n {
				fmt.Println("Choix invalide. Veuillez choisir entre 1 et 3 allumettes.")
				continue
			}

			n -= playerChoice
			playerTurn = false
		} else {
			aiChoice := rand.Intn(3) + 1
			fmt.Printf("L'IA prend %d allumettes.\n", aiChoice)
			n -= aiChoice
			playerTurn = true
		}
	}

	if playerTurn {
		fmt.Println("\nL'IA a pris la dernière allumette. Vous avez gagné!")
	} else {
		fmt.Println("\nVous avez pris la dernière allumette. L'IA a gagné!")
	}
}
