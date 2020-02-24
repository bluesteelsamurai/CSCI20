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
	"math/rand" //random tool for randomness like if you win a bout or get sweet sweet loot
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

//boolean slices
var AntiBoydcheeser = make([][]bool, len(ChartXY))

//wipe is a string map of funcs that execute cls's or clear screen commands from various OS's
var wipe map[string]func()

//Variables used for making actions
var open int
var pickup int
var Delete string
var equip string
var pDirect string  //pDirect is used as input string for TGB function to execute
var FileName string //FileName for opening txt file of map
var ABC bool        //AntiBoydcheeser acronym-for check
//Memories is a variable that is used for plot advancement, like chapters one can advance story,
//create an opportunity to level up, etc
var Memories int

//Slices and structs used for creating inventory, monsters, player stats, and all equipment pieces
var bag = make(map[string]int, 10)
var equipment []string
var allgear []gear
var allMonsters []monster
var P Player
var plot []string

// GetX,GetY,SetX,SetY are variables for getting and setting Playerlocation on ChartXY
var SetX, SetY, DestXY, PrevXY int

type gear struct {
	name    string
	attack  int
	defense int
}
type monster struct {
	name    string
	health  int
	defense int
	attack  int
	XP      float64
}

//Player is a struct holding the staitstics of the protagonist
type Player struct {
	name      string
	level     float64
	health    int
	maxHealth int
	defense   int
	attack    int
}

//Function use to make equipment and player stats occur
func makeEquipment() {
	allgear = []gear{
		/*1*/ gear{name: "Cloth Gloves", attack: 0, defense: 10},
		/*2*/ gear{name: "Cloth Helm", attack: 0, defense: 10},
		/*3*/ gear{name: "Cloth Cuirass", attack: 0, defense: 10},
		/*4*/ gear{name: "Cloth Boots", attack: 0, defense: 10},
		/*5*/ gear{name: "Iron Shield", attack: 0, defense: 10},
		/*6*/ gear{name: "Steel Gloves", attack: 0, defense: 20},
		/*7*/ gear{name: "Steel Helm", attack: 0, defense: 20},
		/*8*/ gear{name: "Steel Curais", attack: 0, defense: 20},
		/*9*/ gear{name: "Steel Boots", attack: 0, defense: 20},
		/*10*/ gear{name: "Steel Shield", attack: 0, defense: 20},
		/*11*/ gear{name: "Iron Sword", attack: 10, defense: 0},
		/*12*/ gear{name: "Iron 2H Axe", attack: 25, defense: 0},
		/*13*/ gear{name: "Iron Warmace", attack: 40, defense: 0},
		/*14*/ gear{name: "Steel Sword", attack: 20, defense: 0},
		/*15*/ gear{name: "Steel Axe", attack: 35, defense: 0},
		/*16*/ gear{name: "Steel Warmace", attack: 60, defense: 0},
	}
}

//Function used to make monsters occur
func makeMonster() {
	allMonsters = []monster{
		monster{name: "weak monster", health: 15, defense: 5, attack: 5, XP: .5},
		monster{name: "monster", health: 30, defense: 10, attack: 7, XP: 1},
		monster{name: "strong monster", health: 35, defense: 15, attack: 14, XP: 1},
		monster{name: "bossmonster", health: 50, defense: 30, attack: 23},
	}
}

//needed a player to test against the encounter function! when i get home on
//wednesday night I'll remove anything debug and submit once I get the ok from you.
func makePlayer() {
	P = Player{name: "Abraxas", level: 1, health: 30, maxHealth: 30, defense: 4, attack: 11}
}

