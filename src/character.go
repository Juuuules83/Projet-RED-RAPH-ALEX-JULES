package main

// Character représente le personnage du joueur.
type Character struct {
	Name       string
	Classe     string
	PvMax      int
	Pv         int
	Inventory map[string]int
	Money int
}

// initCharacter initialise les informations du personnage.
func (c *Character) initCharacter(name string, classe string) {
	c.Name = name
	c.Classe = classe
	c.Money = 100

	switch c.Classe {
	case "mentor":
		c.PvMax = 200
		c.Pv = c.PvMax / 2
	case "étudiant":
		c.PvMax = 50
		c.Pv = c.PvMax / 2
	}

	c.Inventory = map[string]int{
		PotionVie:    3,
		"ordinateur": 1,
	}
}
