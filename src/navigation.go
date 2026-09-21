package main

import (
	"fmt"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

//-------------// AJOUT BUBBLE TEA (uniquement visuel) //----------------//

type navigationScreen int

const (
	mainScreen navigationScreen = iota
	infoScreen
	inventoryScreen
	combatScreen
	combatInventoryScreen
	combatResultScreen
)

// navigationModel contient uniquement l'etat de l'interface clavier.
// Le personnage et les objets existants restent utilises tels quels.
type navigationModel struct {
	player        *Character
	screen        navigationScreen
	cursor        int
	monster       Monster
	combatTurn    int
	statusMessage string
}

func newNavigationModel(player *Character) navigationModel {
	return navigationModel{
		player: player,
		screen: mainScreen,
	}
}

func (m navigationModel) Init() tea.Cmd {
	return nil
}

// Update traite les fleches, Entree et Echap pour tous les ecrans du jeu.
func (m navigationModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	keyMessage, isKeyMessage := message.(tea.KeyMsg)
	if !isKeyMessage {
		return m, nil
	}

	switch keyMessage.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "up", "k":
		m.moveCursor(-1)
	case "down", "j":
		m.moveCursor(1)
	case "esc":
		return m.goBack()
	case "enter", "space":
		return m.validateChoice()
	}

	return m, nil
}

func (m *navigationModel) moveCursor(direction int) {
	optionCount := m.optionCount()
	if optionCount == 0 {
		return
	}

	m.cursor = (m.cursor + direction + optionCount) % optionCount
}

func (m navigationModel) optionCount() int {
	switch m.screen {
	case mainScreen:
		return 4
	case inventoryScreen, combatInventoryScreen:
		return 2
	case combatScreen:
		return 3
	default:
		return 0
	}
}

func (m navigationModel) goBack() (tea.Model, tea.Cmd) {
	switch m.screen {
	case mainScreen:
		return m, tea.Quit
	case infoScreen, inventoryScreen, combatResultScreen:
		m.screen = mainScreen
	case combatInventoryScreen:
		m.screen = combatScreen
	}

	m.cursor = 0
	m.statusMessage = ""
	return m, nil
}

func (m navigationModel) validateChoice() (tea.Model, tea.Cmd) {
	switch m.screen {
	case mainScreen:
		switch m.cursor {
		case 0:
			m.screen = infoScreen
		case 1:
			m.screen = inventoryScreen
		case 2:
			m.screen = combatScreen
			m.monster = initGoblin()
			m.combatTurn = 1
		case 3:
			return m, tea.Quit
		}
	case infoScreen:
		m.screen = mainScreen
	case inventoryScreen:
		if m.cursor == 0 {
			m.statusMessage = useLifePotionFromNavigation(m.player)
		} else {
			m.screen = mainScreen
		}
	case combatInventoryScreen:
		if m.cursor == 0 {
			m.statusMessage = useLifePotionFromNavigation(m.player)
		} else {
			m.screen = combatScreen
		}
	case combatScreen:
		switch m.cursor {
		case 0:
			m.playCombatTurn()
		case 1:
			m.screen = combatInventoryScreen
		case 2:
			m.screen = mainScreen
		}
	case combatResultScreen:
		m.screen = mainScreen
	}

	m.cursor = 0
	return m, nil
}

// useLifePotionFromNavigation reprend la mecanique de takePot sans ecrire
// directement dans le terminal, qui est gere par Bubble Tea.
func useLifePotionFromNavigation(player *Character) string {
	quantity, exists := player.Inventory[PotionVie]
	if !exists || quantity <= 0 {
		return "Vous n'avez plus de potion de vie."
	}

	player.Pv += 50
	if player.Pv > player.PvMax {
		player.Pv = player.PvMax
	}

	player.Inventory[PotionVie]--
	return fmt.Sprintf("Potion utilisee : %d/%d PV.", player.Pv, player.PvMax)
}

