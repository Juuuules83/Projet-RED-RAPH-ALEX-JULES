package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Codes ANSI pour les couleurs et les styles du terminal.
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
	fmt.Println("Appuyez sur Entrée pour continuer...")
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
