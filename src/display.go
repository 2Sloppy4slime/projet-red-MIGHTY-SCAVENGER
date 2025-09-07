package main

import (
	"fmt"
	"strconv"
	"strings"
)

// fonction :
func displayInfo(chara Character) {
	charastrings := []string{"   name : " + chara.name + "   ", "   classe : " + chara.class + "   ", "   level : " + strconv.Itoa(chara.level) + "   ", "   HP : " + strconv.Itoa(chara.hpnow) + "/" + strconv.Itoa(chara.hpmax) + "   "}
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
	fmt.Println("    |" + repeat + "\n" + "    |   " + strings.Repeat("_", width-3) + "|___\n    |  /" + strings.Repeat(" ", width) + "/.\n    \\_/" + strings.Repeat("_", width) + "/.")
}

// fonction :
func accessInventory(chara Character) {
	fmt.Println("\n  .--.      .-'.      .--.      .--.      .--.      .--.      .`-.      .--.\n:::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\ \n'      `--'      `.-'      `--'      `--'      `--'      `-.'      `--'      `")
	fmt.Println(asciiArtLettering(" inventory"))
	inventory := chara.inv
	printmap := make(map[string]int) //map pour stocker la quantité d'items pour afficher un "potion x3"
	for _, item := range inventory {
		_, exists := printmap[item]

		if exists {
			printmap[item]++
		} else {
			printmap[item] = 1
		}
	}
	for key, value := range printmap {
		letter := ""
		if value > 1 {
			letter = asciiArtLettering("-" + key + " x" + strconv.Itoa(value))
		} else {
			letter = asciiArtLettering("-" + key)
		}
		fmt.Println(" ")
		fmt.Println(letter)
	}
	fmt.Println("\n  .--.      .-'.      .--.      .--.      .--.      .--.      .`-.      .--.\n:::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\::::::::.\\ \n'      `--'      `.-'      `--'      `--'      `--'      `-.'      `--'      `")
}

func printPot(chara *Character) {
	fmt.Println("Bu potion, PV : " + strconv.Itoa(chara.hpnow) + " / " + strconv.Itoa(chara.hpmax))
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
