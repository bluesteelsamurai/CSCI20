package main

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
	// variables to be kept within the driver function
	//FileName string variable for opening txt file, later appended with .txt so user simply inputs name
	var FileName string
	pDirect := "" //used to grab input from user for use in driver and TGB
	//Variable for testing location values
	//var location, location2 int = 0, 0
	//this will start player movement.
	Driver(pDirect, FileName)
	//stuff for debugging etc
	/*  f.Printf("the length of Chart: %v \n", len(ChartXY)) //for pathfinding purposes...to be deleted before release
	f.Print("The player start location is denoted by this \u25B2 symbol @ 43,1\n")
	f.Printf("which one you wanna look at?\n X then Y:\n")
	f.Scanln(&location)
	f.Print("Second value:")
	f.Scanln(&location2)
	//location and location2 are reversed because the first index box is associated with the y-axis and second is x-axis
	f.Printf("the value of location %v,%v: %v.\n", location, location2, ChartXY[(YaXis - location2)][(XaXis-location)])
	f.Printf("values and what they mean:\n 0=>░|49=>▓|50=>▲|51=>?|52=>T|53=>║|54=>═ |55=>\\|56=>/|57=>D\n")
	*/

}
