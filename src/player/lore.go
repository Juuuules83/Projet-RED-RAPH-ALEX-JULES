package player

import (
	"fmt"
	"projet-red/utils"
	"time"
)


func dialogueLore(message string) {
    utils.TypeWriter(message, 28*time.Millisecond)
    time.Sleep(1100 * time.Millisecond)
}

func Lore(name string) {
	utils.ClearScreen()

	fmt.Println(utils.Magenta + "╔══════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                             ★ CAMPUS YNOV ★                         ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════╝" + utils.Reset)
	fmt.Println()
	dialogueLore(utils.Magenta + utils.Bold + "Système " + utils.Reset + "Bienvenue à Ynov. Une anomalie vient de traverser les espaces pédagogiques de Ytrack.")
	fmt.Println()
	dialogueLore(utils.Cyan + utils.Bold + "Vito " + utils.Reset + "Le système a détecté une entité inconnue. Elle réécrit des données qui ne devraient jamais bouger.")
	fmt.Println()
	dialogueLore(utils.Magenta + utils.Bold + "Système " + utils.Reset + "Cette entité est désormais identifiée : Gopher. Son origine reste inconnue.")
	fmt.Println()
	dialogueLore(fmt.Sprintf("%s%sVito %s%s, nous avons besoin de toi pour atteindre le noyau Ytrack.", utils.Bold, utils.Cyan, utils.Reset, name))
	fmt.Println()
	dialogueLore(utils.Yellow + utils.Bold + "Cyril " + utils.Reset + "Et si vous me permettez une petite explication... enfin, une petite... qui pourrait durer un moment...")
	fmt.Println()
	dialogueLore(utils.Green + utils.Bold + "Lilian " + utils.Reset + "On a surtout besoin de quelqu'un capable de garder un code propre au milieu du chaos.")
	fmt.Println()
	dialogueLore(fmt.Sprintf("%s%sVito %s%s, ta spécialisation sera ta première arme. Choisis celle qui correspond à ta façon de résoudre les problèmes.", utils.Bold, utils.Cyan, utils.Reset, name))
}
