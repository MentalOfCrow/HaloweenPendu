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
	fmt.Println("BIENVENUE DANS LE JEU DU PENDU EN GO !!!")
	fmt.Println("")

	for {
		fmt.Println("██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗    ██╗  ██╗ █████╗ ██╗     ██╗      ██████╗ ██╗    ██╗███████╗███████╗███╗   ██╗")
		fmt.Println("██║  ██║██╔══██╗████╗  ██║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║    ██║  ██║██╔══██╗██║     ██║     ██╔═══██╗██║    ██║██╔════╝██╔════╝████╗  ██║")
		fmt.Println("███████║███████║██╔██╗ ██║██║  ███╗██╔████╔██║███████║██╔██╗ ██║    ███████║███████║██║     ██║     ██║   ██║██║ █╗ ██║█████╗  █████╗  ██╔██╗ ██║")
		fmt.Println("██╔══██║██╔══██║██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║    ██╔══██║██╔══██║██║     ██║     ██║   ██║██║███╗██║██╔══╝  ██╔══╝  ██║╚██╗██║")
		fmt.Println("██║  ██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║    ██║  ██║██║  ██║███████╗███████╗╚██████╔╝╚███╔███╔╝███████╗███████╗██║ ╚████║")
		fmt.Println("╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝    ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚══════╝ ╚═════╝  ╚══╝╚══╝ ╚══════╝╚══════╝╚═╝  ╚═══╝")
		fmt.Println("Bienvenue dans le jeu du pendu ! Choisissez un mode :")
		fmt.Println("1. Mode Facile")
		fmt.Println("2. Mode Moyen")
		fmt.Println("3. Mode Difficile")
		fmt.Println("4. Mode Halloween")
		fmt.Println("5. Quitter")

		var choixMode string
		fmt.Scanln(&choixMode)

		switch choixMode {
		case "1":
			fmt.Println("Vous avez choisi le Mode Facile !")
			jouerJeu("TextAFacile.txt")
		case "2":
			fmt.Println("Vous avez choisi le Mode Moyen !")
			jouerJeu("TextBMoyen.txt")
		case "3":
			fmt.Println("Vous avez choisi le Mode Difficile !")
			jouerJeu("TextCDifficile.txt")
		case "4":
			fmt.Println("Vous avez choisi le Mode Halloween !")
			jouerJeu("TextDHalloween.txt")
		case "5":
			fmt.Println("Au revoir !")
			return
		}
	}
}

func jouerJeu(nomFichier string) {
	for {
		mots, err := chargerMotsDepuisFichier(nomFichier)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier :", err)
			return
		}

		rand.Seed(time.Now().UnixNano())

		motADeviner := choisirMot(mots)
		motADeviner = strings.ToLower(motADeviner)
		lettresDevinees := make(map[rune]bool)
		lettresUtilisees := []rune{}
		maxTentatives := 10
		tentativesRestantes := maxTentatives
		pendu := ""

		fmt.Println("Le jeu du pendu commence !")
		fmt.Println("Bonne chance !")

		for {
			fmt.Println()
			afficherPendu(pendu)
			afficherMot(motADeviner, lettresDevinees)
			fmt.Printf("Lettres utilisées : %s\n", formatLettresUtilisees(lettresUtilisees))

			if victoire(motADeviner, lettresDevinees) {
				fmt.Println("Félicitations, vous avez gagné ! Le mot était", motADeviner)
				break
			}

			fmt.Printf("Tentatives restantes: %d\n", tentativesRestantes)
			fmt.Print("Devinez une lettre ou un mot: ")
			var tentative string
			fmt.Scanln(&tentative)

			if len(tentative) > 1 {
				if tentativesRestantes > 2 {
					if tentative == motADeviner {
						fmt.Println("Félicitations, vous avez deviné le mot ! Le mot était", motADeviner)
						break
					} else {
						fmt.Println("Désolé, ce n'est pas le bon mot.")
						tentativesRestantes -= 2
					}
				} else {
					fmt.Println("Nombre de tentatives restantes insuffisant pour proposer un mot.")
				}
				continue
			}

			lettre := rune(tentative[0])
			if !estLettreValide(tentative) {
				fmt.Println("Entrez une seule lettre valide à la fois.")
				continue
			}

			if contientLettre(lettre, lettresUtilisees) {
				fmt.Printf("Vous avez déjà utilisé la lettre %c\n", lettre)
				continue
			}

			lettreTrouvee := false
			for _, c := range motADeviner {
				if rune(c) == lettre {
					lettreTrouvee = true
					lettresDevinees[lettre] = true
				}
			}

			if lettreTrouvee {
				fmt.Println("Bonne devinette !")
			} else {
				fmt.Printf("Raté ! La lettre %c n'est pas dans le mot.\n", lettre)
				tentativesRestantes--
				pendu = ajouterEtapePendu(pendu, maxTentatives-tentativesRestantes)
				if tentativesRestantes == 0 {
					fmt.Println("Désolé, vous avez épuisé toutes vos tentatives. Le mot était", motADeviner)
					break
				}
			}

			lettresUtilisees = append(lettresUtilisees, lettre)
		}

		var reponse string
		fmt.Print("Voulez-vous jouer encore ? (oui/non): ")
		fmt.Scanln(&reponse)
		if reponse != "oui" {
			fmt.Println("Au revoir !")
			return
		}
	}
}

func formatLettresUtilisees(lettres []rune) string {
	var result strings.Builder
	for i, lettre := range lettres {
		if i != 0 {
			result.WriteString(", ")
		}
		result.WriteRune(lettre)
	}
	return result.String()
}

func choisirMot(mots []string) string {
	rand.Seed(time.Now().UnixNano())
	return mots[rand.Intn(len(mots))]
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

func afficherPendu(pendu string) {
	fmt.Print(pendu)
}

func ajouterEtapePendu(pendu string, etape int) string {
	etapePendu := []string{
		"",
		"      |  \n      |  \n      |  \n      |  \n      |  \n=========  \n\n",
		"  +---+  \n      |  \n      |  \n      |  \n      |  \n      |  \n=========  \n\n",
		"  +---+  \n  |   |  \n      |  \n      |  \n      |  \n      |  \n=========  \n\n",
		"  +---+  \n  |   |  \n  O   |  \n      |  \n      |  \n      |  \n=========  \n\n",
		"  +---+  \n  |   |  \n  O   |  \n  |   |  \n      |  \n      |  \n=========  \n\n",
		"  +---+  \n  |   |  \n  O   |  \n /|   |  \n      |  \n      |  \n=========  \n\n",
		"  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n      |  \n      |  \n=========  \n\n",
		"  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n /    |  \n      |  \n=========  \n\n",
		"  +---+  \n  |   |  \n  O   |  \n /|\\  |  \n / \\  |  \n      |  \n=========\n\n",
	}

	if etape < len(etapePendu) {
		return pendu + etapePendu[etape]
	}

	return pendu
}

func contientLettre(lettre rune, lettres []rune) bool {
	for _, l := range lettres {
		if l == lettre {
			return true
		}
	}
	return false
}
