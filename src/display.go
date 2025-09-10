package main

import (
	"fmt"
	"strconv"
	"strings"
)

func commandlist(state int) {
	switch state {
	case 0:
		fmt.Println("main menu : \n - info : see your stats \n - inv : access the inventory menu \n - quit : leave the game")
	case 1:
		fmt.Println("Inventory : \n - view : see the contents of your inventory \n - use (item) : consume an item\n - heal : use a healing potion from your inventory \n - shop : enter the shop \n - forge : access the forgemasters shop\n - close : exit out of your inventory")
	case 2:
		displayshop()
		fmt.Println("Merchant : \n - buy (item) \n - close : go back to base menu")
	case 3:
		fmt.Println("oui")
	case 4:
		fmt.Println("oui")
	case 5:
		fmt.Println("bjonur")
	case 6:
		fmt.Println(asciiArtLettering("you have died"))
		fmt.Println("     continue? \n - rez : revive with 50%hp \n - quit : leave the game")

	}

}
func displayshop() {
	fmt.Println(asciiArtLettering("merchant"))
	fmt.Println("the merchant has :")
	for key, value := range shopinv {
		stock := strconv.Itoa(value)
		if value == -1 {
			stock = "inf"
		}
		println(key + " : " + strconv.Itoa(shopprice[key]) + "coins" + "(stock : " + stock)
	}
}
func forgingtext() {
	fmt.Println("the forgemaster can make you : \n - \"adventurers hat\" : +10HP (1 crow feather,1 boar hide)\n - \"adventurers tunic\" : +25HP (2 wolf fur,1 troll skin) \n - \"adventurers boots\" : +10HP (1 wolf fur,1 boar hide)")
	fmt.Println("type the name of the equipment to forge it")
}
func buyingtext() { fmt.Println("oui") }

// fonction :
func displayInfo(chara Character) {
	charastrings := []string{"   name : " + chara.name + "   ", "   class : " + chara.class + "   ", "   level : " + strconv.Itoa(chara.level) + "   ", "   HP : " + strconv.Itoa(chara.hpnow) + "/" + strconv.Itoa(chara.hpmax) + "   ", "   Money : " + strconv.Itoa(chara.money) + "   "}
	longeststring := ""
	for _, word := range charastrings {
		if len(word) > len(longeststring) {
			longeststring = word

		}
	}
	for index, word := range charastrings {
		if len(word) < len(longeststring) {
			charastrings[index] = word + strings.Repeat(" ", len(longeststring)-len(word))
		}
	}
	width := len(longeststring)
	repeat := strings.Repeat(" ", width) + "\\."
	fmt.Println("   __" + strings.Repeat("_", width) + "\n" + " / \\ " + repeat)
	fmt.Println("|   |" + charastrings[0] + "\\.")  //nom
	fmt.Println(" \\_ |" + charastrings[1] + "\\.") //classe
	fmt.Println("    |" + charastrings[2] + "\\.")  //level
	fmt.Println("    |" + charastrings[3] + "\\.")  //pv
	fmt.Println("    |" + charastrings[4] + "\\.")  //money
	fmt.Println("    |" + repeat + "\n" + "    |   " + strings.Repeat("_", width-3) + "|___\n    |  /" + strings.Repeat(" ", width) + "/.\n    \\_/" + strings.Repeat("_", width) + "/.")
}

// fonction :
func printInventory(chara Character) {
	fmt.Println("\n  .--.      .-'.      .--.      .--.      .--.      .--.      .`-.      .--.\n:::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\ \n'      `--'      `.-'      `--'      `--'      `--'      `-.'      `--'      `")
	fmt.Println(asciiArtLettering(" inventory:"))
	inventory := chara.inv
	for key, value := range inventory {
		letter := ""
		if value > 1 {
			letter = "-" + key + " x" + strconv.Itoa(value)
		} else {
			letter = "-" + key
		}
		fmt.Println(" ")
		fmt.Println(letter)
	}
	fmt.Println("\n  .--.      .-'.      .--.      .--.      .--.      .--.      .`-.      .--.\n:::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\ \n'      `--'      `.-'      `--'      `--'      `--'      `-.'      `--'      `")
}

func printPot(chara *Character) {
	fmt.Println("Drunk potion, HP : " + strconv.Itoa(chara.hpnow) + " / " + strconv.Itoa(chara.hpmax))
}

