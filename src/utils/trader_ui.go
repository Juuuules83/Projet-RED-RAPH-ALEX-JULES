package utils

import (
	"fmt"

	"projet-red/player"
	"projet-red/trader"
)

// ------- DEBUT UI TRADER.GO -------//

func (m NavigationModel) RenderMerchant() string {
	lines := []string{
		fmt.Sprintf("%sArgent actuel :%s %s%d €%s", Cyan, Reset, Green, m.Player.Money, Reset),
		fmt.Sprintf("%sInventaire :%s %d/%d", Cyan, Reset, m.Player.TotalItems(), player.StockageMax),
		"",
		Green + "── OBJETS DISPONIBLES ──" + Reset,
		RenderNavigationOption("Potion de vie  •  10 €", m.Cursor == 0),
		RenderNavigationOption("Retour", m.Cursor == 1),
	}

	if m.StatusMessage != "" {
		lines = append(lines, "", Yellow+m.StatusMessage+Reset)
	}

	_ = trader.PotionViePrice
	return RenderNavigationScreen(
		"MARCHAND",
		Green,
		lines,
		"↑ ↓ naviguer  •  Entree acheter l'objet selectionne  •  Echap retour",
	)
}

// ------- FIN UI TRADER.GO -------//
