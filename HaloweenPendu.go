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
	// Affiche le message de bienvenue
	fmt.Println("BIENVENUE DANS LE JEU DU PENDU EN GO !!!")
	fmt.Println("")

	for {
		// Affiche un texte artistique (ASCII art)
		fmt.Println("██╗   ██╗    ██╗  ██╗ █████╗ ██╗     ██╗      ██████╗ ██╗    ██╗███████╗███████╗███╗   ██╗")
		fmt.Println("██║  ██║██╔══██╗████╗  ██║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║    ██║  ██║██╔══██╗██║     ██║     ██╔═══██╗██║    ██║██╔════╝██╔════╝████╗  ██║")
		fmt.Println("███████║███████║██╔██╗ ██║██║  ███╗██╔████╔██║███████║██╔██╗ ██║    ███████║███████║██║     ██║     ██║   ██║██║ █╗ ██║█████╗  █████╗  ██╔██╗ ██║")
		fmt.Println("██╔══██║██╔══██║██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║    ██╔══██║██╔══██║██║     ██║     ██║   ██║██║███╗██║██╔══╝  ██╔══╝  ██║╚██╗██║")
		fmt.Println("██║  ██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║    ██║  ██║██║  ██║███████╗███████╗╚██████╔╝╚███╔███╔╝███████╗███████╗██║ ╚████║")
		fmt.Println("╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝    ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚══════╝ ╚═════╝  ╚══╝╚══╝ ╚══════╝╚══════╝╚═╝  ╚═══╝")

		// Invite l'utilisateur à choisir un mode de jeu
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

		motADeviner := choisirMotAvecLettres(mots)
		motADeviner = strings.ToLower(motADeviner)
		lettresDevinees := make(map[rune]bool)
		lettresUtilisees := []rune{}
		maxTentatives := 10
		tentativesRestantes := maxTentatives
		pendu := ""

		fmt.Println("Le jeu du pendu commence !")
		fmt.Println("Bonne chance !")

		for {
			// Affiche le pendu actuel
			fmt.Println()
			afficherPendu(pendu)
			// Affiche le mot avec les lettres devinées et non devinées
			afficherMot(motADeviner, lettresDevinees)
			// Affiche les lettres utilisées
			fmt.Printf("Lettres utilisées : %s\n", strings.Join(convertirEnChaines(lettresUtilisees), ", "))

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
			if lettresDevinees[lettre] || contientLettre(lettre, lettresUtilisees) {
				fmt.Println("Vous avez déjà deviné cette lettre.")
				continue
			}

			if strings.ContainsRune(motADeviner, lettre) {
				fmt.Println("Bonne devinette !")
				lettresDevinees[lettre] = true
			} else {
				fmt.Printf("Raté ! La lettre %c n'est pas dans le mot.\n", lettre)
				tentativesRestantes--
				pendu = ajouterEtapePendu(pendu, maxTentatives-tentativesRestantes)
			}

			lettresUtilisees = append(lettresUtilisees, lettre)

			if tentativesRestantes == 0 {
				fmt.Println("Désolé, vous avez épuisé toutes vos tentatives. Le mot était", motADeviner)
				break
			}
		}

		// Demande si l'utilisateur souhaite rejouer
		var reponse string
		fmt.Print("Voulez-vous jouer encore ? (oui/non): ")
		fmt.Scanln(&reponse)
		if reponse != "oui" {
			fmt.Println("Au revoir !")
			return
		}
	}
}

// Choisi un mot aléatoire avec des lettres X et Y
func choisirMotAvecLettres(mots []string) string {
	rand.Seed(time.Now().UnixNano())
	motADeviner := mots[rand.Intn(len(mots))]

	motAvecLettres := motADeviner[:2] + "XY" + motADeviner[4:]

	return motAvecLettres
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

func convertirEnChaines(runes []rune) []string {
	chaines := make([]string, len(runes))
	for i, r := range runes {
		chaines[i] = string(r)
	}
	return chaines
}
