package main

import (
	"fmt"
	"strings"
	"time"
)

var exit = false
var state = 0 //0 = base, 1 = inventaire, 2 = shop, 3 = forgeron, 4 = quest, 6 = dead
var buying = false
var forging = false
var menuname = "main"

func main() {
	legacy := false
	a := ""
	choselegacy := false
	for !choselegacy {
		fmt.Println("play the legacy edition? (Y/N)")
		fmt.Scanln(&a)
		switch a {
		case "Y":
			legacy = true
			choselegacy = true
		case "N":
			legacy = false
			choselegacy = true

		case "quit":
			return
		default:
			fmt.Println("unknown command, try again")
		}
	}
	if legacy {
		input := ""
		player := characterCreation()
		for !exit {
			if !forging && !buying {
				fmt.Println("you are in the " + menuname + " menu (input \"list\" to get a list of valid commands)")
			} else if forging {
				forgingtext()
			} else {
				buyingtext()
			}
			fmt.Scanln(&input)
			commands := strings.Split(input, " ")
			for _, value := range commands {
				fmt.Println(value)
				menuHandler(value, &player)
			}
		}
	}
}

func menuHandler(command string, player *Character) {
	if command == "list" || command == "ist" {
		commandlist(state)
		return
	}
	if command == "quit" || command == "uit" {
		fmt.Println(asciiArtLettering("see you when"))
		time.Sleep(1500 * time.Millisecond)
		fmt.Println(asciiArtLettering("you run out"))
		time.Sleep(1500 * time.Millisecond)
		fmt.Println(asciiArtLettering("of food again"))
		exit = true
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

		case "train", "rain":
			enemy := initGoblin()
			trainingFight(player, &enemy)

		default:
			fmt.Println("unrecognized command: " + command + "    Please try again")
		}

	case 1: // inventaire menu
		switch command {
		case "view", "iew":
			printInventory(player)

		case "heal", "eal":
			player.takePot()

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
				player.buyItem(command)
				buying = false
			} else {
				displayshop()
				fmt.Println("item not recognized, type again")
			}
		} else {
			switch command {
			case "buy", "uy":
				buying = true
				fmt.Println("type the name of the item you wish to buy:")
				displayshop()
			case "close", "lose":
				state = 0
				menuname = "main"
			default:
				fmt.Println("unrecognized command: " + command + "    Please try again")
			}
		}

	case 3: //forge
		if forging {
			if command == "stop" {
				forging = false
			}
			if forgelist[command] != 0 {
				player.forgeItem(command)
				buying = false
			} else {
				forgingtext()
				fmt.Println("unrecognized item, type stop if you want to stop forging")
			}
		} else {
			switch command {
			case "make", "ake":
				forging = true
				forgingtext()
			case "close", "lose":
				state = 0
				menuname = "main"
			default:
				fmt.Println("unrecognized command: " + command + "    Please try again")

			}

		}
	case 4: //quest
		switch command {
		case "inv":
			state = 5
			menuname = "combat inventory"

		default:
			fmt.Println("unrecognized command: " + command + "    Please try again")
		}

	case 6: //dead

	default:
		fmt.Println("unrecognized gamestate ,going back to base")
		state = 0
		menuname = "main"
	}

}
