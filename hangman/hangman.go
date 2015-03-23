/*
  This is a program that will allow the user to play
  hangman from the command line.  We will ask them for
  thier guess and then we will display the word with
  underscores for them to guess.  As the user guesses
  letters correctly we will fill in the underscores
  with the letters.  Right now it works with single
  words, I will add support for multiwords soon.
 */

package hangman

import (
		"fmt"
 		"log"
 		"strings"
 		"regexp"
 		"unicode/utf8"
)

var guessedLetters []string = make([]string, 0)

func Play() {
	var fullWord string
	InitGame(&fullWord)
	blankedWord := blankOutLetters(fullWord)
	TakeGuess(fullWord, blankedWord)
}

// InitGame will start the game off and get the word we want to use for the game.
func InitGame(word *string) {
	fmt.Printf("Welcome to the hangman game, please submit the word.\n")
	_, err := fmt.Scanf("%s", word)
	if err != nil {
		fmt.Printf("Hangman has failed with the following fatal error.\n")
		log.Fatal(err)
		return
	}
}

// Show the word to the console blanked out with underscores
// and take in the user guess.
func TakeGuess(fullWord string, oldBlankedWord string) {
	fmt.Printf("Current word: %s.\n", oldBlankedWord)
	fmt.Printf("Guessed letters: %v \n", guessedLetters)
	fmt.Printf("What is your guess?\n")

	var letterGuess string
	_, err := fmt.Scanf("%s", &letterGuess)
	if err != nil {
		fmt.Print("Hangman has failed with the following fatal error.\n")
		log.Fatal(err)
		return
	}

	if (containsStr(guessedLetters, letterGuess)) {
		fmt.Println("You have alredy guessed this letter.")
		TakeGuess(fullWord, oldBlankedWord)
	} else if utf8.RuneCountInString(letterGuess) > 1{
		fmt.Println("You must enter just one letter")
	} else {
		guessedLetters = append(guessedLetters, letterGuess)
	}


	newBlankedWord := fillBlankedWord(string(letterGuess), oldBlankedWord, fullWord)
	fmt.Println("This is the new blanked word: ", newBlankedWord)
	// See if there are more letters to guess.
	if strings.ContainsAny(newBlankedWord, "_") {
		if len(guessedLetters) >= 10 {
			fmt.Println("You guessed 10 times and still do not have the word.\n")
			fmt.Println("Game Over Bitch.\n")
		} else {
        	TakeGuess(fullWord, newBlankedWord)
		}
	} else {
		fmt.Println("You guessed the word: ", fullWord)
	}
}

func blankOutLetters(fullWord string) string {
	letterRegex := regexp.MustCompile("[a-z]")
	blanker := func(r rune) rune {
		switch {
		case letterRegex.MatchString(string(r)):
			return '_'
		default:
			return r
		}
	}
	return strings.Map(blanker, fullWord)
}

func fillBlankedWord(letter string, blankedWord string, fullWord string) string {
	re := regexp.MustCompile(letter)
	// Get the indexes for the guess.
	indexRes := re.FindAllStringIndex(fullWord, -1)
    indexes:= make([]int, 0)
    for _, v := range indexRes {
      indexes = append(indexes, v[0])
    }
	// We will create another string to return with the updated
	// values.
	returnStr := make([]string, 0)
	for i, v := range blankedWord {
		if containsInt(indexes, i) {
			returnStr = append(returnStr, letter)
		} else {
			returnStr = append(returnStr, string(v))
		}
	}

	return strings.Join(returnStr, "")
}

func containsStr(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func containsInt(haystack []int, needle int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}