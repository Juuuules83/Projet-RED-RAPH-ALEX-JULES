

package combat

import (
	"projet-red/player"
	"fmt"
	"time"
)

func UtiliserPotionVie(c *player.Character) string {
	return c.UtiliserPotionVie()	
}

func PoisonPot(m *Monster) {
    for i := 0; i < 3; i++ {
        m.Pv -= 10
        if m.Pv < 0 {
            m.Pv = 0
        }
        fmt.Printf("%s : %d/%d PV\n", m.Name, m.Pv, m.PvMax)
        time.Sleep(1 * time.Second)
    }
}