//hey pat! this is cool but how does it apply armor and weapon values to character?
//also is the list of equipped items accessible to be modified?
//change item to functional items from structs
// not certain where I need to use the pointer since it runs into the problem here
//ok i think i see what we can do to fix it, a few lines of code befroehand oughta do it
//did you see how i set it in the encounter? enemy:=allMonsters[variable]? lets try that and see if it works
//brb
func equipping() {
	doEquip := true
	for doEquip == true {
		textBox("What do you want to equip?")
		f.Printf("Your current equipment is %v\n", equipment)
		f.Printf("You have %v in your inventory.\n", bag)
		f.Scanln(&equip)
		if strings.ToUpper(equip) == "BACK" || strings.ToUpper(equip) == "QUIT" {
			TGB(ChartXY, OrigChart, AntiBoydcheeser)
			doEquip = false
			break
		}
		if strings.ToUpper(equip) == "ESC" {
			TGB(ChartXY, OrigChart, AntiBoydcheeser)
			doEquip = false
			break
		}
		bagItem := bag[equip]
		EquipItem := allgear[bagItem]
		if EquipItem.name != equip {
			textBox("Abraxas wants what he doesn't have")
		}
		f.Printf("You equip %s\n", EquipItem.name)
		P.defense = P.defense + EquipItem.defense
		P.attack = P.attack + EquipItem.attack
		// equipment = append(equipment, bag[equip-1])
		// bag = append(bag[:equip-1], bag[equip:]...)
	}
	TGB(ChartXY, OrigChart, AntiBoydcheeser)
}
func discard() {
	// var item string
	var bagItems = make([]string, 10)
	for bagItem := range bag {
		bagItems = append(bagItems, bagItem)
	}
	f.Printf("What item do you want to discard?[Back][Esc][Quit] to return:\n")
	f.Scanln(&Delete)
	if strings.ToUpper(Delete) == "BACK" || strings.ToUpper(Delete) == "QUIT" {
		TGB(ChartXY, OrigChart, AntiBoydcheeser)
	}
	if strings.ToUpper(Delete) == "ESC" {
		TGB(ChartXY, OrigChart, AntiBoydcheeser)
	}
	for _, item := range bagItems {
		if Delete == item {
			delete(bag, Delete)
		} else if Delete != item {
			textBox("You don't have that item in your bag.")
		}
	}
	textBox("Your bag now contains")
	for _, bagitem := range bag {
		f.Printf("%v\n", bagitem)
	}
}

//Functions
func Chests() {
	number := rand.Intn(5)
	switch number {
	case 1:
		chest1()
	case 2:
		chest2()
	case 3:
		chest3()
	case 4:
		chest4()
	case 5:
		chest5()
	}
}
func chest1() {
	f.Printf("You find a chest, do you want to open the chest?\n1=Yes\n2=No\n")
	f.Scanln(&open)
	if open == 1 {
		f.Printf("You open the chest and find %v!!\n", allgear[9].name)
		bag[allgear[9].name] = allgear[9].defense
	} else if open == 2 {
		textBox("You leave the chest unopened and may never know what is in it.")
	} else {
		textBox("Please make a valid choice.")
	}
}
func chest2() {
	f.Printf("You find a chest, do you want to open the chest?\n1=Yes\n2=No\n")
	f.Scanln(&open)
	if open == 1 {
		f.Printf("You open the chest and find %v!!\n", allgear[8].name)
		bag[allgear[8].name] = allgear[8].defense
	} else if open == 2 {
		textBox("You leave the chest unopened and may never know what is in it.")
	} else {
		textBox("Please make a valid choice.")
	}
}
func chest3() {
	f.Printf("You find a chest, do you want to open the chest?\n1=Yes\n2=No\n")
	f.Scanln(&open)
	if open == 1 {
		f.Printf("You open the chest and find %v!!\n", allgear[5].name)
		bag[allgear[5].name] = allgear[5].defense
	} else if open == 2 {
		textBox("You leave the chest unopened and may never know what is in it.")
	} else {
		textBox("Please make a valid choice.")
	}
}
func chest4() {
	f.Printf("You find a chest, do you want to open the chest?\n1=Yes\n2=No\n")
	f.Scanln(&open)
	if open == 1 {
		f.Printf("You open the chest and find %v!!\n", allgear[6].name)
		bag[allgear[6].name] = allgear[6].defense
	} else if open == 2 {
		textBox("You leave the chest unopened and may never know what is in it.")
	} else {
		textBox("Please make a valid choice.")
	}
}
func chest5() {
	f.Printf("You find a chest, do you want to open the chest?\n1=Yes\n2=No\n")
	f.Scanln(&open)
	if open == 1 {
		f.Printf("You open the chest and find %v!!\n", allgear[7].name)
		bag[allgear[7].name] = allgear[7].defense
	} else if open == 2 {
		textBox("You leave the chest unopened and may never know what is in it.")
	} else {
		textBox("Please make a valid choice.")
	}
}

