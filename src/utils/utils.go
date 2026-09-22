package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const (
	Reset   = "\033[0m"
	Cyan    = "\033[36m"
	Yellow  = "\033[33m"
	Green   = "\033[32m"
	Red     = "\033[31m"
	Magenta = "\033[35m"
	Bold    = "\033[1m"
)

func ClearScreen() {
	var command *exec.Cmd

	if runtime.GOOS == "windows" {
		command = exec.Command("cmd", "/c", "cls")
	} else {
		command = exec.Command("clear")
	}

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	_ = command.Run()
}

func ReadLine() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ReadInt() int {
	for {
		text := ReadLine()
		number, err := strconv.Atoi(text)

		if err == nil {
			return number
		}

		fmt.Print("Veuillez entrer un nombre : ")
	}
}

func Pause() {
	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	ReadLine()
}

func PrintTitle() {
	fmt.Println(Cyan + Bold)
	fmt.Println("██████╗  ██████╗ ██╗    ██╗███╗   ██╗")
	fmt.Println("██╔══██╗██╔═══██╗██║    ██║████╗  ██║")
	fmt.Println("██║  ██║██║   ██║██║ █╗ ██║██╔██╗ ██║")
	fmt.Println("██║  ██║██║   ██║██║███╗██║██║╚██╗██║")
	fmt.Println("██████╔╝╚██████╔╝╚███╔███╔╝██║ ╚████║")
	fmt.Println("╚═════╝  ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝")
	fmt.Println(Reset)
	fmt.Println("            - OF GOPHER -")
	fmt.Println()
}