// playCombatTurn conserve les regles du combat existant : 5 degats par attaque
// et une attaque du gobelin doublee tous les trois tours.
func (m *navigationModel) playCombatTurn() {
	const basicAttackDamage = 5

	m.monster.Pv -= basicAttackDamage
	if m.monster.Pv <= 0 {
		m.monster.Pv = 0
		m.statusMessage = "Victoire ! Le Gobelin d'entrainement est vaincu."
		m.screen = combatResultScreen
		return
	}

	damage := m.monster.Attack
	if m.combatTurn%3 == 0 {
		damage *= 2
	}

	m.player.Pv -= damage
	if m.player.Pv <= 0 {
		m.player.Pv = 0
		m.statusMessage = "Defaite : le Gobelin d'entrainement vous a vaincu."
		m.screen = combatResultScreen
		return
	}

	m.statusMessage = fmt.Sprintf(
		"Vous infligez %d degats. %s inflige %d degats.",
		basicAttackDamage,
		m.monster.Name,
		damage,
	)
	m.combatTurn++
}

func (m navigationModel) View() string {
	switch m.screen {
	case infoScreen:
		return m.renderInfo()
	case inventoryScreen:
		return m.renderInventory(false)
	case combatScreen:
		return m.renderCombat()
	case combatInventoryScreen:
		return m.renderInventory(true)
	case combatResultScreen:
		return renderNavigationResult(m.statusMessage)
	default:
		return m.renderMainMenu()
	}
}

func (m navigationModel) renderMainMenu() string {
	var builder strings.Builder

	for _, line := range navigationLogoLines() {
		builder.WriteString(line + "\n")
	}

	builder.WriteString("\n")
	builder.WriteString(centerStyled(Bold+White+" - O F   G O P H E R - "+Reset, " - O F   G O P H E R - ", menuWidth))
	builder.WriteString("\n")
	builder.WriteString(renderNavigationHeading("MENU PRINCIPAL", Magenta))
	builder.WriteString(renderNavigationOption("Informations du personnage", m.cursor == 0) + "\n")
	builder.WriteString(renderNavigationOption("Acceder a l'inventaire", m.cursor == 1) + "\n")
	builder.WriteString(renderNavigationOption("Combat de test", m.cursor == 2) + "\n")
	builder.WriteString(renderNavigationOption("Quitter", m.cursor == 3))
	builder.WriteString(renderNavigationFooter("↑ ↓ naviguer  •  Entree valider  •  Echap quitter"))

	return builder.String()
}

// navigationLogoLines reprend le logo fourni pour l'ecran d'accueil Bubble Tea.
func navigationLogoLines() []string {
	return []string{
		Cyan + "██████╗  ██████╗ ██╗    ██╗███╗   ██╗███████╗ █████╗ ██╗     ██╗" + Reset,
		Cyan + "██╔══██╗██╔═══██╗██║    ██║████╗  ██║██╔════╝██╔══██╗██║     ██║" + Reset,
		Cyan + "██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║█████╗  ███████║██║     ██║" + Reset,
		Cyan + "██║  ██║██║   ██║██║███╗██║██║╚██╗██║██╔══╝  ██╔══██║██║     ██║" + Reset,
		Cyan + "██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║██║     ██║  ██║███████╗███████╗" + Reset,
		Cyan + "╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝" + Reset,
	}
}

func (m navigationModel) renderInfo() string {
	lines := []string{
		fmt.Sprintf("%sNom%s      : %s", Cyan, Reset, m.player.Name),
		fmt.Sprintf("%sClasse%s   : %s", Cyan, Reset, m.player.Classe),
		fmt.Sprintf("%sPV%s       : %d/%d", Cyan, Reset, m.player.Pv, m.player.PvMax),
		fmt.Sprintf("%sArgent%s   : %d", Cyan, Reset, m.player.Money),
	}

	return renderNavigationScreen("FICHE PERSONNAGE", Cyan, lines, "Entree / Echap : retour")
}

