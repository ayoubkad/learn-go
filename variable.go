// Package main - Point d'entrée du programme Go.
// Ce fichier est un tutoriel complet sur les bases du langage Go (GoLang).
// Il couvre : les variables, les types de données, les constantes,
// la visibilité (export), les tableaux (arrays) et les slices.
package main

import "fmt"

// ================================================================
// 1. Les Zero Values (Valeurs par défaut)
// ================================================================
// Dans d'autres langages de programmation, si vous créez une variable
// sans lui attribuer de valeur, elle peut contenir des « données
// aléatoires (Garbage) ».
// Mais dans Go, si vous créez une variable sans lui attribuer de valeur,
// elle prendra une « valeur nulle » sûre (appelée "Zero Value").

// ================================================================
// 2. Support UTF-8
// ================================================================
// Le langage GoLang supporte nativement l'UTF-8, ce qui permet
// d'écrire directement en arabe, en emoji et dans d'autres alphabets.

// ================================================================
// 3. Résumé des Types de Données Fondamentaux en Go (GoLang)
// ================================================================
//
//  1. Booléens (Booleans) :
//     - Type : bool
//     - Valeurs : true (vrai) ou false (faux).
//     - Valeur par défaut (Zero Value) : false
//
//  2. Chaînes de caractères (Strings) :
//     - Type : string
//     - Description : Séquence de caractères entre guillemets doubles "".
//     - Valeur par défaut : "" (chaîne vide).
//
//  3. Entiers (Integers) :
//     - Type : int (le plus utilisé), int8, int16, int32, int64.
//     - Description : Nombres entiers sans virgule (positifs ou négatifs).
//     - Type : uint (Unsigned Integer) -> entiers positifs uniquement.
//     - Valeur par défaut : 0
//
//  4. Nombres à virgule flottante (Floating Point) :
//     - Type : float64 (standard et précis), float32 (moins de mémoire).
//     - Description : Nombres avec décimales (ex: 10.5, 3.14).
//     - Valeur par défaut : 0.0
//
//  5. Types composites (à voir plus tard) :
//     - Tableaux et Slices (Listes)
//     - Maps (Dictionnaires Clé:Valeur)
//     - Structs (Objets personnalisés)

// printFive affiche un tableau de 5 entiers.
// Note : En Go, la taille d'un tableau fait partie de son type.
// Donc [5]int et [10]int sont deux types différents et incompatibles.
// C'est pour cela qu'on ne peut pas passer un [10]int à cette fonction.
func printFive(arr [5]int) {
	fmt.Println(arr)
}

// main est la fonction principale, point d'entrée du programme.
func main() {

	// ============================================================
	// SECTION 1 : Déclaration de Variables
	// ============================================================
	// Il existe 3 façons de déclarer une variable en Go :
	//   a) var nom type = valeur   → déclaration complète (type explicite)
	//   b) var nom = valeur        → le type est déduit automatiquement (inférence)
	//   c) nom := valeur           → syntaxe courte (uniquement dans une fonction)

	var age int = 25            // (a) Déclaration avec type explicite : int
	var username string = "Ali" // (a) Déclaration avec type explicite : string
	var isActive bool = true    // (a) Déclaration avec type explicite : bool
	//var password = "123456"            // (b) Déclaration sans type → Go déduit "string"
	email := "ayoubkaddioui32@gmail.com" // (c) Syntaxe courte := (autorisée uniquement dans une fonction, pas en global)

	// Si aucune valeur n'est assignée, Go utilise la Zero Value du type.
	var prenom string // Zero Value de string = "" (chaîne vide)

	// Déclaration multiple sur une seule ligne pour organiser le code.
	var x, z, y = 1, 2.5, 3 // x=int, z=float64, y=int (types déduits)

	println(age, username, prenom, email, isActive)
	println(x)
	println(z)
	println(y)

	// ============================================================
	// SECTION 2 : Conversion de Types (Type Casting)
	// ============================================================
	// En Go, on ne peut PAS mélanger les types dans une opération.
	// Il faut convertir explicitement : float64(y) convertit y (int) en float64.

	var somme = float64(y) + z // Conversion nécessaire : int + float64 est interdit
	println(somme)

	// ============================================================
	// SECTION 3 : Les Constantes (Constants)
	// ============================================================
	// Les constantes sont des valeurs immuables définies avec le mot-clé "const".
	// Elles ne peuvent pas être modifiées après leur déclaration.

	const Pi = 3.14           // Constante numérique (float64 par inférence)
	const AppName = "MyGoApp" // Constante de type string

	// ============================================================
	// SECTION 4 : Visibilité / Export (Exporting / Visibility)
	// ============================================================
	// En Go, la visibilité est déterminée par la casse de la première lettre :
	//   - Majuscule (ex: Password) → Public (exporté, visible depuis d'autres packages)
	//   - Minuscule (ex: password) → Privé (accessible uniquement dans ce package)

	var Password string // Public  : visible depuis l'extérieur du package
	var password string // Privé   : réservé uniquement à ce package
	println(Password)
	println(password)

	// ============================================================
	// SECTION 5 : Les Tableaux (Arrays)
	// ============================================================
	// Un tableau a une taille FIXE définie à la déclaration.
	// Syntaxe : var nom [taille]type = [taille]type{valeurs}
	// La taille fait partie du type → [5]int ≠ [10]int

	var small [5]int = [5]int{1, 2, 3, 4, 5}
	var big [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	printFive(small) // ✅ Fonctionne : small est de type [5]int
	//printFive(big)   // ❌ Erreur : big est de type [10]int, incompatible avec [5]int
	fmt.Println(big)

	// ============================================================
	// SECTION 6 : Les Slices
	// ============================================================
	// Un slice est comme un tableau mais avec une taille DYNAMIQUE (flexible).
	// Syntaxe : var nom []type = []type{valeurs}
	// Contrairement aux tableaux, on ne précise pas la taille entre les crochets.
	// Les slices sont plus utilisés que les tableaux en pratique.

	var myFriends []string = []string{"Ali", "Sara", "Ayoub"}
	myNum := []string{"Ali", "Sara", "Ayoub"} // Syntaxe courte pour un slice

	fmt.Println(myNum)
	fmt.Println(myFriends)
}
