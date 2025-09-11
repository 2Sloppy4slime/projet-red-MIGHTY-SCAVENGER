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
func (chara *Character) characterturn(enemy *Monster) {
	input := ""
	acted := false
	for !acted {
		fmt.Println("a " + enemy.name + strconv.Itoa(enemy.hpnow) + " / " + strconv.Itoa(enemy.hpmax) + " stands before you")
		fmt.Scanln(&input)
		switch input {
		case "atk":
			chara.attack(enemy)
			acted = true
		default:
			fmt.Println("unrecognized command, try again or type \"list\" to see a list of actions")
		}
	}
}
func trainingFight(chara *Character, enemy *Monster) {
	turnnumber := 0
	for enemy.hpnow > 0 {
		turnnumber++
		fmt.Println("turn n° " + strconv.Itoa(turnnumber))
		chara.characterturn(enemy)
		if enemy.hpnow > 0 {
			goblinPattern(enemy, chara, turnnumber)
		}
		IsDead(chara)
	}
}
