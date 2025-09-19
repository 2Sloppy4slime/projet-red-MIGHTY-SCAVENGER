package main

import (
	"fmt"
	"strconv"
	"time"
)

var spelldmg = map[string]int{"fireball": 18, "coupdepoing": 8}
var spellprice = map[string]int{"fireball": 50, "coupdepoing": 10}
var maxinvslots = 10

type Character struct {
	name  string
	class string
	level int
	hpmax int
	hpnow int
	money int
	inv   map[string]int
	skill []string
	armor Equipment
	atk   int
	spd   int
	exp   int
	mananow int
	manamax int
}

type Equipment struct {
	headslot  int //nombre de pv en plus
	chestslot int
	feetslot  int
}

func characterCreation() Character {
	name := ""
	namechosen := false
	for !namechosen {
		fmt.Println("What is your name??")
		fmt.Scanln(&name)
		for index, value := range name {
			if !IsLower(string(value)) && !IsUpper(string(value)) {
				fmt.Println("invalid name : " + name + " . please try again only using letters")
				break
			}
			if index == len(name)-1 {
				namechosen = true
			}
		}
	}

	classchosen := false
	health := 100
	atk := 5
	spd := 5
	name = Capitalize(ToLower(name))

	class := ""
	for !classchosen {
		fmt.Println("What is your species?? ( dwarf; elf; human; revenant) (type info for race descriptions)")
		fmt.Scanln(&class)
		switch class {
		case "info":
			fmt.Println(asciiArtLettering("dwarf | human"))
			fmt.Println("            more health, less speed                  |                        fully balanced character")
			fmt.Print("\n")
			fmt.Println(asciiArtLettering("elf | revenant"))
			fmt.Println("  more attack, less health    |                                 more speed, less attack")
			fmt.Print("\n")
		case "human":
			health = 100
			atk = 5
			spd = 5
			classchosen = true
		case "elf":
			health = 80
			atk = 8
			spd = 5
			classchosen = true
		case "dwarf":
			health = 120
			atk = 5
			spd = 3
			classchosen = true
		case "revenant":
			health = 100
			atk = 3
			spd = 8
			classchosen = true
		default:
			fmt.Println("this race does not exist yet... pick another ( dwarf; human; elf; revenant)")
		}
	}
	fmt.Println("Very well " + name + " the " + class + "... Good luck")

	return Character{name, class, 1, health, health / 2, 100, map[string]int{"potion": 3}, []string{"coupdepoing"}, Equipment{0, 0, 0}, atk, spd, 3,100,100}
}
func (chara *Character) getHealth() *int {
	return &chara.hpnow
}
func upgradeInventorySlot() {
	if maxinvslots < 40 {
		maxinvslots += 10
	}
}

func IsInvFull(chara *Character) bool {
	currentslots := 0
	for _, value := range chara.inv {
		currentslots += value
	}
	return currentslots >= maxinvslots
}

func IsDead(chara *Character) {
	if chara.hpnow < 0 {
		commandlist(6)
		command := ""
		fmt.Scanln(&command)
		switch command {
		case "rez":
			chara.hpnow = chara.hpmax / 2
		case "quit":
			exit = true
		}
	}
}

// fonction: retire un objet de l'inventaire
func removeitem(chara *Character, item string) {
	for i, _ := range chara.inv {
		if i == item {
			chara.inv[item] -= 1
		}
	}
}

func (chara *Character) additem(item string) {
	if !IsInvFull(chara) {
		if chara.inv[item] == 0 {
			chara.inv[item] = 1
			fmt.Println("new added!")
		} else {
			chara.inv[item] += 1
			fmt.Println("added!")
		}
	} else {
		fmt.Println("you're carrying too much stuff, item lost")
	}
}

func (chara *Character) Useitem(item string) {
	switch item {
	case "potion":
		chara.takePot()

	case "poisonpot":
		Poisonpot(chara)
	case "spellbook":
		chara.Spellbook()
	case "upgradeinv":
		upgradeInventorySlot()
	case "adventurers_hat":
		chara.equiparmor("adventurers_hat")
		removeitem(chara, "adventurers_hat")
	case "adventurers_tunic":
		chara.equiparmor("adventurers_tunic")
		removeitem(chara, "adventurers_tunic")
	case "adventurers_boots":
		chara.equiparmor("adventurers_boots")
		removeitem(chara, "adventurers_boots")
	case "inv_upgrade":
		upgradeInventorySlot()
	default:
		fmt.Println("item not found, try again")
	}
}

