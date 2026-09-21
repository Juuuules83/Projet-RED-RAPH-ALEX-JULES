func isDead(c *Character) bool {
	if c.CurrentHP <= 0 {
		fmt.Println("WASTED ")
		fmt.Println(c.Name, "est mort...")

		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("%s est ressuscité avec %d / %d PV\n", c.Name, c.CurrentHP, c.MaxHP)
		return true
	}
	return false
}