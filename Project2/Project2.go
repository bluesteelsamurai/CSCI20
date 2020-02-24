/*
Creator: Nigel S Adams
Date started:10-17-19
Date Completed:more like the day I stopped working...I could build on this forever if I had the time, incidentally its 10/30/2019
info:Project #2 is a project in which a random number guessing game and a rock paper scissors game are run via a Driver function, the goal was to capitalize on for loops, randomizing numbers with the nanosecond at the moment as the key, and showcasing the skills learned from previous labs to their fullest such as pretty-printing which is throughout the program

*/
package main

import (
	"fmt"       //format package helps print to screen and gives programmer options for how information is poresented
	"math/rand" //math/rand package is how random numbers and other mathmatical processes are possible such as sqrt which means find the square root of something.
	"os"
	"time" /*time package is for precisely that, time related behaviors that the programmer can use to add unique functions like one to properly randomize the randomizer based off the time down to the nanosecond as the key for that random selection when called*/
)

var playerName string = ""                              //global playerName variable that is better declared here than passed from function to function
var print, pChoice, jerkUser, gameType int = 0, 0, 0, 0 // global print is the variable used for Prettyprints parameter,global pChoice is a universal integer variable for navigating menus, global jerkUser is a variable that holds how many times the user inputs the wrong value type or beyond the range of available values, that value is then used by Driver to dispense the appropriate consequence to the user for not reading carefully.

