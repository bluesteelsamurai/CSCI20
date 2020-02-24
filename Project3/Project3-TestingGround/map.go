package main

/*
##What is it? a toolkit that can be used to make
an interactive text adventure with a working textfile imported
that translates a sheet of integers into a graphic represemtation.
##Who made it? Patrick Woodward & Nigel Adams
##When was it made between 11/12-11/27 2019
*/
//Imports
import (
	bio "bufio" //using placeholder "bio" for bufio cuz i like to complicate yet shorthand everything
	f "fmt"     //using placeholder "f" for fmt to write two less letters per println
	"log"       //imported to record any errors that may occur during runtime
	"os"        //imported to read txt file as well as become method to exit program at end
	"os/exec"   //imported for screen clears and haxy stuff...jk
	"runtime"   //imported to determine user OS for screenwipe function
	"strings"   //to be used later for stringy...stuff
	"time"      //imported for flair and making people wait
)

//Globals
//leave is a boolean used to keep driver loop going
var leave bool

//XaXis is the variable to track x-axis length for pathfinding
var XaXis int

//YaXis is the variable to track Y-axis length for pathfinding
var YaXis int

//ChartXY map that works like xy mapping and is declared globally
var ChartXY = make([][]int, 40)

//OrigChart is the Chart used to replace the locatoins the player goes through so its not a blank spot
var OrigChart = make([][]int, len(ChartXY))

//ChartKey makes the map display correctly delimited to 40 lines.
var ChartKey = make([]int, 40)

//wipe is a string map of funcs that execute cls's or clear screen commands from various OS's
var wipe map[string]func()

//Variables used for making actions
var open int
var pickup int
var delete int
var equip int

//Slices and structs used for creating inventory, monsters, player stats, and all equipment pieces
var bag []string
var equipment []string
var allArmor []armor
var allWeapons []weaps
var allMonsters []monster
var baseStats []base

type armor struct {
	name string
	stat int
}

type weaps struct { //all structs for inventory, equipment, and stat purposes
	name string
	stat int
}

type base struct {
	name string
	stat int
}

type monster struct {
	name    string
	health  int
	defense int
	attack  int
}

//Function use to make equipment and player stats occur
func makeEquipment() {
	allArmor = []armor{
		armor{name: "clothe gloves", stat: 10},
		armor{name: "cloth helm", stat: 10},
		armor{name: "cloth curais", stat: 10},
		armor{name: "boots", stat: 10},
		armor{name: "shield", stat: 10},
	}
	allWeapons = []weaps{
		weaps{name: "sword", stat: 10},
		weaps{name: "axe", stat: 10},
		weaps{name: "mace", stat: 10},
	}
	baseStats = []base{
		base{name: "health", stat: 10},
		base{name: "attack", stat: 10},
		base{name: "defense", stat: 10},
	}
}

//Function used to make monsters occur
func makeMonster() {
	allMonsters = []monster{
		monster{name: "weak monster", health: 15, defense: 10, attack: 5},
		monster{name: "monster", health: 30, defense: 15, attack: 7},
		monster{name: "strong monster", health: 35, defense: 22, attack: 14},
		monster{name: "bossmonster", health: 50, defense: 30, attack: 23},
	}
}

func attaining() {
	for i := 0; i < 10; i++ {
		f.Scanln(&pickup)
		if pickup == 1 {
			f.Printf("You put %s into your bag of holding.\n")
			bag = append(bag, ___)
			break
		} else if pickup == 2 {
			f.Printf("You leave the item where you found it.\n")
			break
		} else {
			f.Printf("Take it or leave it, make a real choice!!\n")
		}
	}
} //change item to functional items from structs

