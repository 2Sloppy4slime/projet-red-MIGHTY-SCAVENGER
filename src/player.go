package main

import (
	"fmt"
	"slices"
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
		}
	}
	fmt.Println("Very well " + name + " the " + class + "... Good luck")

	return Character{name, class, 1, health, health / 2, 100, map[string]int{"potion": 3}, []string{"coupdepoing"}, Equipment{0, 0, 0}}
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
		state = 6
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

func additem(chara *Character, item string) {
	if !IsInvFull(chara) {
		if chara.inv[item] == 0 {
			chara.inv[item] = 1
		} else {
			chara.inv[item] += 1
		}
	} else {
		fmt.Println("you're carrying too much stuff, item lost")
	}
}

func Useitem(chara *Character, item string) {
	switch item {
	case "potion":
		takePot(chara)
	case "poisonpot":
		Poisonpot(chara)
	case "spellbook":
		Spellbook(chara)
	case "upgradeinv":
		upgradeInventorySlot()
	default:
		fmt.Println("item not found, try again")
	}
}

func takePot(chara *Character) {
	if chara.inv["potion"] >0 {
		removeitem(chara, "potion")
		if chara.hpnow+50 >= chara.hpmax {
			chara.hpnow = chara.hpmax
		} else {
			chara.hpnow += 50
		}
		printPot(chara)
	}
}

func Spellbook(chara *Character) {

	chara.inv = append(chara.inv, "fireball")
}
