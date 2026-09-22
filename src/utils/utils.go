package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

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

	MenuWidth = 76
)

func CenterPlain(text string, width int) string {
	padding := width - len([]rune(text))
	if padding <= 0 {
		return text
	}
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}

func CenterStyled(styledText, plainText string, width int) string {
	padding := width - len([]rune(plainText))
	if padding <= 0 {
		return styledText
	}
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + styledText + strings.Repeat(" ", right)
}

func StripANSI(text string) string {
	for _, code := range []string{Reset, Red, Green, Yellow, Blue, Magenta, Cyan, White, Bold, Italic, Underline} {
		text = strings.ReplaceAll(text, code, "")
	}
	return text
}

func PrintGameTitle() string {
	logo := []string{
		"██████╗  ██████╗ ██╗    ██╗███╗   ██╗███████╗ █████╗ ██╗     ██╗     ",
		"██╔══██╗██╔═══██╗██║    ██║████╗  ██║██╔════╝██╔══██╗██║     ██║     ",
		"██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║█████╗  ███████║██║     ██║     ",
		"██║  ██║██║   ██║██║███╗██║██║╚██╗██║██╔══╝  ██╔══██║██║     ██║     ",
		"██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║██║     ██║  ██║███████╗███████╗",
		"╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝",
	}

	var builder strings.Builder
	for _, line := range logo {
		builder.WriteString(Cyan + CenterPlain(line, MenuWidth) + Reset + "\n")
	}
	return builder.String()
}

func PrintBoxTitle(title string) string {
	border := "╔" + strings.Repeat("═", MenuWidth) + "╗"
	return Cyan + border + Reset + "\n" +
		Cyan + "║" + Reset + CenterStyled(Bold+Magenta+title+Reset, title, MenuWidth) + Cyan + "║" + Reset + "\n" +
		Cyan + "╠" + strings.Repeat("═", MenuWidth) + "╣" + Reset
}

func PrintBoxLine(text string) string {
	visibleLength := len([]rune(StripANSI(text)))
	padding := MenuWidth - 4 - visibleLength
	if padding < 0 {
		padding = 0
	}
	return Cyan + "║" + Reset + "  " + text + strings.Repeat(" ", padding) + "  " + Cyan + "║" + Reset
}

func PrintBoxSeparator() string {
	return Cyan + "╠" + strings.Repeat("─", MenuWidth) + "╣" + Reset
}

func PrintBoxBottom() string {
	return Cyan + "╚" + strings.Repeat("═", MenuWidth) + "╝" + Reset
}

func ClearScreen() {
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

func ExitWithError(err error) {
	if err == nil {
		return
	}
	fmt.Println(Red+"Erreur lors du lancement du jeu :"+Reset, err)
}
