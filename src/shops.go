package main

import "fmt"

var shopinv = map[string]int{"potion": -1, "poison_pot": 3, "spellbook:fireball": 1, "wolf_fur": 10, "troll_skin": 5, "boar_hide": 10, "crow_feather": 20}
var shopprice = map[string]int{"potion": 3, "poison_pot": 6, "spellbook:fireball": 25, "wolf_fur": 4, "troll_skin": 7, "boar_hide": 3, "crow_feather": 1}

func buyItem(item string, chara *Character) {

	if shopinv[item] != 0 {
		fmt.Println("succesfully bought :" + item)
		if shopinv[item] > 0 && chara.money >= shopprice[item] {
			shopinv[item]--

			additem(chara, item)
			chara.money -= shopprice[item]
		}
	}
}
