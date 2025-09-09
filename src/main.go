package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var exit = false
var state = 0 //0 = base, 1 = inventaire, 2 = shop, 3 = forgeron, 4 = combat, 5 = combat inv, 6 = dead
var buying = false
var forging = false
var menuname = "main"

func main() {
	input := ""
	player := characterCreation()
	for !exit {
		fmt.Println(" \n\n ")
		fmt.Println("you are in the " + menuname + " menu (input \"list\" to get a list of valid commands)")
		fmt.Scanln(&input)
		commands := strings.Split(input, " ")
		for _, value := range commands {
			fmt.Println(value)
			menuHandler(value, player)
		}
	}
}

func menuHandler(command string, player Character) {
	if command == "list" {
		commandlist(state)
		return
	}
	switch state {
	case 0:
		switch command {
		case "info", "nfo":
			displayInfo(player)

		case "inv", "nv":
			state = 1
			menuname = "inventory"

		case "shop", "hop":
			state = 2
			menuname = "shop"

		case "forge", "orge":
			state = 3
			menuname = "forge"

		case "quit", "uit":
			fmt.Println(asciiArtLettering("see you when"))
			time.Sleep(1500 * time.Millisecond)
			fmt.Println(asciiArtLettering("you run out"))
			time.Sleep(1500 * time.Millisecond)
			fmt.Println(asciiArtLettering("of food again"))
			exit = true

		default:
			fmt.Println("unrecognized command: " + command + "    Please try again")
		}

	case 1: // inventaire menu
		switch command {
		case "view", "iew":
			printInventory(player)

		case "heal", "eal":
			takePot(&player)

		case "close", "lose":
			state = 0
			menuname = "main"
		case "shop", "hop":
			state = 2
			menuname = "shop"
		case "forge", "orge":
			state = 3
			menuname = "forge"
		default:
			fmt.Println("unrecognized command: " + command + "    Please try again")
		}

	case 2: //shop
		if buying {
			if shopinv[command] != 0 {
				buyItem(command, &player)
				buying = false
			}
		} else {
			switch command {
			case "buy", "uy":
				buying = true
				fmt.Println("type the name of the item you wish to buy:")
			case "close", "lose":
				state = 0
				menuname = "main"
			default:
				fmt.Println("unrecognized command: " + command + "    Please try again")
			}
		}

	case 3: //forge
		if forging {
			if forgelist[command] != 0 {
				buyItem(command, &player)
				buying = false
			}
		} else {
			switch command {
			case "make", "ake":
				printInventory(player)
			default:
				fmt.Println("unrecognized command: " + command + "    Please try again")
			}
		}
	case 4: //combat
		switch command {
		case "menu":
			printInventory(player)
		case "atk":
			printInventory(player)
		case "inv":
			state = 5
			menuname = "combat inventory"

		default:
			fmt.Println("unrecognized command: " + command + "    Please try again")
		}
	case 5: //combat inv
		switch command {
		case "view", "iew":
			printInventory(player)

		case "heal", "eal":
			takePot(&player)

		case "close", "lose":
			state = 4
			menuname = "combat"
		default:
			fmt.Println("unrecognized command: " + command + "    Please try again")

		}
	case 6:
		switch command {
		case "rez":
			player.hpnow = player.hpmax / 2
			state = 4
			menuname = "combat"
		case "quit":
			exit = true
		}
	default:
		fmt.Println("unrecognized gamestate ,going back to base")
		state = 0
		menuname = "main"
	}

}

func Poisonpot(chara *Character) {
	chara.hpnow -= 10
	IsDead(chara)
	time.Sleep(1 * time.Second)

	chara.hpnow -= 10
	IsDead(chara)
	time.Sleep(1 * time.Second)

	chara.hpnow -= 10
	fmt.Println("hp left :" + strconv.Itoa(chara.hpnow) + " / " + strconv.Itoa(chara.hpmax))
	IsDead(chara)
	time.Sleep(1 * time.Second)

}
