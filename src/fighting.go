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
	spd   int
}

func initGoblin() Monster {
	return Monster{"Goblin", 40, 40, 5, 5}
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
			acted = true
		case "list":
			fmt.Println(" - atk : Attack \n - inv : Acces inv \n - spell : cast a spell \n - close : exit out of the fight")
		case "atk":
			chara.attack(enemy)
			acted = true
		case "inv":
			chara.combatinv(enemy, &acted)
		case "spell":
			input := ""
			spelled := false
			for !spelled {
				chara.displayspells()
				fmt.Scanln(&input)
				if input == "close" {
					spelled = true
				}
				for _, spell := range chara.skill {

					if input == spell {
						if chara.mananow >= spellprice[spell] {
							enemy.hpnow -= spelldmg[spell]
							chara.mananow -= spellprice[spell]
							spelled = true
							acted = true
						} else {
							fmt.Println("not enough mana for spell, try again")
						}
					}
				}
			}
		default:
			fmt.Println("unrecognized command, try again or type \"list\" to see a list of actions")
		}
	}
}
func trainingFight(chara *Character, enemy *Monster) {
	turnnumber := 0
	youfirst := false
	if chara.spd >= enemy.spd {
		youfirst = true
	}
	for enemy.hpnow > 0 {
		turnnumber++
		fmt.Println("-----------------------------\n turn n° " + strconv.Itoa(turnnumber) + "\n-----------------------------")
		if youfirst {
			chara.characterturn(enemy)
			if quitfight {
				quitfight = false
				return
			}
			if enemy.hpnow > 0 {
				goblinPattern(enemy, chara, turnnumber)
			}
			IsDead(chara)
		} else {
			if enemy.hpnow > 0 {
				goblinPattern(enemy, chara, turnnumber)
			}
			chara.characterturn(enemy)
			if quitfight {
				quitfight = false
				return
			}
			IsDead(chara)
		}

	}

	fmt.Println("----------------------------------\n you have defeated the training goblin! mana has been restored")
	chara.mananow = chara.manamax
	chara.experience(3)
	fmt.Println("+3exp (" + strconv.Itoa(chara.exp) + "/" + strconv.Itoa(chara.level*10) + ")")
	chara.money += 10
	fmt.Println("+10 gold (you have " + strconv.Itoa(chara.money) + " gold now")

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
			used := false
			item := ""
			for !used {
				fmt.Scanln(&item)
				if chara.inv[item] > 0 {
					chara.Useitem(item, enemy)
					*actionbool = true
					used = true
				}
			}
		case "":
			fmt.Println("input not recieved, how the fuck did you do that")
		default:
			fmt.Println("invalid input recieved, type \"list\" to get a list of valid commands")
		}
	}

}
