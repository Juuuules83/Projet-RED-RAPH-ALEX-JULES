package utils

import (
	"fmt"
	"sort"
)

// ------- DEBUT UI INVENTORY.GO -------//

func (m NavigationModel) RenderInventory(fromCombat bool) string {
	itemNames := make([]string, 0, len(m.Player.Inventory))
	for itemName, quantity := range m.Player.Inventory {
		if quantity > 0 {
			itemNames = append(itemNames, itemName)
		}
	}
	sort.Strings(itemNames)

	lines := []string{
		fmt.Sprintf("%sObjets :%s %d/%d", Cyan, Reset, m.Player.TotalItems(), 10),
		"",
		Magenta + "── OBJETS ──" + Reset,
	}

	if len(itemNames) == 0 {
		lines = append(lines, "Aucun objet dans l'inventaire.")
	} else {
		for _, itemName := range itemNames {
			lines = append(lines, fmt.Sprintf("%s•%s %-22s x%d", Cyan, Reset, itemName, m.Player.Inventory[itemName]))
		}
	}

	lines = append(lines, "", Magenta+"── UTILISER ──"+Reset)
	lines = append(lines, RenderNavigationOption("Utiliser une potion de vie", m.Cursor == 0))

	if fromCombat {
		lines = append(lines, RenderNavigationOption("Retour au combat", m.Cursor == 1))
	} else {
		lines = append(lines, RenderNavigationOption("Retour au menu principal", m.Cursor == 1))
	}

	if m.StatusMessage != "" {
		lines = append(lines, "", Green+m.StatusMessage+Reset)
	}

	return RenderNavigationScreen("INVENTAIRE", Magenta, lines, "↑ ↓ naviguer  •  Entree valider  •  Echap retour")
}

// ------- FIN UI INVENTORY.GO -------//
