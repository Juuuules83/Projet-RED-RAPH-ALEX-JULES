package marchand

import (
    "fmt"
    "projet-red/player"
    "projet-red/utils"
)

func Marchand(c *player.Character) {
    for {
        utils.ClearScreen()

        fmt.Println("===== YMATCH =====")
        fmt.Printf("Argent : %d €\n", c.Money)
        fmt.Printf("Inventaire : %d/%d\n\n",
            c.TotalItems(), utils.StockageMax)

        fmt.Println(utils.Bold + "1. Café du dev : 10 €" + utils.Reset)
        fmt.Println(utils.Bold + "2. Soupe de doute : 25 €" + utils.Reset)
        fmt.Println(utils.Bold + "3. Exploit de faille : 75 €" + utils.Reset)
        fmt.Println(utils.Bold + "4. Clé SSH : 4 €" + utils.Reset)
        fmt.Println(utils.Bold + "5. Carte graphique : 7 €" + utils.Reset)
        fmt.Println(utils.Bold + "6. Fragment de code : 3 €" + utils.Reset)
        fmt.Println(utils.Bold + "7. Ticket Ytrack : 1 €" + utils.Reset)
        fmt.Println(utils.Bold + utils.Red + "0. Retour" + utils.Reset)

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
                    fmt.Println("Café du dev [GRATUIT]")
                }
            } else if c.Money >= 10 {
                if c.AddInventory(utils.PotionVie, 1) {
                    fmt.Println("Café du dev acheté [10 €]")
                    c.Money -= 10
                }
            } else {
                fmt.Println("T'as pas d'argent !")
            }

        case 2:
            if c.Money >= 25 {
                if c.AddInventory(utils.PotionPoison, 1) {
                    fmt.Println("Soupe de doute achetée [25 €]")
                    c.Money -= 25
                }
            } else {
                fmt.Println("T'as pas d'argent !")
            }

        case 3:
            if c.Inventory[utils.FireBall] == 0 {
                if c.Money >= 75 {
                    if c.AddInventory(utils.FireBall, 1) {
                        fmt.Println("Exploit de faille acheté [75 €]")
                        c.Money -= 75
                    }
                } else {
                    fmt.Println("T'as pas d'argent !")
                }
            } else {
                fmt.Println("Sort déjà acheté !")
            }

        case 4:
            if c.Money >= 4 {
                if c.AddInventory(utils.WolfFurr, 1) {
                    fmt.Println("Clé SSH achetée [4 €]")
                    c.Money -= 4
                }
            } else {
                fmt.Println("T'as pas d'argent !")
            }

        case 5:
            if c.Money >= 7 {
                if c.AddInventory(utils.TrollSkin, 1) {
                    fmt.Println("Carte graphique achetée [7 €]")
                    c.Money -= 7
                }
            } else {
                fmt.Println("T'as pas d'argent !")
            }

        case 6:
            if c.Money >= 3 {
                if c.AddInventory(utils.BoarLeather, 1) {
                    fmt.Println("Fragment de code acheté [3 €]")
                    c.Money -= 3
                }
            } else {
                fmt.Println("T'as pas d'argent !")
            }

        case 7:
            if c.Money >= 1 {
                if c.AddInventory(utils.RavenFeather, 1) {
                    fmt.Println("Ticket Ytrack acheté [1 €]")
                    c.Money -= 1
                }
            } else {
                fmt.Println("T'as pas d'argent !")
            }

        default:
            fmt.Println("Option invalide. Veuillez réessayer.")
        }

        utils.Pause()
    }
}

