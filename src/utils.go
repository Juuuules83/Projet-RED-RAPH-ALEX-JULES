package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// --------------- COULEURS ET STYLES ---------------

// Codes ANSI pour les couleurs
const (
	Reset     = "\033[0m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	White     = "\033[37m"
	Bold      = "\033[1m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
)

// --------------- DIMENSIONS DE L'INTERFACE ---------------

// menuWidth est la largeur intérieure des grands cadres.
const menuWidth = 76

// --------------- TITRE DU JEU ---------------

// printGameTitle affiche le logo ASCII de Rise of Gopher.
// Le titre est volontairement grand : il donne immédiatement une identité
// visuelle au jeu avant même que le joueur arrive au menu.
// --------------- TITRE DU JEU ---------------

func printGameTitle() {
	border := "╔" + strings.Repeat("═", menuWidth) + "╗"

	fmt.Println(Cyan + border + Reset)
	fmt.Println(Cyan + "║" + Reset + strings.Repeat(" ", menuWidth) + Cyan + "║" + Reset)

	logo := []string{
		"██████╗  ██████╗ ██╗    ██╗███╗   ██╗███████╗ █████╗ ██╗     ██╗     ",
		"██╔══██╗██╔═══██╗██║    ██║████╗  ██║██╔════╝██╔══██╗██║     ██║     ",
		"██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║█████╗  ███████║██║     ██║     ",
		"██║  ██║██║   ██║██║███╗██║██║╚██╗██║██╔══╝  ██╔══██║██║     ██║     ",
		"██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║██║     ██║  ██║███████╗███████╗",
		"╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝",                                                  	
	}

	for _, line := range logo {
		fmt.Println(
			Cyan+"║"+Reset+
				centerPlain(line, menuWidth)+
				Cyan+"║"+Reset,
		)
	}

	fmt.Println(Cyan + "║" + Reset + strings.Repeat(" ", menuWidth) + Cyan + "║" + Reset)

	subtitle := Bold + Magenta + "O F   G O P H E R" + Reset

	fmt.Println(
		Cyan+"║"+Reset+
			centerStyled(subtitle, "O F   G O P H E R", menuWidth)+
			Cyan+"║"+Reset,
	)

	fmt.Println(Cyan + "║" + Reset + strings.Repeat(" ", menuWidth) + Cyan + "║" + Reset)
	fmt.Println(Cyan + "╠" + strings.Repeat("═", menuWidth) + "╣" + Reset)
}

// centerPlain centre un texte qui ne contient pas de codes ANSI.
func centerPlain(text string, width int) string {
	padding := width - len([]rune(text))
	if padding <= 0 {
		return text
	}

	left := padding / 2
	right := padding - left

	return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}

// centerStyled centre un texte coloré en utilisant sa version sans ANSI
// pour calculer correctement les espaces. Cela évite que les codes de couleur
// faussent la largeur du cadre.
func centerStyled(styledText, plainText string, width int) string {
	padding := width - len([]rune(plainText))
	if padding <= 0 {
		return styledText
	}

	left := padding / 2
	right := padding - left

	return strings.Repeat(" ", left) + styledText + strings.Repeat(" ", right)
}

// --------------- CADRES ---------------

// printBoxTitle affiche un petit titre centré dans un cadre.
// Il est utilisé pour les sous-écrans comme l'inventaire ou les erreurs.
func printBoxTitle(title string) {
	border := "╔" + strings.Repeat("═", menuWidth) + "╗"

	fmt.Println(Cyan + border + Reset)
	fmt.Println(Cyan + "║" + Reset + centerStyled(Bold+Magenta+title+Reset, title, menuWidth) + Cyan + "║" + Reset)
	fmt.Println(Cyan + "╠" + strings.Repeat("═", menuWidth) + "╣" + Reset)
}

// printBoxLine affiche une ligne de contenu dans le cadre principal.
func printBoxLine(text string) {
	// Le cadre reste volontairement simple : le contenu est aligné à gauche.
	// Les couleurs ANSI ne sont pas comptées dans la largeur visuelle.
	visibleLength := len([]rune(stripANSI(text)))
	padding := menuWidth - 4 - visibleLength
	if padding < 0 {
		padding = 0
	}

	fmt.Println(Cyan + "║" + Reset + "  " + text + strings.Repeat(" ", padding) + "  " + Cyan + "║" + Reset)
}

// printBoxSeparator sépare les différentes parties d'un menu.
func printBoxSeparator() {
	fmt.Println(Cyan + "╠" + strings.Repeat("─", menuWidth) + "╣" + Reset)
}

// printBoxBottom ferme un cadre.
func printBoxBottom() {
	fmt.Println(Cyan + "╚" + strings.Repeat("═", menuWidth) + "╝" + Reset)
}

// stripANSI retire les codes de couleur pour calculer la longueur visible.
func stripANSI(text string) string {
	for _, code := range []string{
		Reset, Red, Green, Yellow, Blue, Magenta, Cyan, White,
		Bold, Italic, Underline,
	} {
		text = strings.ReplaceAll(text, code, "")
	}

	return text
}

// --------------- OUTILS TERMINAL ---------------

// clearScreen nettoie le terminal.
func clearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

// waitForEnter attend que le joueur appuie sur Entrée.
func waitForEnter() {
	fmt.Println()
	fmt.Print(Cyan + "  Appuyez sur Entrée pour continuer..." + Reset)
	fmt.Scanln()
	fmt.Scanln()
}

// typeWriter affiche un texte caractère par caractère.
func typeWriter(str string, delay time.Duration) {
	for _, r := range str {
		fmt.Printf("%c", r)
		time.Sleep(delay)
	}
	fmt.Println()
}
