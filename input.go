// Package main démontre les différentes méthodes de lecture d'entrée utilisateur en Go.
//
// Ce programme illustre trois approches courantes pour lire les données
// saisies par l'utilisateur depuis l'entrée standard (stdin) :
//   - fmt.Scan : lecture simple (un mot ou un nombre)
//   - bufio.NewReader : lecture de chaînes longues (avec espaces)
//   - bufio.Scanner : lecture ligne par ligne avec gestion d'erreurs
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main exécute les trois méthodes de saisie utilisateur de manière séquentielle.
func main() {

	// -----------------------------------------------------------------------
	// Méthode 1 : fmt.Scan
	// -----------------------------------------------------------------------
	// fmt.Scan lit l'entrée standard en s'arrêtant au premier espace ou saut
	// de ligne. C'est pratique pour lire un seul mot ou une valeur numérique,
	// mais ne convient pas pour lire une phrase complète contenant des espaces.

	var age int
	fmt.Println("Enter age: ")
	fmt.Scan(&age)
	fmt.Println("votre age : ", age)
	var fulName string
	fmt.Println("Enter full name: ")
	fmt.Scan(&fulName)

	// -----------------------------------------------------------------------
	// Méthode 2 : bufio.NewReader
	// -----------------------------------------------------------------------
	// bufio.NewReader crée un lecteur tamponné autour de os.Stdin.
	// ReadString('\n') lit tout jusqu'au caractère de saut de ligne inclus,
	// ce qui permet de capturer une ligne entière (espaces compris).
	// strings.TrimSpace supprime les espaces et le '\n' en début/fin de chaîne.

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter nom complet : ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)
	fmt.Println("bonjour : ", fullName)

	// -----------------------------------------------------------------------
	// Méthode 3 : bufio.Scanner
	// -----------------------------------------------------------------------
	// bufio.Scanner lit l'entrée ligne par ligne via scanner.Scan().
	// scanner.Text() retourne la dernière ligne lue (sans le '\n').
	// scanner.Err() permet de vérifier si une erreur s'est produite pendant
	// la lecture (par exemple une erreur d'E/S).

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter nom complet : ")
	if scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("je écrire %q\n", line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("il exist un erreur : ", err)
	}
}
