package main

import "fmt"

func isDead(c *Character) bool {
	countdead := 0
	if c.Pv <= 0 {
		fmt.Println("WASTED ")
		fmt.Println(c.Name, "est mort...")

		c.Pv = c.PvMax / 2
		fmt.Printf("%s est ressuscité avec %d / %d PV\n", c.Name, c.Pv, c.PvMax)
		countdead++
		return true
	}
	if countdead > 0 {
		return false
	}
	return false
}
