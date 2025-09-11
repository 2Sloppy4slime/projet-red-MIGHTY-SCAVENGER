package main

import (
	"fmt"
	"strconv"
)

type Monster struct {
	name  string
	hpnow int
	hpmax int
	atk   int
}

func initGoblin() Monster {
	return Monster{"Goblin", 40, 40, 5}
}
func goblinPattern(goblin *Monster, player *Character, turnnumber int) {
	if turnnumber%3 == 0 {
		player.hpnow -= goblin.atk * 2
		fmt.Println("CCCCRRRAAASSSHHH! " + goblin.name + " deals " + strconv.Itoa(goblin.atk*2) + " damage to : " + player.name)
		fmt.Println(player.name + "s HP : " + strconv.Itoa(player.hpnow) + "/" + strconv.Itoa(player.hpmax))
	} else {
		player.hpnow -= goblin.atk
		fmt.Println(goblin.name + " deals " + strconv.Itoa(goblin.atk) + " damage to : " + player.name)
		fmt.Println(player.name + "s HP : " + strconv.Itoa(player.hpnow) + "/" + strconv.Itoa(player.hpmax))
	}

}

func (chara *Character) characterTurn(enemy *Monster) {
	turnnumber := 0
	for enemy.hpnow > 0 {
		turnnumber++
		input := ""
		acted := false
		for !acted {
			fmt.Println("a " + enemy.name + strconv.Itoa(enemy.hpnow) + " / " + strconv.Itoa(enemy.hpmax) + " stands before you")
			fmt.Scanln(&input)
			switch input {
			case "atk" :

			default:
				fmt.Println("unrecognized command, try again or type \"list\" to see a list of actions")
			}
		}
		goblinPattern(enemy, chara, turnnumber)
	}
}
