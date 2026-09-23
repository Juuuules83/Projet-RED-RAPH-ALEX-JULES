package main

import (
	"projet-red/player"
	"projet-red/ui"
)

func main() {
	character := player.CharacterCreation()
	ui.MenuPrincipal(&character)
}