package player

import (
    "fmt"
    "projet-red/utils"
)

func (c *Character) RecruterMentor(nom string) {
    if c.Mentors == nil {
        c.Mentors = make(map[string]bool)
    }
    if c.Mentors[nom] {
        return
    }
    c.Mentors[nom] = true
    if nom == "Cyril" {
        fmt.Println(utils.Bold + utils.Yellow + "Cyril recruté - il multiplie les dégâts par 2 en mettant des supra Tunnel à tout le monde !" + utils.Reset)
    } else {
        c.PvMax += 20
        c.Pv += 20
        fmt.Println(utils.Bold + utils.Green + "Mentor recruté - 20 PV supplémentaires !" + utils.Reset)
    }
}
