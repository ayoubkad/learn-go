package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var age int
	var techFrontend string
	var techBackend string
	var anneeExp string
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
		" développer une application avec %v en frontend et %v en backend.", techFrontend, techBackend)
	fmt.Printf("Enter combien d'années d'expérience en programmation avez-vous ?")
	fmt.Scan(&anneeExp)

	exp, err := strconv.Atoi(anneeExp)
	if err != nil {
		fmt.Println("enter juster les nombres!!!")
		return
	}
	if exp <= 1 {
		fmt.Println("Parfait pour apprendre, tu seras dans l'équipe Débutant !")
	} else if exp > 1 && exp <= 3 {
		fmt.Println("Super, tu seras dans l'équipe Intermédiaire.")
	} else {
		fmt.Println("Génial, tu peux être Mentor ou Lead technique !")
	}

}
