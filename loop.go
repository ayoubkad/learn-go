// =============================================================================
// loop.go — Les boucles en Go
// =============================================================================
//
// En Go, il n'existe qu'un seul mot-clé de boucle : "for".
// Contrairement à C, Java ou Python, il n'y a pas de "while" ni de "do-while".
// Le mot-clé "for" couvre tous les cas :
//
//   1. Boucle classique      : for init; condition; post { ... }
//   2. Boucle "while"        : for condition { ... }
//   3. Boucle infinie        : for { ... }
//
// Ce fichier illustre ces trois formes avec des exemples concrets.
// =============================================================================

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ---------------------------------------------------------
	// 1. Boucle classique (comme "for" en C / Java)
	//    Syntaxe : for initialisation; condition; post-instruction { corps }
	//    Ici on affiche les nombres de 1 à 9.
	// ---------------------------------------------------------
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// ---------------------------------------------------------
	// 2. Boucle "while" (condition seule)
	//    Syntaxe : for condition { corps }
	//    Équivalent de "while" en C/Java ou Python.
	//    Ici on multiplie `power` par 100 tant qu'il est < 100.
	// ---------------------------------------------------------
	power := 1
	for power < 100 {
		fmt.Println(power)
		power *= 100
	}

	// ---------------------------------------------------------
	// 3. Boucle infinie
	//    Syntaxe : for { corps }
	//    S'exécute sans fin jusqu'à un "break" ou "return".
	//
	//    Exemple commenté — compteur infini :
	//    counter := 1
	//    for {
	//        fmt.Println(counter)
	//        counter++
	//    }
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// 4. Exemple pratique : validation d'entrée utilisateur
	//    On demande l'année de naissance en boucle jusqu'à ce
	//    que l'utilisateur saisisse un nombre valide.
	//    - fmt.Scan  : lit l'entrée standard (clavier)
	//    - strconv.Atoi : convertit une chaîne en entier (int)
	//      Retourne une erreur si la chaîne n'est pas un nombre.
	//    - break : sort de la boucle infinie quand l'entrée est valide.
	// ---------------------------------------------------------
	for {
		var anneeNessance string
		fmt.Print("Enter Année de naissance : ")
		fmt.Scan(&anneeNessance)

		_, err := strconv.Atoi(anneeNessance)
		if err != nil {
			fmt.Println("Erreur : veuillez saisir uniquement des chiffres.")
		} else {
			fmt.Println("Une année de naissance acceptable !")
			break // ← on quitte la boucle infinie
		}
	}

}
