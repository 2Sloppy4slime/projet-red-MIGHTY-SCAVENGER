package main

import (
	"slices"
)

func initCharacter() Character {
	return Character{"Luca", "Elfe", 1, 100, 40, []string{"potion", "potion", "potion"}}
}

func takePot(chara *Character) {
	if slices.Contains(chara.inv, "potion") {
		chara.inv = remove(chara.inv, "potion")
		if chara.hpnow+50 >= chara.hpmax {
			chara.hpnow = chara.hpmax
		} else {
			chara.hpnow += 50
		}
		printPot(chara)
	}
}

type Character struct {
	name  string
	class string
	level int
	hpmax int
	hpnow int
	inv   []string
}

func remove[T comparable](l []T, item T) []T {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}
