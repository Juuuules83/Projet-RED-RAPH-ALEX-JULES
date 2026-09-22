package utils

import "fmt"

// ------- DEBUT UI COMBAT.GO -------//

func (m NavigationModel) RenderCombat() string {
	lines := []string{
		fmt.Sprintf("%s%s%s  •  %sPV %d/%d%s", Bold, m.Player.Name, Reset, Green, m.Player.Pv, m.Player.PvMax, Reset),
		fmt.Sprintf("%s%s%s  •  %sPV %d/%d%s", Bold, m.Monster.Name, Reset, Red, m.Monster.Pv, m.Monster.PvMax, Reset),
		"",
		RenderNavigationOption("Attaquer", m.Cursor == 0),
		RenderNavigationOption("Inventaire", m.Cursor == 1),
		RenderNavigationOption("Abandonner le combat", m.Cursor == 2),
	}

	if m.StatusMessage != "" {
		lines = append(lines, "", Cyan+m.StatusMessage+Reset)
	}

	return RenderNavigationScreen("COMBAT // GOBELIN D'ENTRAINEMENT", Red, lines, "↑ ↓ naviguer  •  Entree valider  •  Echap retour")
}

// ------- FIN UI COMBAT.GO -------//
