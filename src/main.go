package main

import "fmt"

func main() {
	player := characterCreation()

	//-------------// BUBBLE TEA //----------------//
	if err := startNavigation(&player); err != nil {
		fmt.Println("Erreur lors du lancement de l'interface :", err)
	}
	//-------------// FIN BUBBLE TEA //----------------//
}
