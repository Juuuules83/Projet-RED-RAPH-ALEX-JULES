    package marchand

    import (
        "fmt"
        "projet-red/player"
        "projet-red/utils"
    )

func Marchand(c *player.Character) {
    for {
        utils.ClearScreen()

        fmt.Println("===== MARCHAND =====")
        fmt.Printf("Argent : %d €\n", c.Money)
        fmt.Printf("Inventaire : %d/%d\n\n",
            c.TotalItems(), utils.StockageMax)

        fmt.Println("1. Potion de vie : 10 €")
        fmt.Println("2. Potion de poison : 25 €")
        fmt.Println("3. Boule de feu : 75 €")
        fmt.Println("4. Fourrure de loup : 4 €")
        fmt.Println("5. Peau de troll : 7 €")
        fmt.Println("6. Cuir de sanglier : 3 €")
        fmt.Println("7. Plume de corbeau : 1 €")
        fmt.Println("0. Retour")

        fmt.Print("\nVotre choix : ")
        choice := utils.ReadInt()

        switch choice {
        case 0:
            utils.ClearScreen()
            fmt.Println("Retour")
            return

        case 1:
            if utils.Countfree == 0 {
                if c.AddInventory(utils.PotionVie, 1) {
                    utils.Countfree++
                    fmt.Println("Potion de vie [GRATUIT]")
                }
            } else if c.Money >= 10 {
                if c.AddInventory(utils.PotionVie, 1) {
                    fmt.Println("Potion de vie achetée [10 €]")
                    c.Money -= 10
                }
            } else {
                fmt.Println("T'es PAUVRE... PAUUUUVRE !!!!")
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

        case 4:
            if c.Money >= 4 {
                if c.AddInventory(utils.WolfFurr, 1) {
                    fmt.Println("Fourrure de loup achetée [4 €]")
                    c.Money -= 4
                }
            } else {
                fmt.Println("VA BOSSER ET REVIENS QUAND TU POURRAS TE LE PERMETTRE!")
            }

        case 5:
            if c.Money >= 7 {
                if c.AddInventory(utils.TrollSkin, 1) {
                    fmt.Println("Peau de troll achetée [7 €]")
                    c.Money -= 7
                }
            } else {
                fmt.Println("VA ET FAIRE VOIR !")
            }

        case 6:
            if c.Money >= 3 {
                if c.AddInventory(utils.BoarLeather, 1) {
                    fmt.Println("Cuir de sanglier acheté [3 €]")
                    c.Money -= 3
                }
            } else {
                fmt.Println(" .... SANS COMMENTAIRE SALE PAUVRE !")
            }

        case 7:
            if c.Money >= 1 {
                if c.AddInventory(utils.RavenFeather, 1) {
                    fmt.Println("Plume de corbeau achetée [1 €]")
                    c.Money -= 1
                }
            } else {
                fmt.Println("T'es PAUVRE... PAUUUUVRE !!!!")
            }

        default:
            fmt.Println("Option invalide. Veuillez réessayer.")
        }

        // Attend que le joueur appuie sur Entrée
        // avant de réafficher le marchand.
        utils.Pause()
    }
}
