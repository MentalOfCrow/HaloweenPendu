package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Charger les mots depuis un fichier
	mots, err := chargerMotsDepuisFichier("text.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	rand.Seed(time.Now().UnixNano())

	// Sélectionnez un mot au hasard
	motADeviner := mots[rand.Intn(len(mots))]
	motADeviner = strings.ToLower(motADeviner) // Convertir en minuscules pour la comparaison
	lettresDevinees := make(map[rune]bool)
	maxTentatives := 6
	tentativesRestantes := maxTentatives

	fmt.Println("Bienvenue dans le jeu du pendu en Go!")

	for {
		fmt.Println()
		afficherMot(motADeviner, lettresDevinees)

		if victoire(motADeviner, lettresDevinees) {
			fmt.Println("Félicitations, vous avez gagné ! Le mot était", motADeviner)
			break
		}

		fmt.Printf("Tentatives restantes: %d\n", tentativesRestantes)
		fmt.Print("Devinez une lettre: ")
		var tentative string
		fmt.Scanln(&tentative)

		if len(tentative) != 1 || !estLettreValide(tentative) {
			fmt.Println("Entrez une seule lettre valide à la fois.")
			continue
		}

		lettre := rune(tentative[0])
		if lettresDevinees[lettre] {
			fmt.Println("Vous avez déjà deviné cette lettre.")
			continue
		}

		if strings.ContainsRune(motADeviner, lettre) {
			fmt.Println("Bonne devinette !")
			lettresDevinees[lettre] = true
		} else {
			fmt.Printf("Raté ! La lettre %c n'est pas dans le mot.\n", lettre)
			tentativesRestantes--
		}

		if tentativesRestantes == 0 {
			fmt.Println("Désolé, vous avez épuisé toutes vos tentatives. Le mot était", motADeviner)
			break
		}
	}
}

func afficherMot(mot string, lettresDevinees map[rune]bool) {
	for _, c := range mot {
		if c == ' ' {
			fmt.Print("  ")
		} else if lettresDevinees[rune(c)] {
			fmt.Printf("%c ", c)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func estLettreValide(lettre string) bool {
	return len(lettre) == 1 && 'a' <= lettre[0] && lettre[0] <= 'z'
}

func chargerMotsDepuisFichier(filename string) ([]string, error) {
	var mots []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mot := strings.TrimSpace(scanner.Text())
		mots = append(mots, mot)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return mots, nil
}

func victoire(mot string, lettresDevinees map[rune]bool) bool {
	for _, c := range mot {
		if c != ' ' && !lettresDevinees[rune(c)] {
			return false
		}
	}
	return true
}
