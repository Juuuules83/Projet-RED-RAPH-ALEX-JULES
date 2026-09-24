package marchand

import (
	"fmt"
	"projet-red/player"
	"projet-red/utils"
)

const PrixPotionVie = 10

func AcheterPotionVie(c *player.Character) string {
    if c.Money < PrixPotionVie {
        return "trop pauvre espèce de GUEUX !!! BOUUUUH LA LA LA LA LA LA"
    }

    if !c.AddInventory(utils.PotionVie, 1) {
        return "Achat impossible."
    }

    c.Money -= PrixPotionVie

    return "Potion achetée !"
}

func Marchand(c *player.Character) {
 
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
            c.AddInventory(utils.PotionVie, 1)
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


/* func Marchand(c *player.Character) {
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

        var item string
        var prix int

        switch choice {
        case 0:
            return

        case 1:
            item = utils.PotionVie
            prix = 10

            // la prmeière potion de vie est gratuite...
            if utils.Countfree == 0 {
                prix = 0
            }

        case 2:
            item = utils.PotionPoison
            prix = 25

        case 3:
            item = utils.FireBall
            prix = 75

            if c.Inventory[item] > 0 {
                fmt.Println("Sort déjà acheté !")
                utils.Pause()
                continue
            }

        case 4:
            item = utils.WolfFurr
            prix = 4

        case 5:
            item = utils.TrollSkin
            prix = 7

        case 6:
            item = utils.WildBoarLeather
            prix = 3

        case 7:
            item = utils.CrowFeather
            prix = 1

        default:
            fmt.Println("Choix invalide.")
            utils.Pause()
            continue
        }

        if c.Money < prix {
            fmt.Println("t'es pauvre clochard !")
            utils.Pause()
            continue
        }

        // verif de la place dans l'inventaire
        if !c.AddInventory(item, 1) {
            fmt.Println("Achat impossible.")
            utils.Pause()
            continue
        }

        // l'argent est retiré si l'objet a bien été ajouté à l'inventaire !
        c.Money -= prix

        if choice == 1 && utils.Countfree == 0 {
            utils.Countfree++
        }

        fmt.Printf(
            "Achat réussi : %s pour %d € !\n",
            item, prix,
        )

        utils.Pause()
    }
} */