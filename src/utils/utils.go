package utils

import (
	"bufio" // Lire entrée clavier
	"fmt" // Afficher texte
	"os" // permet d'écrire dans le CMD
	"os/exec" // Permet d'exécuter des commadandes CMD
	"strconv" // Convertir string en int
	"strings" // écrire texte
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
    command := exec.Command("cmd", "/c", "cls")
	command.Stdout = os.Stdout  
	// command.Stdout = ce que la commande veut afficher
	//os.Stdout = le terminal de mon programme
    
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

   var TotalItems int
   var Choose int

   const (
	PotionVie    = "potion de vie"
	PotionPoison = "potion de poison"
	StockageMax  = 10
)

 var Countfree int  // mise hors de la fonction sinon ça reprend toujours à 0 lorsque la boucle est recommencé
 var CountFB int // mise hors de la fonction sinon ça reprend toujours à 0 lorsque la boucle est recommencé

// les Countfree et CountFB sont raccordé au marchand