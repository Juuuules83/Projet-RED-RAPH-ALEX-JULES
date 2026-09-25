package combat

type Monster struct {
	Name   string
	PvMax  int
	Pv     int
	Attack int
}

func InitGoblin() Monster {
	return Monster{
		Name:   "Gopher",
		PvMax:  125,
		Pv:     125,
		Attack: 15,
	}
}