/*ReadFromFile takes any sized grid of text in a text file and puts it into an array to be used as a Chart*/
func ReadFromFile(path string, ChartXY [][]int, AntiBoydcheeser [][]bool) {
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
			ChartXY[i] = make([]int, XaXis)
			AntiBoydcheeser[i] = make([]bool, XaXis)
			for j := range ChartXY[i] {
				ChartXY[i][j] = 0
				AntiBoydcheeser[i][j] = false
			}
		}
		//splitting the lines of ChartY into individual chars for ChartXY as elements
		for i, obj := range ChartY {
			for j := 0; j < len(obj)-1; j++ {
				ChartXY[i][j] = int(obj[j])
				AntiBoydcheeser[i][j] = false
			}
		}
	}
}

/*drawChart draws the map of the region the game is played in, its referencing the
global 2 dimensional slice populated by readfromfile and then uses a nested for loop to print out the map iteratively, interpreting numbers as strings to then print out based on the format of the supplied txt file*/
func drawChart(ChartXY [][]int) {
	//this nested for range loop checks each element pair from the Chart slice and its assigned value then prints out the value's associated icon to the Chart!
	for i, xIs := range ChartXY {
		for j := range xIs {
			if string(ChartXY[i][j]) == "0" {
				f.Print("░") //grass
			} else if string(ChartXY[i][j]) == "1" {
				f.Print("▓") //path
			} else if string(ChartXY[i][j]) == "2" {
				f.Print("Ü") //Player
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
			} else if string(ChartXY[i][j]) == "t" {
				f.Print("♠") //tree!
			} else if string(ChartXY[i][j]) == "B" {
				f.Print("B")
			}
		}
		//pushes onto next row
		f.Print("\n")
	}
	//finishes off map and newlines for non-map stuff
	f.Print("\n")
}

//Collison will detect whether or not player can go a direction
func Collision(DestXY int) (Blocked bool) {
	if string(DestXY) == "0" || string(DestXY) == "1" {
		Blocked = false
	} else if string(DestXY) == "9" || string(DestXY) == "B" {
		Blocked = false
	} else if string(DestXY) == "3" || string(DestXY) == "4" {
		Blocked = false
	} else {
		Blocked = true
	}
	return Blocked
}

/*TGB is a function for Traversing Game Board hence TGB for ease of use
this function will take user input based on a compass rose so North,South,East,West,
NorthEast,NorthWest,SouthEast,SouthWest. Also functions a bit like a controller in a videogame
recieving the input from the player via the driver and executing the logic behind its associated prompts. ie quit, soon to come; inv, look, run, rest */
func TGB(ChartXY [][]int, OrigChart [][]int, AntiBoydcheeser [][]bool) {
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
	//here is where additional choices during exploration can be made...wanna sharpen your blade before a fight? make a switch case to a function that buffs att for a certain # of moves
	f.Print("choose a direction!\n[N][S][W][E]\n[Stats][Inv][Equip][Rest]\n [Quit]=>:")
	f.Scanln(&pDirect)
	pDirect = strings.ToUpper(pDirect)
	//this switch functions like a controller you input the commands and stuff happens
	switch pDirect {
	case "N": //focus on SetY in negative value cuz map is written down to up
		DestXY = ChartXY[SetY-1][SetX]
		ABC = AntiBoydcheeser[SetY-1][SetX]
		collide := Collision(DestXY)
		if collide == false {
			ChartXY[SetY-1][SetX] = ChartXY[SetY][SetX]
			if AntiBoydcheeser[SetY][SetX] == false {
				ChartXY[SetY][SetX] = PrevXY
			} else {
				PrevXY = 1
			}
		} else {
			defer textBox("You can't go that way")
		}
	case "S": //focus on SetY in positive value because ...
		DestXY = ChartXY[SetY+1][SetX]
		ABC = AntiBoydcheeser[SetY+1][SetX]
		collide := Collision(DestXY)
		if collide == false {
			ChartXY[SetY+1][SetX] = ChartXY[SetY][SetX]
			ChartXY[SetY][SetX] = PrevXY
		} else {
			defer textBox("You can't go that way")
		}
	case "W": //focus on SetX in negative value because map is written left to right.
		DestXY = ChartXY[SetY][SetX-1]
		ABC = AntiBoydcheeser[SetY][SetX-1]
		collide := Collision(DestXY)
		if collide == false {
			ChartXY[SetY][SetX-1] = ChartXY[SetY][SetX]
			ChartXY[SetY][SetX] = PrevXY
		} else {
			defer textBox("You can't go that way")
		}
	case "E": //focus on SetX in positive value because...
		DestXY = ChartXY[SetY][SetX+1]
		ABC = AntiBoydcheeser[SetY][SetX+1]
		collide := Collision(DestXY)
		if collide == false {
			ChartXY[SetY][SetX+1] = ChartXY[SetY][SetX]
			ChartXY[SetY][SetX] = PrevXY
		} else {
			defer textBox("You can't go that way")
		}
	case "QUIT":
		leave = true
		return
	case "INV":
		discard()
	case "EQUIP":
		equipping()
	case "REST":
		rest()
	case "STATS":
		f.Print(P)
		f.Print("Press[ENTER] to continue\n")
		f.Scanln()
	}
	SCRWipe()
	drawChart(ChartXY)
	trigChk(DestXY, ABC, SetX, SetY, AntiBoydcheeser)
	return
}