func (m navigationModel) renderInventory(fromCombat bool) string {
	itemNames := make([]string, 0, len(m.player.Inventory))
	for itemName := range m.player.Inventory {
		itemNames = append(itemNames, itemName)
	}
	sort.Strings(itemNames)

	lines := make([]string, 0, len(itemNames)+6)
	lines = append(lines, Magenta+"── OBJETS ──"+Reset)
	for _, itemName := range itemNames {
		lines = append(lines, fmt.Sprintf("%s•%s %-22s x%d", Cyan, Reset, itemName, m.player.Inventory[itemName]))
	}

	lines = append(lines, "", Magenta+"── UTILISER ──"+Reset)
	lines = append(lines, renderNavigationOption("Utiliser une potion de vie", m.cursor == 0))
	if fromCombat {
		lines = append(lines, renderNavigationOption("Retour au combat", m.cursor == 1))
	} else {
		lines = append(lines, renderNavigationOption("Retour au menu principal", m.cursor == 1))
	}
	if m.statusMessage != "" {
		lines = append(lines, "", Green+m.statusMessage+Reset)
	}

	return renderNavigationScreen("INVENTAIRE", Magenta, lines, "↑ ↓ naviguer  •  Entree valider  •  Echap retour")
}

func (m navigationModel) renderCombat() string {
	lines := []string{
		fmt.Sprintf("%s%s%s  •  %sPV %d/%d%s", Bold, m.player.Name, Reset, Green, m.player.Pv, m.player.PvMax, Reset),
		fmt.Sprintf("%s%s%s  •  %sPV %d/%d%s", Bold, m.monster.Name, Reset, Red, m.monster.Pv, m.monster.PvMax, Reset),
		"",
		renderNavigationOption("Attaquer", m.cursor == 0),
		renderNavigationOption("Inventaire", m.cursor == 1),
		renderNavigationOption("Abandonner le combat", m.cursor == 2),
	}
	if m.statusMessage != "" {
		lines = append(lines, "", Cyan+m.statusMessage+Reset)
	}

	return renderNavigationScreen("COMBAT // GOBELIN D'ENTRAINEMENT", Red, lines, "↑ ↓ naviguer  •  Entree valider  •  Echap retour")
}

func renderNavigationOption(label string, selected bool) string {
	if selected {
		return Yellow + Bold + "▶ " + label + Reset
	}

	return White + "  " + label + Reset
}

// renderNavigationHeading conserve un titre structure sans encadrer tout
// l'ecran : les contenus restent ouverts et aeres.
func renderNavigationHeading(title string, color string) string {
	label := "✦ " + title + " ✦"
	lineLength := 72 - len([]rune(label))
	if lineLength < 3 {
		lineLength = 3
	}

	return "\n\n" + color + "── " + Bold + label + Reset + color + " " + strings.Repeat("─", lineLength) + Reset + "\n"
}

// renderNavigationScreen garde les ecrans lisibles sans les enfermer dans
// une grande boite : seul le titre est encadre, comme dans les maquettes.
func renderNavigationScreen(title string, titleColor string, lines []string, footer string) string {
	var builder strings.Builder

	builder.WriteString(renderNavigationHeading(title, titleColor))
	builder.WriteString("\n")

	for _, line := range lines {
		builder.WriteString(line + "\n")
	}

	builder.WriteString(renderNavigationFooter(footer))
	return builder.String()
}

func renderNavigationResult(statusMessage string) string {
	title := "COMBAT TERMINE"
	color := Red
	if strings.HasPrefix(statusMessage, "Victoire") {
		title = "VICTOIRE"
		color = Green
	}

	return renderNavigationScreen(title, color, []string{Bold + statusMessage + Reset}, "Entree / Echap : retour au menu")
}

func renderNavigationFooter(text string) string {
	return "\n" + Yellow + Bold + text + Reset + "\n"
}

func startNavigation(player *Character) error {
	program := tea.NewProgram(newNavigationModel(player), tea.WithAltScreen())
	_, err := program.Run()
	return err
}

//-------------// FIN AJOUT BUBBLE TEA //----------------//
