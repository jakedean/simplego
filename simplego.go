package main

import(
	"github.com/jakedean/simplego/hangman"
	"github.com/jakedean/simplego/rotator"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must choose a package to run.")
	} else {
		packageToRun := os.Args[1]
		switch packageToRun {
		case "hangman":
			hangman.Play()
		case "rotator":
			rotator.Begin()
		default:
			fmt.Println("You must chose a package to run.")
		}
	}
}