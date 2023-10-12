package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var mots = []string{
	"Sorcière",
	"Fantôme",
	"Loup-garou",
	"Vampire",
	"Momie",
	"Zombie",
	"Chauve-souris",
	"Citrouille",
	"Goule",
	"Squelette",
	"Mort",
	"Cimetière",
	"Chaudron",
	"Balai",
	"Magicien",
	"Crâne",
	"Potion",
	"Épouvantail",
	"Frankenstein",
	"Effrayant",
	"Terrifiant",
	"Horrifiant",
	"Épouvantable",
	"Monstrueux",
	"Mortel",
	"Mort-vivant",
	"Mortuaire",
	"Mortifère",
	"Sanglant",
	"Nuit",
	"Sorcellerie",
	"Tombe",
	"Magie noire",
	"Hanté",
	"Ténébreux",
	"Diable",
	"Maquillage",
	"Château hanté",
	"Déguisement",
}

func choisirMotAleatoire() string {
	rand.Seed(time.Now().UnixNano())
	return mots[rand.Intn(len(mots))]
}

func afficherMot(mot, lettresTrouvées string) {
	for _, lettre := range mot {
		if strings.ContainsRune(lettresTrouvées, lettre) {
			fmt.Printf("%c ", lettre)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("Bienvenue au jeu du pendu!")

	motSecret := choisirMotAleatoire()
	lettresTrouvées := ""
	essaisRestants := 6

	for essaisRestants > 0 {
		fmt.Printf("\nMot à deviner : ")
		afficherMot(motSecret, lettresTrouvées)

		fmt.Printf("Lettres déjà utilisées : %s\n", lettresTrouvées)
		fmt.Printf("Essais restants : %d\n", essaisRestants)

		var lettre string
		fmt.Print("Entrez une lettre : ")
		fmt.Scan(&lettre)

		if len(lettre) != 1 || !strings.ContainsAny(lettre, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			fmt.Println("Veuillez entrer une seule lettre valide.")
			continue
		}

		if strings.ContainsRune(motSecret, rune(lettre[0])) {
			lettresTrouvées += lettre
			if strings.Contains(motSecret, lettresTrouvées) {
				fmt.Printf("\nFélicitations, vous avez deviné le mot : %s!\n", motSecret)
				break
			}
		} else {
			essaisRestants--
			fmt.Printf("\nLa lettre %s n'est pas dans le mot. Il vous reste %d essais.\n", lettre, essaisRestants)
		}
	}

	if essaisRestants == 0 {
		fmt.Printf("Désolé, vous avez perdu. Le mot était : %s\n", motSecret)
	}
}
