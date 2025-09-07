package main

import (
	"fmt"
	"time"
)

var exit = false
var state = 0 //0 = base, 1 = inventaire, 2 = shop, 3 = combat
func main() {
	input := ""
	player := initCharacter()
	for !exit {
		fmt.Scanln(&input)
		
		menuHandler(input, player)
	}
}

func menuHandler(command string, player Character) {
	switch state {
	case 0:
		switch command {
		case "info":
			displayInfo(player)

		case "inventory":
			state = 1

		case "close":
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
		case "view":
			displayInfo(player)

		case "heal":
			printInventory(player)

		case "exit":

		default:
			fmt.Println("unrecognized command: " + command + "    Please try again")
		}
	}

}
