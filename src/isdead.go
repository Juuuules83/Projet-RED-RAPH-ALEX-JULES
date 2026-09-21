package main

import "fmt"

func isDead(c *Character) bool {
	countdead := 0
	if c.Pv <= 0 {
		fmt.Println("WASTED ")
		fmt.Println(c.Name, "est mort...")

		c.Pv = c.PvMax / 2
		fmt.Printf("%s est ressuscité avec %d / %d PV\n", c.Name, c.Pv, c.PvMax)
<<<<<<< HEAD
	
=======
		countdead++
>>>>>>> 0d2d4114f103aa05b6b956a4f9cc89b922b33080
		return true
	}
	if countdead > 0 {
		return false
	}
	return false
}