//PrettyPrints prints out my text art for the games and title screen and menu...just all the shit that would take forever to manually code into each function
func PrettyPrints(print int) {
	/*this section is going to define all my big pretty printing and follow with a simple switch that chooses the particular graphic by its integer number associated with it.
	the images come out distorted without modifying them here piece by piece so they look weird here but they look great in use.*/

	//this graphic and a few other presented weird so adjustments were made to achieve correct appearance
	var pridePic string = `
	 _____ _____ _____    _____
	|   | |   __|  _  |  |   __|___ _____ ___ ___
	| | | |__   |     |  |  |  | .'|     | -_|_ -|
	|_|___|_____|__|__|  |_____|__,|_|_|_|___|___|
	         _____                     _
		|  _  |___ ___ ___ ___ ___| |_ ___
		|   __|  _| -_|_ -| -_|   |  _|_ -|
		|__|  |_| |___|___|___|_|_|_| |___|
`
	var titlePic string = `
	 _____           _         _
	|  _  |___ ___  |_|___ ___| |_
	|   __|  _| . | | | -_|  _|  _|
	|__|  |_| |___|_| |___|___|_|
		      |___|
	  	    _ _
		  _| | |_    ___
	         |_     _|  |_  |
		 |_     _|  |  _|
		   |_|_|    |___|
`

	var spellPic string = `
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
	//
	var swordPic string = `
     /
O===[====================-
     \
`

	var shieldPic string = `
	\_              _/
	] --__________-- [
	|       ||       |
	\       ||       /
	 [      ||      ]
	 |______||______|
	 |------..------|
	 ]      ||      [
	  \     ||     /
	   [    ||    ]
	   \    ||    /
	    [   ||   ]
	     \__||__/
	        --
	`
	var winPic string = `
	 __ __            _ _ _         __
	|  |  |___ _ _   | | | |___ ___|  |
	|_   _| . | | |  | | | | . |   |__|
	  |_| |___|___|  |_____|___|_|_|__|
	`

	var losePic string = `
	 __ __            ____  _       _
	|  |  |___ _ _   |    \|_|___ _| |
	|_   _| . | | |  |  |  | | -_| . |
	  |_| |___|___|  |____/|_|___|___|
	`
	var tiePic string = `
	 _____ _
	|_   _|_|___
	  | | | | -_|
	  |_| |_|___|
	`
	var sssPic string = `
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
	switch print {
	case 1:
		fmt.Printf("%s\n", pridePic)
	case 2:
		fmt.Printf("%s\n", titlePic)
	case 3:
		fmt.Printf("%s\n", spellPic)
	case 4:
		fmt.Printf("%s\n", swordPic)
	case 5:
		fmt.Printf("%s\n", shieldPic)
	case 6:
		fmt.Printf("%s\n", winPic)
	case 7:
		fmt.Printf("%s\n", losePic)
	case 8:
		fmt.Printf("%s\n", sssPic)
	case 9:
		fmt.Printf("%s\n", tiePic)

	}
}

/*WhatNum is the number guessing game with a random number decided from a range defined by the player and they decide how many attempts they wanna make afte rwhich they must guess or lose */
func WhatNum(playerName string) {
	pAnswer := ""                              /*function specific variable for handling user response...going with detecting simple string YyNn char-like responses from user to keep things simple*/
	theNumb := 0                               /*function specific variable which takes input from user for range of numbers to make a random choice for guessing*/
	theRange := 0                              // variable used for displaying the range user guesses from
	pGuess, pAttempts, pMaxAttempts := 0, 0, 0 //player associated variables whih are kinda self explanatory I hope

	fmt.Printf("This will be a number guessing game with some freedom on rules ...\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("You get to decide how many tries to guess the number.\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("You also choose the size of the range of numbers [1-infinity,etc]\nit's kind of a free for all here just make it as hard or easy as you want.\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("are you ready to play %v? \n[Y/N]:", playerName)
	fmt.Scanln(&pAnswer)
	if pAnswer == "y" || pAnswer == "Y" {
		fmt.Printf("Excellent, lets get started.\n")
		fmt.Printf("%v, please pick a number to start the range from.\n", playerName)
		fmt.Printf("[please note that higher numbers will increase difficulty]\n")
		fmt.Scanln(&theNumb)
		theRange = theNumb
		theNumb = NumRand(theNumb)
		fmt.Printf("Now Choose how many rounds you wanna play\n")
		fmt.Printf("[note that too many rounds will make it too easy and kinda boring then]\n")
		fmt.Scanln(&pMaxAttempts)
	}
	if pAnswer == "N" || pAnswer == "n" {
		fmt.Printf("Returning to Main Menu\n")
		time.Sleep(2 * time.Second)
		Driver()
	}
	fmt.Printf("All set! I'm computing a number between 1 and %v!\nmake your guess!\n", theRange)

	for {
		fmt.Printf("Choose a # between [1-%v]\n", theRange)
		fmt.Scanln(&pGuess)
		if pAttempts <= pMaxAttempts {
			if pGuess < theNumb {
				pAttempts++
				fmt.Printf("tis higher I'm afraid! you have %v left\n", pMaxAttempts-pAttempts)
			}
			if pGuess > theNumb {
				pAttempts++
				fmt.Printf("tis lower I'm afraid! you have %v left\n", pMaxAttempts-pAttempts)
			}
			if pGuess == theNumb {
				PrettyPrints(6)

				fmt.Printf("tis exactly so! you guessed in %v turns\n", pAttempts+1)
				Retry(0, 0, 0, 2)
			}
		} else {
			PrettyPrints(7)
			fmt.Printf("you didn't guess in time...\n")
			Retry(0, 0, 0, 2)
		}
	}
}

/*NumRand is the function which randomly selects the value that you guess for in the number guessing game should be a simple call to get a random number returned to be used in each playthrough */
func NumRand(pRange int) int {
	return rand.Intn(pRange) + 1
}

/*SsS is Sword, Shield, Spell; a rock paper scissors game that has three choices, sword shield and spell and they work in this order sword>spell>shield>sword or sword beats spell, spell beats shield and shield beats sword. This version has a predictive opponent that uses data from your continued playthroughs to try and not use moves that it lost with previously*/
func SsS(playerName string, difficulty int) {
	pChoice = 0
	wins, loss, tie := 0, 0, 0
	PrettyPrints(8)

	fmt.Printf("BEHOLD! YOU HAVE ENTERED THE FIGHT PIT OF TRIPLE S!\n CHOOSE YOUR WEAPON MORTAL OR DIE A DOG'S DEATH AT THE HANDS OF MOGTAR THE OGRE!\n")
	SssWeapon(wins, loss, tie)

}

/*SssWeapon drives the process of choosing a weapon in the SSS function */
func SssWeapon(wins int, loss int, tie int) {
	fmt.Printf("WINS:%d LOSS:%d TIES:%d\n\n", wins, loss, tie)
	fmt.Printf("Before you lies a table with three items:\n[1]a gleaming Sword honed to deadly sharpness.\n")
	fmt.Printf("[2]a Tower Shield with a Crucifix emblazoned on its face \n")
	fmt.Printf("[3]a Wand and parchment that glows ominously\n")

	fmt.Scanln(&pChoice)
	switch pChoice {
	case 1:
		PrettyPrints(4)
		fmt.Printf("You have chosen the Sword!\nPress [Enter] to continue.\n")
		fmt.Scanln()
		BattleResult(1, MogTar(), wins, loss, tie)
	case 2:
		PrettyPrints(5)
		fmt.Printf("You have chosen the Shield!\nPress [Enter] to continue.\n")
		fmt.Scanln()
		BattleResult(2, MogTar(), wins, loss, tie)
	case 3:
		PrettyPrints(3)
		fmt.Printf("You have chosen the Spell!\nPress [Enter] to continue.\n")
		fmt.Scanln()
		BattleResult(3, MogTar(), wins, loss, tie)
	default:
		PrettyPrints(7)
		fmt.Printf("You die a dog's death for not choosing...you mangy cur\n")
		fmt.Scanln()
		loss++
		Retry(wins, loss, tie, 1)

	}
}

/*BattleResult is the results of the users choice vs the pc randomly choosing 1-3, the wins,loss,tie parameters are passed through this in order to increment them effectively.*/
func BattleResult(pChoice int, MogTar int, wins int, loss int, tie int) {

	if pChoice == 1 && MogTar == 2 {
		PrettyPrints(7)
		fmt.Printf("MogTar knocks you down as your sword lodges itself uselessly in the edge of his shield and he \nsubsequently beats you to death with his shield.\n")
		fmt.Printf("Press [Enter] to continue.\n")
		fmt.Scanln()
		loss++
		Retry(wins, loss, tie, 1)
	}
	if pChoice == 2 && MogTar == 3 {
		PrettyPrints(7)
		fmt.Printf("MogTar points the wand as you bravely charge him with your shield only to see him finish the spell\n and turn you into a frog which he gleefuly squishes.\n")
		fmt.Printf("Press [Enter] to continue.\n")
		fmt.Scanln()
		loss++
		Retry(wins, loss, tie, 1)
	}
	if pChoice == 3 && MogTar == 1 {
		PrettyPrints(7)
		fmt.Printf("MogTar rushes you with his sword arm drawn back to strike while you struggle to aim the spell,\nclosing the distance MogTar plants the sword firmly in your chest.\n")
		fmt.Printf("Press [Enter] to continue.\n")
		fmt.Scanln()
		loss++
		Retry(wins, loss, tie, 1)
	}
	if pChoice == 1 && MogTar == 3 {
		PrettyPrints(6)
		fmt.Printf("Surging forward with a predators grace you rush MogTar with sword in hand and turn him into a grotesque\nsheathe for the deadly blade.\n")
		fmt.Printf("Press [Enter] to continue.\n")
		fmt.Scanln()
		wins++
		Retry(wins, loss, tie, 1)
	}
	if pChoice == 2 && MogTar == 1 {
		PrettyPrints(6)
		fmt.Printf("With the might of a Thousand Bulls you batter MogTar to the ground with your shield as his sword lodges itself\nimplacably so in your shield leaving you to freely rearrange his face.\n")
		fmt.Printf("Press [Enter] to continue.\n")
		fmt.Scanln()
		wins++
		Retry(wins, loss, tie, 1)
	}
	if pChoice == 3 && MogTar == 2 {
		PrettyPrints(6)
		fmt.Printf("MogTar boldly charges you with his shield but your spell is faster, by the time he reaches you he \nis no longer MogTar the Ogre but MogTar the frog...frog-legs anyone?\n")
		fmt.Printf("Press [Enter] to continue.\n")
		fmt.Scanln()
		wins++
		Retry(wins, loss, tie, 1)
	}
	if pChoice == MogTar {
		PrettyPrints(9)
		fmt.Printf("With your respective weapons being evenly matched MogTar and you clash to no avail with no clear winner...quite the show still\n")
		tie++
		Retry(wins, loss, tie, 1)
	}
}

//MogTar function is the computers random choice between 1-3 based off random and time.now().unix nano() as the key to randomize by.
func MogTar() int {
	return rand.Intn(3) + 1
}

/*Retry is a function that handles the end of game routine, like play again or exit*/
func Retry(wins int, loss int, tie int, gameType int) {

	again := "" //user input variable for handling response to query to Retry

	fmt.Printf("Try again?[y/n]\n")
	fmt.Scanln(&again)
	if again == "Y" || again == "y" {
		if gameType == 1 {
			SssWeapon(wins, loss, tie)
		}
		if gameType == 2 {
			WhatNum(playerName)
		}
	}
	if again == "N" || again == "n" {
		Driver()
	} else {
		fmt.Printf("I'm sorry I don't understand ('Y','y'/'N','n') please ")
		Retry(wins, loss, tie, gameType)
	}
}

/*DriverMenu is a function that is called from Driver that handles user response in tandem with Driver to avoid any superfluous text entry like when they give a string instead of an int etc by simply reiterating the menu and instructing the user of accepted entries...it's the idiot catcher if you will.*/
func DriverMenu(playerName string) int {
	fmt.Printf("%s, Please make a selection;\n", playerName)
	fmt.Printf("[1]Game of Chance\n[2]Game of Guessing\n[3]Get me outta here!\n#[1-3]:")
	fmt.Scanln(&pChoice)
	return pChoice
}

/*Driver runs the show so to speak, its started in main with a bool parameter that acts like a on/off switch for the program. It also has this neat feature where if the user just wants to see what happens if they keep trying to input the wrong data it just closes after admonishing the user...just kidding it goes into a repeating iteration of a phrase one hundred million times...don't do it if your pc is slower than average*/
func Driver() {
	for {
		DriverMenu(playerName)
		switch pChoice {
		case 1:
			SsS(playerName, 1)
		case 2:
			WhatNum(playerName)
		case 3:
			fmt.Printf("Bye bye now have a good one!\n")
			for i := 4; i > 0; i-- {
				time.Sleep(1 * time.Second)
				fmt.Printf("%d\n", i-1)
			}
			os.Exit(1)

		default:
			if jerkUser == 2 {
				fmt.Printf("Are you trying to piss me off?!\n")
				jerkUser++
				Driver()
			}
			if jerkUser == 4 {
				fmt.Printf("This is your last warning before I unleash holy hell on your computer.\n")
				jerkUser++
				Driver()
			}
			if jerkUser == 6 {
				for k := 0; k < 999999999; k++ {
					fmt.Printf("You dun goofed BRAH!")
				}
			}
			fmt.Printf("\nuh what?...lets try that again...*gives side-eyed look*\n\n")
			jerkUser++
			Driver()

		}
	}
}
func main() {

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
	PrettyPrints(1) //Prints my textart for NSA GAMES
	PrettyPrints(2) //Title text art
	fmt.Printf("Greetings and Salutations Friend!\n")
	fmt.Printf("How may I address you from here on?\n")
	fmt.Printf("Please type your name below\n")
	fmt.Scanln(&playerName)

	Driver() //This starts the driver which leads to the games or exit

}
