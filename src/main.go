package main

import "fmt"

func main() {
	exit := false
	input := ""
	player := initCharacter()
	displayInfo(player)
	accessInventory(player)
	takePot(&player)
	displayInfo(player)
	accessInventory(player)
	for !exit {
		fmt.Scan(input)
		switch input {
		case "exit":
			fmt.Println(asciiArtLettering("Goodbye"))
		}
	}
}
