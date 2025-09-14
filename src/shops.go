package main

import "fmt"

var shopinv = map[string]int{"potion": -1, "poison_pot": 3, "spellbook:fireball": 1, "wolf_fur": 10, "troll_skin": 5, "boar_hide": 10, "crow_feather": 20, "inv_upgrade": 4}
var shopprice = map[string]int{"potion": 3, "poison_pot": 6, "spellbook:fireball": 25, "wolf_fur": 4, "troll_skin": 7, "boar_hide": 3, "crow_feather": 1, "inv_upgrade": 30}
var forgelist = map[string]int{"adventurers_hat": 1, "adventurers_tunic": 1, "adventurers_boots": 1}

func (chara *Character) buyItem(item string) {

	if shopinv[item] != 0 && chara.money >= shopprice[item] {
		fmt.Println("succesfully bought :" + item)
		shopinv[item]--
		chara.additem(item)
		chara.money -= shopprice[item]
	}
}

func (chara *Character) forgeItem(item string) {
	switch item {
	case "adventurers_hat":
		if !IsInvFull(chara) && chara.money >= 5 && chara.inv["crow_feather"] > 0 && chara.inv["boar_hide"] > 0 {
			chara.money -= 5
			removeitem(chara, "crow_feather")
			removeitem(chara, "boar_hide")
			chara.additem("adventurers_hat")
		}
	case "adventurers_tunic":
		if !IsInvFull(chara) && chara.money >= 5 && chara.inv["wolf_fur"] > 1 && chara.inv["troll_skin"] > 0 {
			chara.money -= 5
			removeitem(chara, "wolf_fur")
			removeitem(chara, "wolf_fur")
			removeitem(chara, "troll_skin")
			chara.additem("adventurers_tunic")
		}
	case "adventurers_boots":
		if !IsInvFull(chara) && chara.money >= 5 && chara.inv["wolf_fur"] > 0 && chara.inv["boar_hide"] > 0 {
			chara.money -= 5
			removeitem(chara, "wolf_fur")
			removeitem(chara, "boar_hide")
			chara.additem("adventurers_boots")
		}
	}
}