/*trigChk looks at destination of TGB input and checks if its a path or something to
contend with then calls assigned function i.e those icons and the associated #'s
from drawChart()? You can even add more and since they translate to alt codes you
can look up windows alt codes and find ton's of viable icons codes and likely get
them to work*/
func trigChk(DestXY int, ABC bool, SetX int, SetY int, AntiBoydcheeser [][]bool) {
	switch string(DestXY) {
	case "0":
		textBox("A patch of grass lies before you devoid of anything interesting.")
	case "1":
		textBox("A path can be seen through the forest that you can make your way through.")
	case "2":
		panic("It's IMPOSSIBLE TO ACCOMPLISH THIS WITHOUT BREAKING THE GAME!\n") //no touchie!
	case "3":
		if ABC == false {
			textBox("BATTLE!")
			time.Sleep(3 * time.Second)
			Encounter()
		} else if ABC == true {
			textBox("nothing remains of your now dead foe but a smear of blood")
		}
	case "4":
		if ABC == false {
			Chests()
		} else if ABC == true {
			textBox("There is only an empty box here.")
		}
	case "9":
		textBox("Before you is a sturdy door...")
		textBox("Will you proceed?")
	case "B":
		BossFight()
	}
}

//rest is a function to restore player health...seems silly for a player that heals with time to need bandages.
func rest() {
	P.health = P.maxHealth
	textBox("Health Restored!")
}

//Bossfight is the end of the game
func BossFight() {
	Boss := allMonsters[3]
	for P.health > 0 {
		f.Printf("HP:%v\nATTK:%v\nDEF:%v\n", P.health, P.attack, P.defense)
		Boss.health = Boss.health - ((P.attack / Boss.defense) + (P.attack % Boss.defense))
		textBox("Abraxas strikes The SpellSword")
		if Boss.health > 0 {
			//debugging tool for combat smoothness
			f.Printf("Enemy HP:%v\nEnemy ATT:%v\nEnemy DEF:%v\n", Boss.health, Boss.attack, Boss.defense)
			textBox("SpellSword strikes back!")
			P.health = P.health - ((Boss.attack / P.defense) + (Boss.attack % P.defense))
		} else if P.health == 0 {
			textBox("You have fainted from too much damage")
			textBox("Taking advantage of the situation Spellsword...")
			textBox("BURNS YOU TO ASH! YOU ARE DEAD FOR GOOD!")
			os.Exit(666)
		} else {
			textBox("You have won!")
			textBox(story(len(plot) - 1)) //this advances the story to end...
			os.Exit(42)
			break
		}
		time.Sleep(3 * time.Second)
	}
}

