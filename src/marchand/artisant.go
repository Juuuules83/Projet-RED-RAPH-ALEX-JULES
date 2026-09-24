package marchand

import (
	"fmt"
	"projet-red/player"
	"projet-red/utils"
)

func Artisant(c *player.Character) {
 
    switch utils.Choose {
    case 0:
        utils.ClearScreen()
        fmt.Println("Retour")
 
        return
 
    case 1:
        if utils.Countfree >= 0{
            utils.Countfree++
            fmt.Println("potion de vie [GRATUIT]")
            c.AddInventory(utils.PotionVie, 1)
        }else if c.Money >= 10 && utils.StockageMax > utils.TotalItems {
            fmt.Println("potion de vie [10]")
            c.AddInventory(utils.PotionPoison, 1)
            c.Money -= 10
        }else{
            fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
        }
 
    case 2:
        if c.Money >= 25 && utils.StockageMax > utils.TotalItems {
            fmt.Println("potion de poison [25]")
            c.AddInventory(utils.PotionPoison, 1)
            c.Money -= 25
        }else{
            fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
        }
    case 3:
        if utils.CountFB > 0{
            if c.Money >= 75 && utils.StockageMax > utils.TotalItems {
                fmt.Println("Sort : boule de feu [75]")
                c.AddInventory(utils.FireBall, 1)
                c.Money -= 75
                utils.CountFB++
            }else{
                fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
            }
        }else{
            fmt.Print("Sort déja acheté")
        }
     case 4:
        if c.Money >= 4 && utils.StockageMax > utils.TotalItems {
            fmt.Println("fourrure de loup [4]")
            c.AddInventory(utils.WolfFurr, 1)
            c.Money -= 4 
        }else{
            fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
        }
     case 5:
        if c.Money >= 7 && utils.StockageMax > utils.TotalItems {
            fmt.Println("peau de troll [7]")
            c.AddInventory(utils.TrollSkin, 1)
            c.Money -= 7
        }else{
            fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
        }
     case 6:
        if c.Money >= 3 && utils.StockageMax > utils.TotalItems {
            fmt.Println("cuir de sanglier [3]")
            c.AddInventory(utils.BoarLeather, 1)
            c.Money -= 3
        }else{
            fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
        }
    case 7:
        if c.Money >= 1 && utils.StockageMax > utils.TotalItems {
            fmt.Println("plume de corbeau [1]")
            c.AddInventory(utils.RavenFeather, 1)
            c.Money -= 1
        }else{
            fmt.Println("Pas assez d'argent ou d'espace dans l'inventaire")
        }
    default:
        utils.ClearScreen()
        fmt.Println("Option invalide. Veuillez réessayer.")
    }
}