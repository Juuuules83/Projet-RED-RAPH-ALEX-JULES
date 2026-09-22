package utils

import (
	"fmt"
	"strings"
)

// ------- DEBUT UI MENU.GO -------//

func (m NavigationModel) RenderMainMenu() string {
	var builder strings.Builder

	logo := []string{
		Cyan + "██████╗  ██████╗ ██╗    ██╗███╗   ██╗███████╗ █████╗ ██╗     ██╗" + Reset,
		Cyan + "██╔══██╗██╔═══██╗██║    ██║████╗  ██║██╔════╝██╔══██╗██║     ██║" + Reset,
		Cyan + "██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║█████╗  ███████║██║     ██║" + Reset,
		Cyan + "██║  ██║██║   ██║██║███╗██║██║╚██╗██║██╔══╝  ██╔══██║██║     ██║" + Reset,
		Cyan + "██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║██║     ██║  ██║███████╗███████╗" + Reset,
		Cyan + "╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝" + Reset,
	}

	for _, line := range logo {
		builder.WriteString(CenterPlain(StripANSI(line), MenuWidth))
		builder.WriteString("\n")
	}

	builder.WriteString("\n")
	builder.WriteString(CenterStyled(Bold+White+" - O F   G O P H E R - "+Reset, " - O F   G O P H E R - ", MenuWidth))
	builder.WriteString(RenderNavigationHeading("MENU PRINCIPAL", Magenta))
	builder.WriteString(RenderNavigationOption("Informations du personnage", m.Cursor == 0) + "\n")
	builder.WriteString(RenderNavigationOption("Acceder a l'inventaire", m.Cursor == 1) + "\n")
	builder.WriteString(RenderNavigationOption("Marchand", m.Cursor == 2) + "\n")
	builder.WriteString(RenderNavigationOption("Combat de test", m.Cursor == 3) + "\n")
	builder.WriteString(RenderNavigationOption("Quitter", m.Cursor == 4))
	builder.WriteString(RenderNavigationFooter("↑ ↓ naviguer  •  Entree valider  •  Echap quitter"))

	return builder.String()
}

func (m NavigationModel) RenderInfo() string {
	lines := []string{
		fmt.Sprintf("%sNom%s      : %s", Cyan, Reset, m.Player.Name),
		fmt.Sprintf("%sClasse%s   : %s", Cyan, Reset, m.Player.Classe),
		fmt.Sprintf("%sPV%s       : %d/%d", Cyan, Reset, m.Player.Pv, m.Player.PvMax),
		fmt.Sprintf("%sNiveau%s   : %d", Cyan, Reset, m.Player.Niveau),
		fmt.Sprintf("%sArgent%s   : %d €", Cyan, Reset, m.Player.Money),
	}
	return RenderNavigationScreen("FICHE PERSONNAGE", Cyan, lines, "Entree / Echap : retour")
}

// ------- FIN UI MENU.GO -------//
