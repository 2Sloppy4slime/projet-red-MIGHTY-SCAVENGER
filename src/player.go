package main

import (
	"fmt"
	"strconv"
	"time"
)

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
	name = Capitalize(ToLower(name))

	class := ""
	for !classchosen {
		fmt.Println("What is your species?? ( dwarf; human; elf)")
		fmt.Scanln(&class)
		switch class {
		case "human":
			health = 100
			classchosen = true
		case "elf":
			health = 80
			classchosen = true
		case "dwarf":
			health = 120
			classchosen = true
		default:
			fmt.Println("what the fuck is that race I don't know it")
		}
	}
	fmt.Println("Very well " + name + " the " + class + "... Good luck")

	return Character{name, class, 1, health, health / 2, 100, map[string]int{"potion": 3}, []string{"coupdepoing"}, Equipment{0, 0, 0}, 5}
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

	chara.skill = append(chara.skill, "fireball")
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