func asciiArtLettering(s string) string {
	res := ""
	lines := []string{"", "", "", "", "", ""}
	for _, letter := range s {
		switch letter {
		case '0':
			lines[0] += "   ___  "
			lines[1] += "  / _ \\ "
			lines[2] += " | | | |"
			lines[3] += " | |_| |"
			lines[4] += "  \\___/ "
			lines[5] += "        "

		case '1':
			lines[0] += "  _ "
			lines[1] += " / |"
			lines[2] += " | |"
			lines[3] += " | |"
			lines[4] += " |_|"
			lines[5] += "    "
		case '2':
			lines[0] += "  ____  "
			lines[1] += " |___ \\ "
			lines[2] += "   __) |"
			lines[3] += "  / __/ "
			lines[4] += " |_____|"
			lines[5] += "        "
		case '3':
			lines[0] += "  _____ "
			lines[1] += " |___ / "
			lines[2] += "   |_ \\ "
			lines[3] += "  ___) |"
			lines[4] += " |____/ "
			lines[5] += "        "
		case '4':
			lines[0] += "  _  _   "
			lines[1] += " | || |  "
			lines[2] += " | || |_ "
			lines[3] += " |__   _|"
			lines[4] += "    |_|  "
			lines[5] += "        "
		case '5':
			lines[0] += "  ____  "
			lines[1] += " | ___| "
			lines[2] += " |___ \\ "
			lines[3] += "  ___) |"
			lines[4] += " |____/ "
			lines[5] += "        "
		case '6':
			lines[0] += "   __   "
			lines[1] += "  / /_  "
			lines[2] += " | '_ \\ "
			lines[3] += " | (_) |"
			lines[4] += "  \\___/ "
			lines[5] += "        "
		case '7':
			lines[0] += "  _____ "
			lines[1] += " |___  |"
			lines[2] += "    / / "
			lines[3] += "   / /  "
			lines[4] += "  /_/   "
			lines[5] += "        "
		case '8':
			lines[0] += "   ___  "
			lines[1] += "  ( _ ) "
			lines[2] += "  / _ \\ "
			lines[3] += " | (_) |"
			lines[4] += "  \\___/ "
			lines[5] += "        "
		case '9':
			lines[0] += "   ___  "
			lines[1] += "  / _ \\ "
			lines[2] += " | (_) |"
			lines[3] += "  \\__, |"
			lines[4] += "    /_/ "
			lines[5] += "        "
		case 'a':
			lines[0] += "        "
			lines[1] += "   __ _ "
			lines[2] += "  / _` |"
			lines[3] += " | (_| |"
			lines[4] += "  \\__,_|"
			lines[5] += "        "

		case 'b':
			lines[0] += "  _     "
			lines[1] += " | |__  "
			lines[2] += " | '_ \\ "
			lines[3] += " | |_) |"
			lines[4] += " |_.__/ "
			lines[5] += "        "
		case 'c':
			lines[0] += "       "
			lines[1] += "   ___ "
			lines[2] += "  / __|"
			lines[3] += " | (__ "
			lines[4] += "  \\___|"
			lines[5] += "       "
		case 'd':
			lines[0] += "      _ "
			lines[1] += "   __| |"
			lines[2] += "  / _` |"
			lines[3] += " | (_| |"
			lines[4] += "  \\__,_|"
			lines[5] += "        "
		case 'e':
			lines[0] += "       "
			lines[1] += "   ___ "
			lines[2] += "  / _ \\"
			lines[3] += " |  __/"
			lines[4] += "  \\___|"
			lines[5] += "       "
		case 'f':
			lines[0] += "   __ "
			lines[1] += "  / _|"
			lines[2] += " | |_ "
			lines[3] += " |  _|"
			lines[4] += " |_|  "
			lines[5] += "      "
		case 'g':
			lines[0] += "        "
			lines[1] += "   __ _ "
			lines[2] += "  / _` |"
			lines[3] += " | (_| |"
			lines[4] += "  \\__, |"
			lines[5] += "  |___/ "
		case 'h':
			lines[0] += "  _     "
			lines[1] += " | |__  "
			lines[2] += " | '_ \\ "
			lines[3] += " | | | |"
			lines[4] += " |_| |_|"
			lines[5] += "        "
		case 'i':
			lines[0] += "  _ "
			lines[1] += " (_)"
			lines[2] += " | |"
			lines[3] += " | |"
			lines[4] += " |_|"
			lines[5] += "    "
		case 'j':
			lines[0] += "    _ "
			lines[1] += "   (_)"
			lines[2] += "   | |"
			lines[3] += "   | |"
			lines[4] += "  _/ |"
			lines[5] += " |__/ "
		case 'k':
			lines[0] += "  _    "
			lines[1] += " | | __"
			lines[2] += " | |/ /"
			lines[3] += " |   < "
			lines[4] += " |_|\\_\\"
			lines[5] += "       "
		case 'l':
			lines[0] += "  _ "
			lines[1] += " | |"
			lines[2] += " | |"
			lines[3] += " | |"
			lines[4] += " |_|"
			lines[5] += "    "
		case 'm':
			lines[0] += "            "
			lines[1] += "  _ __ ___  "
			lines[2] += " | '_ ` _ \\ "
			lines[3] += " | | | | | |"
			lines[4] += " |_| |_| |_|"
			lines[5] += "            "
		case 'n':
			lines[0] += "        "
			lines[1] += "  _ __  "
			lines[2] += " | '_ \\ "
			lines[3] += " | | | |"
			lines[4] += " |_| |_|"
			lines[5] += "        "
		case 'o':
			lines[0] += "        "
			lines[1] += "   ___  "
			lines[2] += "  / _ \\ "
			lines[3] += " | (_) |"
			lines[4] += "  \\___/ "
			lines[5] += "        "
		case 'p':
			lines[0] += "        "
			lines[1] += "  _ __  "
			lines[2] += " | '_ \\ "
			lines[3] += " | |_) |"
			lines[4] += " | .__/ "
			lines[5] += " |_|    "
		case 'q':
			lines[0] += "        "
			lines[1] += "   __ _ "
			lines[2] += "  / _` |"
			lines[3] += " | (_| |"
			lines[4] += "  \\__, |"
			lines[5] += "     |_|"
		case 'r':
			lines[0] += "       "
			lines[1] += "  _ __ "
			lines[2] += " | '__|"
			lines[3] += " | |   "
			lines[4] += " |_|   "
			lines[5] += "       "
		case 's':
			lines[0] += "      "
			lines[1] += "  ___ "
			lines[2] += " / __|"
			lines[3] += " \\__ \\"
			lines[4] += " |___/"
			lines[5] += "      "
		case 't':
			lines[0] += "  _   "
			lines[1] += " | |_ "
			lines[2] += " | __|"
			lines[3] += " | |_ "
			lines[4] += "  \\__|"
			lines[5] += "      "
		case 'u':
			lines[0] += "        "
			lines[1] += "  _   _ "
			lines[2] += " | | | |"
			lines[3] += " | |_| |"
			lines[4] += "  \\__,_|"
			lines[5] += "        "
		case 'v':
			lines[0] += "        "
			lines[1] += " __   __"
			lines[2] += " \\ \\ / /"
			lines[3] += "  \\ V / "
			lines[4] += "   \\_/  "
			lines[5] += "        "
		case 'w':
			lines[0] += "           "
			lines[1] += " __      __"
			lines[2] += " \\ \\ /\\ / /"
			lines[3] += "  \\ V  V / "
			lines[4] += "   \\_/\\_/  "
			lines[5] += "           "
		case 'x':
			lines[0] += "       "
			lines[1] += " __  __"
			lines[2] += " \\ \\/ /"
			lines[3] += "  >  < "
			lines[4] += " /_/\\_\\"
			lines[5] += "       "
		case 'y':
			lines[0] += "        "
			lines[1] += "  _   _ "
			lines[2] += " | | | |"
			lines[3] += " | |_| |"
			lines[4] += "  \\__, |"
			lines[5] += "  |___/ "
		case 'z':
			lines[0] += "      "
			lines[1] += "  ____"
			lines[2] += " |_  /"
			lines[3] += "  / / "
			lines[4] += " /___|"
			lines[5] += "      "
		case '-':
			lines[0] += "           "
			lines[1] += "           "
			lines[2] += "  _________"
			lines[3] += " /        /"
			lines[4] += "/________/ "
			lines[5] += "           "
		case '_':
			lines[0] += "           "
			lines[1] += "           "
			lines[2] += "           "
			lines[3] += "           "
			lines[4] += " _________ "
			lines[5] += "/________/ "
		case ':':
			lines[0] += "     _     "
			lines[1] += "    (_)    "
			lines[2] += "           "
			lines[3] += "     _     "
			lines[4] += "    (_)    "
			lines[5] += "           "
		default:
			lines[0] += "        "
			lines[1] += "        "
			lines[2] += "        "
			lines[3] += "        "
			lines[4] += "        "
			lines[5] += "        "
		}
	}
	res = lines[0] + "\n" + lines[1] + "\n" + lines[2] + "\n" + lines[3] + "\n" + lines[4] + "\n" + lines[5]
	return res
}
