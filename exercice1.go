package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func evaluateExperience(anneeExp string) string {
	exp, err := strconv.Atoi(anneeExp)
	if err != nil {
		return "enter juster les nombres!!!"
	}
	if exp <= 1 {
		return "Parfait pour apprendre, tu seras dans l'équipe Débutant !"
	} else if exp > 1 && exp <= 3 {
		return "Super, tu seras dans l'équipe Intermédiaire."
	} else {
		return "Génial, tu peux être Mentor ou Lead technique !"
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var age int
	var techFrontend string
	var techBackend string

	fmt.Print("Enter age : ")
	fmt.Scan(&age)
	//scanner.Scan()
	fmt.Println("Vous êtes probablement né(e) en", 2026-age)
	fmt.Print("Enter Frontend : ")
	if scanner.Scan() {
		techFrontend = scanner.Text()
	}
	fmt.Print("Enter Backend : ")
	if scanner.Scan() {
		techBackend = scanner.Text()
	}
	fmt.Printf("✅ Résumé de votre projet : Vous allez"+
		" développer une application avec %v en frontend et %v en backend.\n\n", techFrontend, techBackend)
	var anneeExp string
	fmt.Printf("Enter combien d'années d'expérience en programmation avez-vous : ")
	fmt.Scan(&anneeExp)
	fmt.Print(evaluateExperience(anneeExp))

}
