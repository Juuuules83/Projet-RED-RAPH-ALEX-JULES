package main

import (
	"projet-red/src/player"
	"projet-red/src/ui"
)

func main() {
	character := player.CharacterCreation()
	ui.MenuPrincipal(&character)
}