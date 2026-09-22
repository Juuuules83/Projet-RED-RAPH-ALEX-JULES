package utils

import (
	"fmt"
	"strings"
)

// ------- DEBUT UI GLOBAL -------//

func RenderNavigationOption(label string, selected bool) string {
	if selected {
		return Yellow + Bold + "▶ " + label + Reset
	}
	return White + "  " + label + Reset
}

func RenderNavigationHeading(title string, color string) string {
	label := "✦ " + title + " ✦"
	lineLength := 72 - len([]rune(label))
	if lineLength < 3 {
		lineLength = 3
	}
	return "\n\n" + color + "── " + Bold + label + Reset + color + " " + strings.Repeat("─", lineLength) + Reset + "\n"
}

func RenderNavigationScreen(title string, titleColor string, lines []string, footer string) string {
	var builder strings.Builder
	builder.WriteString(RenderNavigationHeading(title, titleColor))
	builder.WriteString("\n")
	for _, line := range lines {
		builder.WriteString(line + "\n")
	}
	builder.WriteString(RenderNavigationFooter(footer))
	return builder.String()
}

func RenderNavigationFooter(text string) string {
	return "\n" + Yellow + Bold + text + Reset + "\n"
}

func RenderNavigationResult(statusMessage string) string {
	title := "COMBAT TERMINE"
	color := Red
	if strings.HasPrefix(statusMessage, "Victoire") {
		title = "VICTOIRE"
		color = Green
	}
	return RenderNavigationScreen(title, color, []string{Bold + statusMessage + Reset}, "Entree / Echap : retour au menu")
}

func RenderGameOver(playerName string) string {
	lines := []string{
		Red + Bold + "GAME OVER" + Reset,
		"",
		fmt.Sprintf("%s est mort au combat.", playerName),
		"",
		Red + "Le jeu va se fermer automatiquement..." + Reset,
	}
	return RenderNavigationScreen("GAME OVER", Red, lines, "Fermeture automatique dans 3 secondes...")
}

// ------- FIN UI GLOBAL -------//
