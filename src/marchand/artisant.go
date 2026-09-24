    package marchand

    import (
        "fmt"
        "projet-red/player"
        "projet-red/utils"
    )

func Artisant(c *player.Character) {
    for {
        utils.ClearScreen()

        fmt.Println("===== MARCHAND =====")
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
            if c.Money >= 25 {
                if c.AddInventory(utils.PotionPoison, 1) {
                    fmt.Println("Potion de poison achetée [25 €]")
                    c.Money -= 25
                }
            } else {
                fmt.Println("ESPECE DE GUEUX VAAAA BOUUUUH !")
            }

        case 2:
            if c.Money >= 25 {
                if c.AddInventory(utils.PotionPoison, 1) {
                    fmt.Println("Potion de poison achetée [25 €]")
                    c.Money -= 25
                }
            } else {
                fmt.Println("ESPECE DE GUEUX VAAAA BOUUUUH !")
            }

        case 3:
            if c.Inventory[utils.FireBall] == 0 {
                if c.Money >= 75 {
                    if c.AddInventory(utils.FireBall, 1) {
                        fmt.Println("Sort : boule de feu acheté [75 €]")
                        c.Money -= 75
                    }
                } else {
                    fmt.Println("AHHHH T'AS PAS D'ARGENT TU M'DEGOUTES !")
                }
            } else {
                fmt.Println("Sort déjà acheté !")
            }
        }
    }
}