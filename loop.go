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
	var anneeNessance string
	for {
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

	// =============================================================================
	// 5. Le Switch en Go
	// =============================================================================
	//
	// Le "switch" permet de remplacer une longue chaîne de if/else if/else.
	// Contrairement à C ou Java, Go n'a PAS besoin de "break" à la fin de
	// chaque cas : l'exécution s'arrête automatiquement après le premier cas
	// correspondant (pas de "fall-through" par défaut).
	//
	// Syntaxe :
	//   switch expression {
	//   case valeur1:
	//       // code si expression == valeur1
	//   case valeur2, valeur3:
	//       // code si expression == valeur2 OU valeur3
	//   default:
	//       // code si aucun cas ne correspond
	//   }
	//
	// Particularités de Go :
	//   - Les cas peuvent contenir plusieurs valeurs séparées par des virgules.
	//   - On peut utiliser "switch" SANS expression (comme un if/else if).
	//   - Le mot-clé "fallthrough" force l'exécution du cas suivant.
	// =============================================================================

	// ---------------------------------------------------------
	// 5a. Switch classique sur une valeur
	//     On évalue le jour de la semaine et on affiche un message.
	// ---------------------------------------------------------
	jour := "mercredi"

	switch jour {
	case "lundi":
		fmt.Println("C'est le début de la semaine 😴")
	case "mardi", "mercredi", "jeudi":
		fmt.Println("On est en plein milieu de la semaine 💪")
	case "vendredi":
		fmt.Println("C'est bientôt le week-end ! 🎉")
	case "samedi", "dimanche":
		fmt.Println("C'est le week-end ! 🥳")
	default:
		fmt.Println("Jour inconnu 🤔")
	}

	// ---------------------------------------------------------
	// 5b. Switch SANS expression (remplace if/else if)
	//     Chaque "case" contient une condition booléenne.
	//     Utile pour des comparaisons plus complexes.
	// ---------------------------------------------------------
	age := 25

	switch {
	case age < 13:
		fmt.Println("Vous êtes un enfant.")
	case age < 18:
		fmt.Println("Vous êtes un adolescent.")
	case age < 65:
		fmt.Println("Vous êtes un adulte.")
	default:
		fmt.Println("Vous êtes un senior.")
	}

	// ---------------------------------------------------------
	// 5c. Switch avec initialisation (comme le "if" en Go)
	//     On peut déclarer une variable dans le switch.
	// ---------------------------------------------------------
	switch note := 15; {
	case note >= 16:
		fmt.Println("Très bien !")
	case note >= 14:
		fmt.Println("Bien !")
	case note >= 10:
		fmt.Println("Passable.")
	default:
		fmt.Println("Insuffisant.")
	}

}
