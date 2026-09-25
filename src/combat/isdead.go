package combat

import (
	"fmt"
	"os"

	"projet-red/player"
	"projet-red/utils"
)

func IsDead(characterPV int) bool {
	return characterPV <= 0
}

func HandlePlayerDeath(c *player.Character) {
    if !IsDead(c.Pv) {
        return
    }

    c.CountDeath++

    if c.CountDeath == 1 {
        c.Pv = c.PvMax / 2

        fmt.Println(utils.Bold + utils.Green +
            "\nVous avez été ressuscité avec 50% de vos PV !\n" +
            utils.Reset)

        utils.Pause()
        return
    }

    if c.CountDeath > 1 {
utils.ClearScreen()
	fmt.Println(utils.Red + utils.Bold)
	fmt.Println(" ██████╗  █████╗ ███╗   ███╗███████╗     ██████╗ ██╗   ██╗███████╗██████╗")
	fmt.Println("██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ██╔═══██╗██║   ██║██╔════╝██╔══██╗")
	fmt.Println("██║  ███╗███████║██╔████╔██║█████╗      ██║   ██║██║   ██║█████╗  ██████╔╝")
	fmt.Println("██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║   ██║╚██╗ ██╔╝██╔══╝  ██╔══██╗")
	fmt.Println("╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ╚██████╔╝ ╚████╔╝ ███████╗██║  ██║")
	fmt.Println(" ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═╝")
	fmt.Println(utils.Reset)
	utils.Pause()
	os.Exit(0) 
    }
}