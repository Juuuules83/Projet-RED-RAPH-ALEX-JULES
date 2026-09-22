package main

import (
	"fmt"

	"projet-red/player"
	"projet-red/utils"
)

// main est le point d'entrée unique du projet.
// Toute l'initialisation du jeu part d'ici.
func main() {
	character := player.CharacterCreation()

	if err := utils.StartNavigation(&character); err != nil {
		fmt.Println("Erreur lors du lancement du jeu :", err)
	}
}