func (chara *Character) takePot() {
	health := chara.getHealth()
	if chara.inv["potion"] > 0 {
		removeitem(chara, "potion")
		if chara.hpnow+50 >= chara.hpmax {
			fmt.Println("drink")
			*health = chara.hpmax
		} else {
			fmt.Println("drink")
			*health += 50
		}
		printPot(chara)
	}
}

func (chara *Character) Spellbook() {
	for _, val := range chara.skill {
		if val == "fireball" {
			fmt.Println("you have learned this already, item removed :p")
			return
		}
	}
	chara.skill = append(chara.skill, "fireball")
	fmt.Println("learned fireball!!!!")
}

func (chara *Character) equiparmor(item string) {
	switch item {
	case "adventurers hat":
		chara.armor.headslot += 10
		chara.hpmax += 10
	case "adventurers tunic":
		chara.armor.chestslot += 25
		chara.hpmax += 10
	case "adventurers boots":
		chara.armor.feetslot += 15
		chara.hpmax += 10
	}
}

func (chara *Character) GetInv() *map[string]int {
	return &chara.inv
}

func (chara *Character) attack(enemy *Monster) {
	enemy.hpnow -= 5
	fmt.Println(chara.name + " deals " + strconv.Itoa(chara.atk) + " dmg to " + enemy.name)
}

func Poisonpot(chara *Character) {
	chara.hpnow -= 10
	fmt.Println("hp left :" + strconv.Itoa(chara.hpnow) + " / " + strconv.Itoa(chara.hpmax))
	IsDead(chara)
	time.Sleep(1 * time.Second)

	chara.hpnow -= 10
	fmt.Println("hp left :" + strconv.Itoa(chara.hpnow) + " / " + strconv.Itoa(chara.hpmax))
	IsDead(chara)
	time.Sleep(1 * time.Second)

	chara.hpnow -= 10
	fmt.Println("hp left :" + strconv.Itoa(chara.hpnow) + " / " + strconv.Itoa(chara.hpmax))
	IsDead(chara)
	time.Sleep(1 * time.Second)

}

func (chara *Character) experience(n int) {
	if chara.exp+n > 10*chara.level {
		chara.exp = chara.exp + n - 10*chara.level
		chara.level++
		switch chara.class {
		case "human":
			fmt.Print("LEVEL UP ! You are now level " + strconv.Itoa(chara.level) + "! (+1hp +1atk +1 spd)\n")
			chara.atk++
			chara.spd++
			chara.hpmax++
			chara.hpnow = chara.hpmax
		case "dwarf":
			fmt.Print("LEVEL UP ! You are now level " + strconv.Itoa(chara.level) + "! (+1atck +3hp ")
			chara.atk++
			chara.hpmax += 3
			chara.hpnow = chara.hpmax

			if chara.level%3 == 0 {
				fmt.Print("+1 spd)")
				chara.spd++
			} else {
				fmt.Print(")")
			}
		case "elf":
			fmt.Print("LEVEL UP ! You are now level " + strconv.Itoa(chara.level) + "! (+3atk +1 spd")
			chara.atk += 3
			chara.spd++

			if chara.level%3 == 0 {
				fmt.Print("+1 hp)")
				chara.hpmax++
				chara.hpnow = chara.hpmax
			} else {
				chara.hpnow = chara.hpmax
				fmt.Print(")")
			}
		case "revenant":
			fmt.Print("LEVEL UP ! You are now level " + strconv.Itoa(chara.level) + "! (+3spd +1 hp")
			chara.spd += 3
			chara.hpmax++
			chara.hpnow = chara.hpmax
			if chara.level%3 == 0 {
				fmt.Print("+1 atk)")
				chara.atk++
			} else {
				fmt.Print(")")
			}
		}

	}
}
