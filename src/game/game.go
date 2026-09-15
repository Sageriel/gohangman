package game

import (
	"fmt"
)

func Start() {
	//Entrypoint
	ClearScreen()
	var i = 0
	fmt.Println(Blue + "Hangman" + Reset)
	fmt.Println("")
	fmt.Println(Green + "1. Start Game" + Reset)
	fmt.Println(Red + "2. Exit" + Reset)
	fmt.Println("")
	fmt.Print("Enter a number: ")
	fmt.Scan(&i)
	switch i {
	case 1:
		Game()

	case 2:
		ClearScreen()
		fmt.Println(Yellow + ":((((((" + Reset)
		return

	default:
		fmt.Println(Red + "Invalid choice!" + Reset)
		Start()
		return
	}

	//New Game Loop
	var k = ""
	for true {
		fmt.Println("")
		fmt.Print("Play again? (y/n) ")
		fmt.Scan(&k)
		switch k {
		case "y":
			Game()

		case "n":
			ClearScreen()
			fmt.Println(Yellow + ":((((((" + Reset)
			return
		}

	}
}
