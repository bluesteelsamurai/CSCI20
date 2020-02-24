// Lab 8 -- control flow (if/else)
//
// Programmer name: Nigel S Adams
// Date completed:  10/8/19

package main

import (
	"fmt"
	"math/rand"

	//	"os"
	//"os/exec"
	//"runtime"
	"time"
)

var playerName string = ""
var won, lost, ties, playerWep int = 0, 0, 0, 0

//Pretty Text Title
var title string = `
 _____                 _
|   __|_ _ _ ___ ___ _| |
|__   | | | | . |  _| . |
|_____|_____|___|_| |___|
 _____ _   _     _   _
|   __| |_|_|___| |_| |
|__   |   | | -_| | . |
|_____|_|_|_|___|_|___|
 _____         _ _
|   __|___ ___| | |
|__   | . | -_| | |
|_____|  _|___|_|_|
      |_|
`

// ComputerPlay returns a random integer value representing
// the computer's choice in a game of Rock, Scissors, Paper.
// 0=rock, 1=scissors, 2=paper
func ComputerPlay() int {

	return rand.Intn(3) + 1
}

// PlayerPlay prompts for, reads in, and returns
// an integer value representing the players's choice
// in a game of Rock, Scissors, Paper.
// 0=rock, 1=scissors, 2=paper
//but in this case 1=Sword 2=Shield 3=Spell
func PlayerPlay(playerName string) int {
	playerChoice := 0
	fmt.Printf("Choose your Weapon %s! \n", playerName)
	fmt.Printf("[1] Sword!\n")
	fmt.Printf("[2] Shield!\n")
	fmt.Printf("[3] Spell!\n")
	fmt.Scanln(&playerChoice)
	switch playerChoice {
	case 1:
		PrintPlay(playerName, 1)
		playerWep = 1
	case 2:
		PrintPlay(playerName, 2)
		playerWep = 2
	case 3:
		PrintPlay(playerName, 3)
		playerWep = 3
	default:
		fmt.Printf("Imbecile! You will die if you pick anything else!\n")
		PlayerPlay(playerName)

	}
	return playerChoice
}

// PrintPlay "pretty prints" the player name (human or computer)
// and what they played (rock, scissors, or paper).
// Example: "Computer plays rock"
func PrintPlay(playerName string, play int) {
	sword := `
      /
 O===[====================-
      \
 `
	shield := `
 \_________________/
 |       | |       |
 |       | |       |
 |       | |       |
 |_______| |_______|
 |_______   _______|
 |       | |       |
 |       | |       |
  \      | |      /
   \     | |     /
    \    | |    /
     \   | |   /
      \  | |  /
       \ | | /
        \| |/
         \_/
 `
	spell := `
                             /\
                            /  \
                           |    |
                         --:'''':--
                           :'_' :
                           _:"":\___
            ' '      ____.' :::     '._
           . *=====<<=)           \    :
            .  '      '-'-'\_      /'._.'
                             \====:_ ""
                            .'     \\
                           :       :
                          /   :    \
                         :   .      '.
         ,. _            :  : :      :
      '-'    ).          :__:-:__.;--'
    (        '  )        '-'   '-'
 ( -   .00.   - _
(    .'  O )     )
'-  ()_.\,\,   -
`

	//sword art!
	if play == 1 {
		fmt.Printf("%s", sword)
		fmt.Printf("\n%s chooses the sword!\n", playerName)
	}
	//shield art
	if play == 2 {
		fmt.Printf("%s", shield)
		fmt.Printf("\n%s chooses the shield!\n", playerName)

	}
	//spell art!
	if play == 3 {
		fmt.Printf("%s", spell)
		fmt.Printf("\n%s chooses the Spell!\n", playerName)
	}

}

// ShowResult "pretty prints" the result of a round
// of Rock, Scissors, Paper.
// Example: "Rock crushes scissors. Player wins."
func ShowResult(computerPlay int, humanPlay int) {
	//Player wins prettypicture
	win := `
 __ __            _ _ _
|  |  |___ _ _   | | | |___ ___
|_   _| . | | |  | | | | . |   |
  |_| |___|___|  |_____|___|_|_|
`
	//Player lost prettypicture
	lose := `
 __ __            ____  _       _
|  |  |___ _ _   |    \|_|___ _| |
|_   _| . | | |  |  |  | | -_| . |
  |_| |___|___|  |____/|_|___|___|
`

	if computerPlay == 1 && humanPlay == 2 {
		fmt.Print(win)
		fmt.Printf("You have won by the Shield\n")
		won = won + 1
	} else if computerPlay == 2 && humanPlay == 3 {
		fmt.Print(win)
		fmt.Printf("You have won by the Spell\n")
		won = won + 1
	} else if computerPlay == 3 && humanPlay == 1 {
		fmt.Print(win)
		fmt.Printf("\nYou have won by the Sword!\n")
		won = won + 1
	}
	//Computer wins
	if computerPlay == 1 && humanPlay == 3 {
		fmt.Print(lose)
		fmt.Printf("\nYou have been Slain by the Sword!\n")
		lost = lost + 1
	} else if computerPlay == 2 && humanPlay == 1 {
		fmt.Print(lose)
		fmt.Printf("\nYou have been Slain by the Shield!\n")
		lost = lost + 1
	} else if computerPlay == 3 && humanPlay == 2 {
		fmt.Print(lose)
		fmt.Printf("\nYou have been Slain by the Spell!\n")
		lost = lost + 1
	}
	//ties
	if computerPlay == humanPlay {

		fmt.Printf(" _____ _     \n")
		fmt.Printf("|_   _|_|___ \n")
		fmt.Printf("  | | | | -_|\n")
		fmt.Printf("  |_| |_|___|\n")
		fmt.Printf("             \n")
		fmt.Printf("it's a tie!\n")
		ties = ties + 1
	}
	//time for the continue to play rounds switch
	var proceed string
	fmt.Printf("%s, do you wish to try again?(Y\\N)\n", playerName)
	fmt.Scanln(&proceed)
	switch proceed {
	case "Y", "y":
		Driver(won, lost, ties)
	case "N", "n":
		break
	}
}

//Driver
////This runs the show ladies and gents
func Driver(won int, lost int, ties int) {
	C := ComputerPlay() // assign variable to function

	fmt.Printf("%s\n", title)
	fmt.Printf("WON: %d||LOSSES: %d||TIES: %d\n", won, lost, ties)

	PlayerPlay(playerName)   // calls PlayerPlay and PrintPlay for human player as choice is made.
	PrintPlay("Computer", C) // call PrintPlay for computer

	// call ShowResult
	ShowResult(C, playerWep)
}

func main() {

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("%s", title)
	fmt.Printf("You wake up in an empty room save three symbols etched into a pedestal.\n ")
	fmt.Printf("a voice booms out from the darkness\n")
	fmt.Printf("WHAT IS YOUR NAME!?\n")
	fmt.Scanln(&playerName)
	//this is the first run of driver to start everything, after this showresult opts to have it called again or exit program
	Driver(0, 0, 0)
}