//Encounter is the battle function that is more or less automatic, just stat based
func Encounter() {
	//define necessary temp stats like who is fighting and what they might carry
	enemy := allMonsters[0]
	if P.level <= 2 {
		enemy = allMonsters[0]
	} else if P.level > 2 && P.level < 5 {
		enemy = allMonsters[rand.Intn(3)]
	} else if P.level > 5 && P.level < 10 {
		enemy = allMonsters[rand.Intn(2)+1]
	} else {
		enemy = allMonsters[3]
	}
	//fights go automatically, no intervention by player
	for P.health > 0 {
		//debug for combat to checkvals
		f.Printf("HP:%v\nATTK:%v\nDEF:%v\n", P.health, P.attack, P.defense)
		enemy.health = enemy.health - ((P.attack / enemy.defense) + (P.attack % enemy.defense))
		textBox("Abraxas strikes his foe")
		if enemy.health > 0 {
			//debugging tool for combat smoothness
			f.Printf("Enemy HP:%v\nEnemy ATT:%v\nEnemy DEF:%v\n", enemy.health, enemy.attack, enemy.defense)
			textBox("the foe strikes back!")
			P.health = P.health - ((enemy.attack / P.defense) + (enemy.attack % P.defense))
		} else if P.health == 0 {
			textBox("You have fainted from too much damage")
			textBox("your body disintegrates and you re-appear at the start")
			defer Driver(pDirect, FileName)
			textBox("Abraxas has lost a Level!")
			P.level--
		} else {
			textBox("You have won!")
			textBox(story(int(P.level + 3.0))) //this advances the story dood!...
			P.attack = P.attack + 3
			P.defense = P.defense + 2
			P.health = P.health + 5
			P.maxHealth = P.maxHealth + 5
			P.level = P.level + enemy.XP
			//looting time
			break
		}
		time.Sleep(3 * time.Second)
	}
	TGB(ChartXY, OrigChart, AntiBoydcheeser)
}
func textBox(message string) {
	//mRows := (len(message) % 76)
	spacer := ""
	middlesection := ""
	middlelength := len([]rune(message)) + 1

	for i := 0; i <= middlelength; i++ {
		middlesection += "═"
	}
	topsect := "╔" + middlesection + "╗"
	for i := 0; i <= middlelength; i++ {
		spacer += " "
	}
	sides := "║" + spacer + "║"
	bottomsect := "╚" + middlesection + "╝"

	//message box!
	f.Printf("%s\n", topsect)
	f.Printf("%s\n", sides)
	f.Printf("║ %s ║\n", message)
	f.Printf("%s\n", sides)
	f.Printf("%s\n", bottomsect)
}

//MakeOriginalChart is a function which is used to create a comparative slice for moving across Game map without leaving blank spaces
func MakeOriginalChart(ChartXY [][]int, OrigChart [][]int) {
	for i := range ChartXY {
		OrigChart[i] = make([]int, len(ChartXY[i]))
		copy(OrigChart[i], ChartXY[i])
	}
}

//init initializes string func map of OSes for SCRWipe()
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
		panic("unrecognized platform, run program on Windows OS, MacOS, or Linux OS")
	}
}

//using a file scanner to read lines from txt file to save space in source code.
//also cuz this makes it possible for others to write their own lore etc while still-
//maintaining the mechanics etc.
func story(Memories int) string {
	GamePlot := "story.txt"
	if len(plot) <= 0 {
		story, err := os.Open(GamePlot)
		if err != nil {
			log.Fatalf("either incorrect format or associated with %s please rewrite.\n", err)
		} else {
			StoryReader := bio.NewScanner(story)
			StoryReader.Split(bio.ScanLines)
			defer story.Close()
			for StoryReader.Scan() {
				plot = append(plot, StoryReader.Text())
			}
		}
	}
	return plot[Memories]
}

//Driver runs the show, takes a string for direction to travel including QUIT
//also uses the FileName string to import txt file just type in name no .txt necessary
func Driver(pDirect string, FileName string) {
	ReadFromFile("map.txt", ChartXY, AntiBoydcheeser)
	MakeOriginalChart(ChartXY, OrigChart)
	drawChart(ChartXY)
	textBox(story(0))
	time.Sleep(3 * time.Second)
	textBox(story(1))
	time.Sleep(3 * time.Second)
	textBox(story(2))
	for leave == false {
		TGB(ChartXY, OrigChart, AntiBoydcheeser)
	}
}
func main() {
	//Seed randomNumberGenerator first things first!
	rand.Seed(time.Now().UnixNano())
	// creates armor, weapon, and stats for use
	makeEquipment()
	// creates all monsters
	makeMonster()
	//makes player stats
	makePlayer()
	Driver(pDirect, FileName)
}