func equipped() {
	f.Printf("What do you want to equip?\n")
	f.Printf("You have %v in your inventory.\n", bag)
	f.Scanln(&equip)
	if equip == 1 {
		equipment = append(equipment, bag[equip-1])
		bag = append(bag[:equip-1], bag[equip:]...)
	} else if equip == 2 {
		equipment = append(equipment, bag[equip-1])
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if equip == 3 {
		equipment = append(equipment, bag[equip-1])
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if equip == 4 {
		equipment = append(equipment, bag[equip-1])
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if equip == 5 {
		equipment = append(equipment, bag[equip-1])
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if equip == 6 {
		equipment = append(equipment, bag[equip-1])
		bag = append(bag[:delete-1], bag[delete:]...)
	}
}

func discard() {
	f.Printf("What item do you want to discard? 1-10\n")
	for i, bag := range bag {
		f.Printf("%d) %s\n", i+1, bag)
	}
	f.Scanln(&delete)
	if delete == 1 {
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if delete == 2 {
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if delete == 3 {
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if delete == 4 {
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if delete == 5 {
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if delete == 6 {
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if delete == 7 {
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if delete == 8 {
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if delete == 9 {
		bag = append(bag[:delete-1], bag[delete:]...)
	} else if delete == 10 {
		bag = append(bag[:delete-1], bag[delete:]...)
	} else {
		f.Printf("Your bag is not that big.\n")
	}
	f.Printf("Your bag now contains\n")
	for i, bag := range bag {
		f.Printf("%d) %v\n", i+1, bag)
	}
}

//Functions

func chest1() {
	f.Println("You find a chest, do you want to open the chest?\n1)Yes\n2)No\n")
	f.Scanln(&open)
	if open == 1 {
		f.Printf("You open the chest and find %v!!\n", allArmor[1].name)
		bag = append(bag, allArmor[1].name)
		bag = append(bag[:open-1], bag[open:]...)
	} else if open == 2 {
		f.Printf("You leave the chest unopened and may never know what is in it.\n")
	} else {
		f.Printf("Please make a valid choice.")
	}
}

func chest2() {
	f.Println("You find a chest, do you want to open the chest?\n1)Yes\n2)No\n")
	f.Scanln(&open)
	if open == 1 {
		f.Printf("You open the chest and find %v!!\n", allArmor[4].name)
		bag = append(bag, allArmor[4].name)
		bag = append(bag[:open-1], bag[open:]...)
	} else if open == 2 {
		f.Printf("You leave the chest unopened and may never know what is in it.\n")
	} else {
		f.Printf("Please make a valid choice.")
	}
}

func chest3() {
	f.Println("You find a chest, do you want to open the chest?\n1)Yes\n2)No\n")
	f.Scanln(&open)
	if open == 1 {
		f.Printf("You open the chest and find %v!!\n", allArmor[0].name)
		bag = append(bag, allArmor[0].name)
		bag = append(bag[:open-1], bag[open:]...)
	} else if open == 2 {
		f.Printf("You leave the chest unopened and may never know what is in it.\n")
	} else {
		f.Printf("Please make a valid choice.")
	}
}

func chest4() {
	f.Println("You find a chest, do you want to open the chest?\n1)Yes\n2)No\n")
	f.Scanln(&open)
	if open == 1 {
		f.Printf("You open the chest and find %v!!\n", allArmor[2])
		bag = append(bag, allArmor[2].name)
		bag = append(bag[:open-1], bag[open:]...)
	} else if open == 2 {
		f.Printf("You leave the chest unopened and may never know what is in it.\n")
	} else {
		f.Printf("Please make a valid choice.")
	}
}

func chest5() {
	f.Println("You find a chest, do you want to open the chest?\n1)Yes\n2)No\n")
	f.Scanln(&open)
	if open == 1 {
		f.Printf("You open the chest and find %v!!\n", allArmor[3])
		bag = append(bag, allArmor[3].name)
		bag = append(bag[:open-1], bag[open:]...)
	} else if open == 2 {
		f.Printf("You leave the chest unopened and may never know what is in it.\n")
	} else {
		f.Printf("Please make a valid choice.")
	}
}

/*ReadFromFile takes any sized grid of text in a text file and puts it into an array to be used as a Chart*/
func ReadFromFile(path string, ChartXY [][]int) {
	//slice to read file into
	ChartY := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("uhhh, something broke. could be your wording or %s is not there.", err)
	} else {
		Fscanner := bio.NewScanner(file)
		Fscanner.Split(bio.ScanLines)
		defer file.Close()
		//Fscanner reads data into a
		//range for loop that reads file till error or EoF(End of File)into
		i := 0
		for Fscanner.Scan() {
			ChartY = append(ChartY, strings.Trim(Fscanner.Text(), "\u00A0"))
			i++
			XaXis = len(strings.Trim(Fscanner.Text(), "\u00A0"))
			YaXis = i
		}
		//adding second level to ChartXY
		for i := range ChartXY {
			ChartXY[i] = make([]int, len(ChartY[i]))
			for j := range ChartXY[i] {
				ChartXY[i][j] = 0
			}
		}
		//splitting the lines of ChartY into individual chars for ChartXY as elements
		for i, obj := range ChartY {
			for j := 0; j < len(obj)-1; j++ {
				ChartXY[i][j] = int(obj[j])
			}
		}
	}
}

/*drawChart draws the map of the region the game is played in, its referencing the
global 2 dimensional slice populated by readfromfile and then uses a nested for loop to print out the map iteratively, interpreting numbers as strings to then print out based on the format of the supplied txt file*/
func drawChart(ChartXY [][]int) {
	//this nested for range loop checks each element pair from the Chart slice and its assigned value then prints out the value's associated icon to the Chart!
	for i, xis := range ChartXY {
		for j := range xis {
			if string(ChartXY[i][j]) == "0" {
				f.Print("░") //grass
			} else if string(ChartXY[i][j]) == "1" {
				f.Print("▓") //path
			} else if string(ChartXY[i][j]) == "2" {
				f.Print("\u25B2") //Player
			} else if string(ChartXY[i][j]) == "3" {
				f.Print("?") //Monster
			} else if string(ChartXY[i][j]) == "4" {
				f.Print("T") //Treasure
			} else if string(ChartXY[i][j]) == "5" {
				f.Print("║") //wall-vertical
			} else if string(ChartXY[i][j]) == "6" {
				f.Print("═") //wall-horizontal
			} else if string(ChartXY[i][j]) == "7" {
				f.Print("\\") //slope wall right
			} else if string(ChartXY[i][j]) == "8" {
				f.Print("/") //slope wall left
			} else if string(ChartXY[i][j]) == "9" {
				f.Print("D") //Door
			}
		}
		//marks map with values for navigation...remove before release
		// f.Printf("║#%v", YaXis-i)
		//pushes onto next row
		f.Print("\n")
	}
	//a for loop to draw a series of nums for pathfinding
	// for i := 5; i < XaXis+5; {
	// f.Printf("  \\%v", i)
	// i = i + 5
	// }
	//finishes off map and newlines for non-map stuff
	f.Print("\n")
}

// TGB is a function for Traversing Game Board hence TGB for ease of use
// TGB will take user input based on a compass rose so North,South,East,West,
//NorthEast,NorthWest,SouthEast,SouthWest.
func TGB(ChartXY [][]int, OrigChart [][]int, pDirect string, lstPosVal string, leave bool) bool {
	// GetX,GetY,SetX,SetY are variables for getting and setting Playerlocation on ChartXY
	var SetX, SetY, DestXY, PrevXY int

	var trigger bool
	for trigger != true {
		// for loop to get player location by iterating through ChartXY
		for i, yAxis := range ChartXY {
			for j := range yAxis {
				if ChartXY[i][j] == int('2') {
					SetX = j
					SetY = i
					if ChartXY[i][j] == int('2') && OrigChart[i][j] == int('2') {
						OrigChart[i][j] = int('1')
					}
					PrevXY = OrigChart[i][j]
					trigger = true
				} else if i == len(ChartXY) {
					log.Fatalf("Player could not be found!")
				}
			}
		}
	}

	//needs more work somethings not right here, places get swapped what if each move destroys origin spot?
	//either that or additional variable that records intended location, saves its value then replaces it after player leaves.
	switch pDirect {
	case "N": //focus on SetY in negative value cuz map is written down to up
		DestXY = ChartXY[SetY-1][SetX]
		if trigChk(DestXY) == false {
			//fight or door or treasure
		}
		ChartXY[SetY-1][SetX] = ChartXY[SetY][SetX]
		ChartXY[SetY][SetX] = PrevXY
	case "S": //focus on SetY in positive value because ...
		DestXY = ChartXY[SetY+1][SetX]
		if trigChk(DestXY) == false {
			//fight or door or treasure
		}
		ChartXY[SetY+1][SetX] = ChartXY[SetY][SetX]
		ChartXY[SetY][SetX] = PrevXY
	case "W": //focus on SetX in negative value because map is written left to right.
		DestXY = ChartXY[SetY][SetX-1]
		if trigChk(DestXY) == false {
			//fight or door or treasure
		}
		ChartXY[SetY][SetX-1] = ChartXY[SetY][SetX]
		ChartXY[SetY][SetX] = PrevXY
	case "E": //focus on SetX in positive value because...
		DestXY = ChartXY[SetY][SetX+1]
		if trigChk(DestXY) == false {
			//fight or door or treasure
		}
		ChartXY[SetY][SetX+1] = ChartXY[SetY][SetX]
		ChartXY[SetY][SetX] = PrevXY
	case "NW": //focus on SetY as negative and SetX as negative because its north and west combined
		DestXY = ChartXY[SetY-1][SetX-1]
		if trigChk(DestXY) == false {
			//fight or door or treasure
		}
		ChartXY[SetY-1][SetX-1] = ChartXY[SetY][SetX]
		ChartXY[SetY][SetX] = PrevXY
	case "NE": //focus on SetY as ... and SetX as positive because East is postive
		DestXY = ChartXY[SetY-1][SetX+1]
		if trigChk(DestXY) == false {
			//fight or door or treasure
		}
		ChartXY[SetY-1][SetX+1] = ChartXY[SetY][SetX]
		ChartXY[SetY][SetX] = PrevXY
	case "SW": //focus on SetY as Positive...SetX as Negative...
		DestXY = ChartXY[SetY+1][SetX]
		if trigChk(DestXY) == false {
			//fight or door or treasure
		}
		ChartXY[SetY-1][SetX] = ChartXY[SetY][SetX]
		ChartXY[SetY][SetX] = PrevXY
	case "SE": //Focus on SetY as Positive...SetX as Positive
		DestXY = ChartXY[SetY+1][SetX+1]
		if trigChk(DestXY) == false {
			//fight or door or treasure
		}
		ChartXY[SetY+1][SetX+1] = ChartXY[SetY][SetX]
		ChartXY[SetY][SetX] = PrevXY
	case "QUIT":
		leave = true
		break
	}
	drawChart(ChartXY)
	return leave
}

//looks at destination of TGB input and checks if its a path or something to contend with
func trigChk(DestXY int) bool {
	if DestXY != 1 {
		return true
	}

	return false

}

//MakeOriginalChart is a function which is used to create a comparative slice for moving across Game map without leaving blank spaces
func MakeOriginalChart(ChartXY [][]int, OrigChart [][]int) {
	for i := range ChartXY {
		OrigChart[i] = make([]int, len(ChartXY[i]))
		copy(OrigChart[i], ChartXY[i])
	}
}

//init initializes string func map of OSes for SCR_Wipe()
func init() {
	wipe = make(map[string]func())
	wipe["windows"] = func() {
		commnd := exec.Command("cmd", "/c", "cls")
		commnd.Stdout = os.Stdout
		commnd.Run()
	}
	wipe["darwin"] = func() {
		commnd := exec.Command("clear")
		commnd.Stdout = os.Stdout
		commnd.Run()
	}
	wipe["linux"] = func() {
		commnd := exec.Command("clear")
		commnd.Stdout = os.Stdout
		commnd.Run()
	}
}

//SCRWipe wipes the screen clean of all that text debris...if it recognizes the OS
func SCRWipe() {
	//
	cmd, allGreen := wipe[runtime.GOOS]
	if allGreen {
		cmd()
	} else {
		panic("unrecognized platform, run program on Windows OS,MacOS, or Linux OS")
	}
}

//using a file scanner to read lines from txt file to save space in source code.
//also cuz this makes it possible for others to write their own lore etc while still-
//maintaining the mechanics etc.
// func story(){
// 	plot:=[]string
//
// 	switch plot {
//
// }
// }

//Driver runs the show, takes a string for direction to travel including QUIT
//also uses the FileName string to import txt file just type in name no .txt necessary
func Driver(pDirect string, FileName string) {
	lstPosVal := ""
	f.Printf("Welcome to the Tgrounds!\n")
	time.Sleep(3 * time.Second)
	f.Print("Please type a filename,no .txt necessary\n=>:")
	f.Scanln(&FileName)
	f.Print("\nThank you, now loading...\n")
	SCRWipe()
	FileName = FileName + ".txt"
	ReadFromFile(FileName, ChartXY)
	MakeOriginalChart(ChartXY, OrigChart)
	drawChart(ChartXY)
	for leave == false {

		f.Printf("choose a direction!\n")
		f.Printf("N,S,W,E,NW,NE,SW,SE\nor press enter to continue previous command\n or type Quit=>:")
		f.Scanln(&pDirect)
		pDirect = strings.ToUpper(pDirect)
		if pDirect == "QUIT" {
			break
		} else {
			SCRWipe()
			TGB(ChartXY, OrigChart, pDirect, lstPosVal, leave)
		}
	}

}
func main() {
	// creates armor, weapon, and stats for use
	makeEquipment()
	// creates all monsters
	makeMonster()

	//declare FileName for opening txt file
	var FileName string
	pDirect := "" //ued to grab input from user for use in Driver and TGB
	Driver(pDirect, FileName)

}

/*SOLUTION!!!!!*/
/* PATRICK SAID*/
/*
tah tah for now haha
for sure man, i will poke around at things tomorrow and see what I can do
for some polishing touches on my code and I will also see if I can find
anyting funky with the map stuffs
I will drink grog from their skulls!!!
g'night
*/
/*NIGEL SAID*/
/*
righto, tah
i will upload what i got so far for the movement and such to the tgorund.go file
try and see whats up with the map.go file as well
do itttt look for any weaknesses!
crush the enemies, see them driven before you and hear the lamentations of their women!
lol goodnight

the memories are part of the story driving part, if we move that too fast
the player might get to the end before they even hit a high enough level
for a boss fight
yeah
control the story, control the world!
mwahahahaha
im a bit of a maniac and a malicious one at that
*cackles in goblin* keepin dis
weird, i can't copy the imports from tground
*/
