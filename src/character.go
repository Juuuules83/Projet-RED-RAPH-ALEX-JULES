package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Character représente le personnage du joueur.
type Character struct {
	Name       string
	Classe     string
	PvMax      int
	Pv         int
	Niveau     int
	Sorts      []string
	Inventory map[string]int
	Money int
}

// initCharacter initialise les informations du personnage.
func (c *Character) initCharacter(name string, classe string) {
	c.Name = name
	c.Classe = classe
	c.Money = 100
	c.Niveau = 1
	c.Sorts = []string{"Coup de Poing"}

	switch c.Classe {
	case "AI & Data":
		c.PvMax = 100
	case "Info":
		c.PvMax = 80
	case "Cyber":
		c.PvMax = 120
	}
	c.Pv = c.PvMax / 2

	c.Inventory = map[string]int{
		PotionVie:    3,
		PotionPoison: 1,
		"ordinateur": 1,
	}
}


// isLettersOnly vérifie que la chaîne ne contient que des lettres.
func isLettersOnly(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

// formatName met en forme le nom : une majuscule au début, le reste en minuscule.
func formatName(name string) string {
	lower := strings.ToLower(name)
	return strings.ToUpper(lower[:1]) + lower[1:]
}

// characterCreation permet à l'utilisateur de créer son propre personnage.
func characterCreation() Character {
	var player Character

	var name string
	for {
		fmt.Print("Entrez le nom de votre personnage (lettres uniquement) : ")
		fmt.Scan(&name)

		if isLettersOnly(name) {
			break
		}

		fmt.Println(Red + "Le nom ne doit contenir que des lettres." + Reset)
	}

	name = formatName(name)

	var classe string
	for {
		fmt.Println("Choisissez votre classe :")
		fmt.Println(Yellow + "[1]" + Reset + "AI & Data")
		fmt.Println(Yellow + "[2]" + Reset + "Info")
		fmt.Println(Yellow + "[3]" + Reset + "Cyber")
		fmt.Print("> ")

		var choix int
		if _, err := fmt.Scan(&choix); err != nil {
			fmt.Println(Red + "Veuillez entrer un nombre." + Reset)
			continue
		}

		switch choix {
		case 1:
			classe = "AI & Data"
		case 2:
			classe = "Info"
		case 3:
			classe = "Cyber"
		default:
			fmt.Println(Red + "Option invalide. Veuillez réessayer." + Reset)
			continue
		}

		break
	}

	player.initCharacter(name, classe)

	return player
}

