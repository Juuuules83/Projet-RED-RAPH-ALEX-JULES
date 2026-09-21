package main

//-------------------// STRUCTURE MONSTER + INIT GOBLIN //-------------------//
// Monster représente un adversaire du joueur
type Monster struct {
	Name   string
	PvMax  int
	Pv     int
	Attack int
}

// initialise goblin d'entraînement
func initGoblin() Monster {
	return Monster{
		Name:   "Gobelin d'entraînement",
		PvMax:  40,
		Pv:     40,
		Attack: 5,
	}
}
