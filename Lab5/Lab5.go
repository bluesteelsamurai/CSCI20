// Lab 5 -- functions, pt 1
//
// Programmer Name: Nigel S Adams
// Date completed:  5/24/19

package main

import (
	"fmt"
	"time"
)

//Global variables, declared and intialized.
var byebye, intro, bannshow bool = false, true, false //Boolean variables that operate like switches mainly in Handler()
var Name string = "Nigel"                             //String variable for user name---ubiquitous in program
var birthyear, Age int = 1988, 31                     //integer variables for EstimateAge()
var Pearned, Ppossib float64 = 0.0, 0.0               //Float variables for percent values in ComputeAverage()

func ShowBanner(byebye bool, Name string, bannshow bool) { // ShowBanner prints a banner for the program.
	if byebye != true {
		fmt.Printf("***************************************************************************\n")
		fmt.Printf("*                ``                                            /\\         *\n")
		fmt.Printf("*                                ``         ``                /  \\        *\n")
		fmt.Printf("*          $             ahhh what a lovely day...           / .  \\       *\n")
		fmt.Printf("*         $$$           ``                  .               / |  | \\      *\n")
		fmt.Printf("*        $$$$$                                             /        \\     *\n")
		fmt.Printf("*       $$$$$$$                  ``                       /     O    \\    *\n")
		fmt.Printf("*          |                                             /            \\   *\n")
		fmt.Printf("*__________|____________________________________________/              \\  *\n")
		fmt.Printf("*#########################################################################*\n")
		fmt.Printf("*#########################################################################*\n")
		fmt.Printf("***************************************************************************\n")
		fmt.Scanln()
	}
	if byebye == true {
		fmt.Printf("*****************************************************************************\n")
		fmt.Printf("*                                                                           *\n")
		fmt.Printf("*      _______________________________________________________________      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                  This Door's for you                        |      *\n")
		fmt.Printf("*      |                  %v                                                 \n", Name)
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |     ________     _______  ________     _________     /      |      *\n")
		fmt.Printf("*      |      |_____] \\_/ |______   |_____] \\_/ |______      /       |      *\n")
		fmt.Printf("*      |      |_____]  |  |______   |_____]  |  |______     .        |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                         oo  |      *\n")
		fmt.Printf("*      |                                                     OOOOOo  |      *\n")
		fmt.Printf("*      |                                                         oo  |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      |                                                             |      *\n")
		fmt.Printf("*      _______________________________________________________________      *\n")
		fmt.Printf("*                                                                           *\n")
		fmt.Printf("*                                                                           *\n")
		fmt.Printf("*                                                                           *\n")
		fmt.Printf("*                                                                           *\n")
		fmt.Printf("*****************************************************************************\n")
		fmt.Scanln()
		if bannshow == true {
			Handler(false, false, Name)
		} //this catches the Handler() if user chooses Show me the banners!
	}
}

//Determines your age based on user supplied int values
func EstimateAge(birthyear int, Age int, Name string) {
	N := time.Now()
	fmt.Println("Please enter your 4 digit year of birth")
	fmt.Scanln(&birthyear)
	Age = N.Year() - birthyear

	fmt.Printf("You are %v, years old, %v ", Age, Name)
	Handler(false, false, Name)

}

//Determines scores based on user supplied float64 values
func ComputeAverage(Pearned float64, Ppossib float64, Name string) {
	fmt.Printf("Alight %v lets see if you passed\n", Name)
	fmt.Printf("what was your score %v?\n", Name)
	fmt.Scanln(&Pearned)
	fmt.Printf("\n what was the best score to get?\n")
	fmt.Scanln(&Ppossib)
	fmt.Printf("you scored a  %.1f percent on that whatever it was!!!\n", (Pearned/Ppossib)*100)
	if ((Pearned / Ppossib) * 100) >= 90.0 {
		fmt.Printf("Thats an A buddy!\n")
	}
	if ((Pearned/Ppossib)*100) >= 80.0 && ((Pearned/Ppossib)*100) <= 89.99 {
		fmt.Printf("Thats a B guy!\n")
	}
	if ((Pearned/Ppossib)*100) >= 70.0 && ((Pearned/Ppossib)*100) <= 79.99 {
		fmt.Printf("Thats a C pal!\n")
	}
	if ((Pearned/Ppossib)*100) >= 60.0 && ((Pearned/Ppossib)*100) <= 69.99 {
		fmt.Printf("Thats a D friend!\n")
	}
	if ((Pearned/Ppossib)*100) >= 50.0 && ((Pearned/Ppossib)*100) <= 59.99 || ((Pearned/Ppossib)*100) < 50.0 {
		fmt.Printf("Thats a F buddy!\n")
	}
	fmt.Println("Press ENTER to return to Main Menu.\n")
	fmt.Scanln()
	Handler(false, false, Name)
}

//GreetUser and SayGoodbye rolled into one
func Handler(intro bool, byebye bool, Name string) {
	Handmen := 0

	if intro == true {
		ShowBanner(false, Name, false)
		fmt.Printf("Hello and welcome to the B.A.G. program my friend!\n")
		fmt.Printf("what might be your name?\nPlease enter your name and press ENTER:")
		fmt.Scanln(&Name)
		intro = false
	}
	if intro != true {

		fmt.Printf("\nPlease make a selection \n")
		fmt.Printf("1) Show me the Banners\n")
		fmt.Printf("2) Guess my Age\n")
		fmt.Printf("3) Tell me what my Grade is\n")
		fmt.Printf("4)Show me the door!\n\n\nPlease enter your choice(1-4) and press ENTER:")
		fmt.Scanln(&Handmen)
		if Handmen == 1 {
			ShowBanner(false, Name, false)
			ShowBanner(true, Name, true)
			byebye = false
		}
		if Handmen == 2 {
			EstimateAge(birthyear, Age, Name)
		}
		if Handmen == 3 {
			ComputeAverage(Pearned, Ppossib, Name)
		}
		if Handmen == 4 {
			ShowBanner(true, Name, false)
		} else {
			Handler(false, false, Name)
		}
	}

}

func main() {

	Handler(true, false, "") //this starts everything into motion in a "default" setting

	/*

	   This section helps me keep track of my lab objectives...utterly superfluous
	   	++// GreetUser accepts a Name and prints a greeting message
	   	// using the Name.

	   	++// EstimateAge accepts a birth year and computes and
	   	// prints user age in years.

	   	// ComputeAverage accepts points earned and points possible,
	   	// computes and prints the average.

	   	++// SayGoodbye accepts a Name and prints a goodbye message
	   	// using the Name.

	   	++// call ShowBanner

	   	++// prompt for and read in user Name
	   	// call GreetUser -- pass the user Name

	   	++// prompt for and read in user birth year
	   	// call EstimateAge -- pass user birth year

	   	++// prompt for and read in points earned
	   	++// prompt for and read in points possible
	   	++// call ComputeAverage -- pass points earned and points possible

	   	++// call SayGoodbye -- pass the user Name
	*/
}
