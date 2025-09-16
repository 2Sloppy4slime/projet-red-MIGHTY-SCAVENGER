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

var quitfight bool = false

func (chara *Character) characterturn(enemy *Monster) {
	input := ""
	acted := false
	fmt.Println("a " + enemy.name + strconv.Itoa(enemy.hpnow) + " / " + strconv.Itoa(enemy.hpmax) + " stands before you")
	for !acted {

		fmt.Scanln(&input)
		switch input {
		case "close":
			quitfight = true
		case "list":
			fmt.Println("atk : Attack \n inv : Acces inv")
		case "atk":
			chara.attack(enemy)
			acted = true
		case "inv":
			chara.combatinv(enemy,&acted)

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
		if quitfight {
			quitfight = false
			return
		}
		if enemy.hpnow > 0 {
			goblinPattern(enemy, chara, turnnumber)
		}
		IsDead(chara)
	}
}
func (chara *Character) combatinv(enemy *Monster, actionbool *bool) {
	printInventory(chara)
	input := ""
	closeinv := false
	for !closeinv {
		fmt.Scanln(&input)
		switch input {
		case "list":
			fmt.Println(" -use : use an item \n -close : close your inv ")
		case "close":
			closeinv = true
		case "use":
			used:= false
			item := ""
			for !used {
				fmt.Scanln(&item)
				if chara.inv[item] > 0 {
					chara.Useitem(item)
					*actionbool = true
				}
			}
		case "":
			fmt.Println("input not recieved, how the fuck did you do that")
		default:
			fmt.Println("invalid input recieved, type \"list\" to get a list of valid commands")
		}
	}

}
