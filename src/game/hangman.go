package game

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

/*
* This function is designed to create a string containing a specified number of underscores ('_')
* based on the number it receives. We use this string to build the hidden representation
* of the word being guessed.
 */
func BlankBuilder(i int) string {
	var s = ""
	var count = 0
	for count != i {
		s = s + "_"
		count++
	}
	return s
}

/*
* This function ensures that the blank representation of the word is not displayed as
* a continuous sequence (e.g., "___"), but rather with spaces between the underscores
* (e.g., "_ _ _"). It inserts a space after each underscore character.
 */
func BlankDisplay(s string) string {
	var result = ""
	for i := 0; i < len(s); i++ {
		result += string(s[i]) + " "
	}
	return result
}

/*
* This function is used to clear the terminal screen,
* providing a cleaner and more user-friendly experience during gameplay.
 */
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}

/*
* This function is responsible for displaying the hangman figure.
* It has six stages, and the displayed stage advances
* whenever the player's fault count increases.
 */
func DisplayHangman(fault int) {
	stages := []string{
		`
+---+
|   |
    |
    |
    |
    |
=========`,
		`
+---+
|   |
O   |
    |
    |
    |
=========`,
		`
+---+
|   |
O   |
|   |
    |
    |
=========`,
		`
 +---+
 |   |
 O   |
/|   |
     |
     |
=========`,
		`
 +---+
 |   |
 O   |
/|\  |
     |
     |
=========`,
		`
 +---+
 |   |
 O   |
/|\  |
/    |
     |
=========`,
		`
 +---+
 |   |
 O   |
/|\  |
/ \  |
     |
=========`,
	}

	fmt.Println(stages[fault])
	fmt.Println("")
}

func HP(fault int) {
	for i := 0; i < 6; i++ {
		if i < 6-fault {
			fmt.Print("❤ ")
		} else {
			fmt.Print(" ♡")
		}
	}
	fmt.Println()
}

/*
* Colors used to create a more visually appealing and user-friendly interface.
 */
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

/*
* This is the main game logic. It is a simple implementation of the classic Hangman game.
* The objective is to guess the hidden word by entering letters. The player is
* allowed up to five incorrect guesses; the game ends on the sixth mistake.
* The game also displays all incorrectly guessed letters, helping the player
* keep track of previously used characters.
 */
func Game() {
	var input = ""
	word := GetRandomWord()
	wordRunes := []rune(word)
	wordLength := len(wordRunes)
	var fault = 0
	myWordRunes := []rune(BlankBuilder(wordLength))
	var wrongLetters []string
	var hint = 1
	var score int = 0

	ClearScreen()
	for string(myWordRunes) != word {
		DisplayHangman(fault)
		fmt.Print(BlankDisplay(string(myWordRunes)))
		fmt.Print(" ")
		HP(fault)
		fmt.Print("Wrong letters: ")
		for _, letter := range wrongLetters {
			fmt.Print(string(letter), ", ")
		}
		fmt.Println("")

		fmt.Print("Guess a letter: ")
		fmt.Scan(&input)

		ClearScreen()
		if len([]rune(input)) != 1 {
			fmt.Println(Red + "Enter only one letter!" + Reset)
			fmt.Println("")
			continue
		}
		input = strings.ToLower(input)

		if input == "-" {
			ClearScreen()
			fmt.Println(Yellow + ":((((((" + Reset)
			return
		}

		if input == "," {
			ClearScreen()
			fault = 6
		}

		if fault >= 5 {
			fmt.Println(Red + "You lost! :(" + Reset)
			fmt.Println(Red + "The word you were supposed to guess was: " + word + Reset)
			if score <= 0 {
				score = 0
			}
			fmt.Print(Red + "Score: " + Reset)
			fmt.Print(score)
			DisplayHangman(6)
			return
		}

		if input == "." {
			if hint <= 0 {
				fmt.Println(Red + "Already used your hint!" + Reset)
				fmt.Println("")
				continue
			}

			var remaining []int
			for k := 0; k < wordLength; k++ {
				if myWordRunes[k] == '_' {
					remaining = append(remaining, k)
				}
			}
			if len(remaining) == 0 {
				continue
			}

			randomIndex := remaining[rand.IntN(len(remaining))]
			hintLetter := wordRunes[randomIndex]

			for k := 0; k < wordLength; k++ {
				if wordRunes[k] == hintLetter {
					myWordRunes[k] = hintLetter
				}
			}

			wrongLetters = append(wrongLetters, input)
			hint--
			fmt.Println(Green + "Hint used!" + Reset)
			fmt.Println("")
			continue
		}

		inputRune := []rune(input)[0]

		exists := false
		for _, letter := range wrongLetters {
			if letter == input {
				exists = true
				break
			}
		}
		if exists {
			fmt.Println(Red + "You already tried that letter!" + Reset)
			fmt.Println("")
			continue
		}

		found := false
		for k := 0; k < wordLength; k++ {
			if inputRune == wordRunes[k] {
				myWordRunes[k] = inputRune
				found = true
			}
		}

		if found {
			fmt.Println(Green + "Nice one!" + Reset)
			fmt.Println("")
			score = score + 10
		} else {
			wrongLetters = append(wrongLetters, input)
			fault++
			fmt.Println(Red + "Wrong guess!" + Reset)
			fmt.Println("")
			score = score - 5
		}
	}
	ClearScreen()
	fmt.Println(Green + "Congratulations, you won!" + Reset)
	fmt.Println(Green + "The word was: " + word + Reset)
	if score <= 0 {
		score = 0
	}
	fmt.Print(Green + "Score: " + Reset)
	fmt.Print(score)
	DisplayHangman(fault)
}
