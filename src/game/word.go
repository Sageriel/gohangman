package game

import (
	"encoding/json"
	"math/rand/v2"
	"os"
)

type Word struct {
	Word string `json:"word"`
}

type WordList struct {
	Words []Word `json:"words"`
}

/*
* This function is responsible for reading and loading the words from a JSON file.
* It parses the JSON data and makes the word list available for use within the game.
 */
func GetRandomWord() string {
	data, err := os.ReadFile("data/words.json")
	if err != nil {
		panic(err)
	}
	var random = rand.IntN(20)
	var wordList WordList

	err = json.Unmarshal(data, &wordList)
	if err != nil {
		panic(err)
	}
	var randomWord = wordList.Words[random]
	randomString := randomWord.Word

	return randomString

}
