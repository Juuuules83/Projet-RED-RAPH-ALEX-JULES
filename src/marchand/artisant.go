package marchand

import (
	"fmt"
)

func Artisant() {
  for {
        utils.ClearScreen()

        fmt.Println("===== ARTISANT =====")
        fmt.Printf("Argent : %d €\n", c.Money)
        fmt.Printf("Inventaire : %d/%d\n\n",
            c.TotalItems(), utils.StockageMax)

        fmt.Println("1. Chapeau de l'aventurier \n \t Objets Necéssaire : x1 Plume de corbeau, x1 Cuir de sanglier")
        fmt.Println("2. Tunique de l'aventurier \n \t Objets Necéssaire : x2 Fourrure de loup, x1 Peau de Troll")
        fmt.Println("3. Bottes de l'aventurier \n \t Objets Necéssaire : x1 Fourrure de loup, x1 Cuir de sanglier")
        fmt.Println("0. Retour")

        fmt.Print("\nVotre choix : ")
        choice := utils.ReadInt()

        switch choice {
        case 0:
            utils.ClearScreen()
            fmt.Println("Retour")
            return

        case 1:
           

        case 2:
           
        case 3:
            
        }
    }
}
