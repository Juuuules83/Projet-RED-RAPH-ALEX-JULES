package utils

import (
	"fmt"
	"time"

	"projet-red/combat"
	"projet-red/player"
	"projet-red/trader"

	tea "github.com/charmbracelet/bubbletea"
)

// ------- DEBUT NAVIGATION -------//

type navigationScreen int

const (
	mainScreen navigationScreen = iota
	infoScreen
	inventoryScreen
	merchantScreen
	combatScreen
	combatInventoryScreen
	combatResultScreen
	gameOverScreen
)

// NavigationModel contient l'état de l'interface et du combat.
type NavigationModel struct {
	Player        *player.Character
	Screen        navigationScreen
	Cursor        int
	Monster       combat.Monster
	CombatTurn    int
	StatusMessage string
}

type gameOverMessage struct{}

func NewNavigationModel(character *player.Character) NavigationModel {
	return NavigationModel{
		Player: character,
		Screen: mainScreen,
	}
}

func (m NavigationModel) Init() tea.Cmd {
	return nil
}

func (m NavigationModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := message.(gameOverMessage); ok {
		return m, tea.Quit
	}

	keyMessage, ok := message.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch keyMessage.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "up", "k":
		m.MoveCursor(-1)
	case "down", "j":
		m.MoveCursor(1)
	case "esc":
		return m.GoBack()
	case "enter", "space":
		return m.ValidateChoice()
	}

	return m, nil
}

func (m *NavigationModel) MoveCursor(direction int) {
	count := m.OptionCount()
	if count == 0 {
		return
	}
	m.Cursor = (m.Cursor + direction + count) % count
}

func (m NavigationModel) OptionCount() int {
	switch m.Screen {
	case mainScreen:
		return 5
	case inventoryScreen, combatInventoryScreen:
		return 2
	case merchantScreen:
		return 2
	case combatScreen:
		return 3
	default:
		return 0
	}
}

func (m NavigationModel) GoBack() (tea.Model, tea.Cmd) {
	switch m.Screen {
	case mainScreen:
		return m, tea.Quit
	case infoScreen, inventoryScreen, merchantScreen, combatResultScreen:
		m.Screen = mainScreen
	case combatScreen:
		m.Screen = mainScreen
	case combatInventoryScreen:
		m.Screen = combatScreen
	case gameOverScreen:
		return m, tea.Quit
	}

	m.Cursor = 0
	m.StatusMessage = ""
	return m, nil
}

func (m NavigationModel) ValidateChoice() (tea.Model, tea.Cmd) {
	switch m.Screen {
	case mainScreen:
		switch m.Cursor {
		case 0:
			m.Screen = infoScreen
		case 1:
			m.Screen = inventoryScreen
		case 2:
			m.Screen = merchantScreen
		case 3:
			m.Screen = combatScreen
			m.Monster = combat.InitGoblin()
			m.CombatTurn = 1
			m.StatusMessage = ""
		case 4:
			return m, tea.Quit
		}

	case infoScreen:
		m.Screen = mainScreen

	case inventoryScreen:
		if m.Cursor == 0 {
			m.StatusMessage = m.Player.UseLifePotion()
		} else {
			m.Screen = mainScreen
		}

	case merchantScreen:
		if m.Cursor == 0 {
			m.StatusMessage = trader.BuyLifePotion(m.Player)
		} else {
			m.Screen = mainScreen
		}

	case combatInventoryScreen:
		if m.Cursor == 0 {
			m.StatusMessage = combat.UseLifePotion(m.Player)
		} else {
			m.Screen = combatScreen
		}

	case combatScreen:
		switch m.Cursor {
		case 0:
			return m, m.PlayCombatTurn()
		case 1:
			m.Screen = combatInventoryScreen
		case 2:
			m.Screen = mainScreen
		}

	case combatResultScreen:
		m.Screen = mainScreen
	}

	m.Cursor = 0
	return m, nil
}

func (m *NavigationModel) PlayCombatTurn() tea.Cmd {
	playerDamage := combat.PlayerAttack(m.Player, &m.Monster)
	if m.Monster.Pv <= 0 {
		m.StatusMessage = fmt.Sprintf("Victoire ! Vous infligez %d dégâts. Le Gobelin est vaincu.", playerDamage)
		m.Screen = combatResultScreen
		return nil
	}

	monsterDamage := combat.MonsterAttack(m.Player, &m.Monster, m.CombatTurn)
	if combat.IsDead(m.Player.Pv) {
		m.Player.Pv = 0
		m.StatusMessage = "Vous êtes mort. Le Gobelin d'entraînement vous a vaincu."
		m.Screen = gameOverScreen
		return tea.Tick(3*time.Second, func(time.Time) tea.Msg {
			return gameOverMessage{}
		})
	}

	m.StatusMessage = combat.CombatStatus(m.Player, &m.Monster, playerDamage, monsterDamage)
	m.CombatTurn++
	return nil
}

func (m NavigationModel) View() string {
	switch m.Screen {
	case infoScreen:
		return m.RenderInfo()
	case inventoryScreen:
		return m.RenderInventory(false)
	case merchantScreen:
		return m.RenderMerchant()
	case combatScreen:
		return m.RenderCombat()
	case combatInventoryScreen:
		return m.RenderInventory(true)
	case combatResultScreen:
		return RenderNavigationResult(m.StatusMessage)
	case gameOverScreen:
		return RenderGameOver(m.Player.Name)
	default:
		return m.RenderMainMenu()
	}
}

// StartNavigation lance le moteur Bubble Tea du jeu.
func StartNavigation(character *player.Character) error {
	program := tea.NewProgram(NewNavigationModel(character), tea.WithAltScreen())
	_, err := program.Run()
	return err
}

// ------- FIN NAVIGATION -------//
