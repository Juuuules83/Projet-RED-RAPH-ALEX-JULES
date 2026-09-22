package combat

type Monster struct {
	Name   string
	PvMax  int
	Pv     int
	Attack int
}

func InitGoblin() Monster {
	return Monster{
		Name:   "Gobelin d'entraînement",
		PvMax:  40,
		Pv:     40,
		Attack: 5,
	}
}
