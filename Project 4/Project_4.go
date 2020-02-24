package main

import (
	f "fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"text/tabwriter"
	"time"
	t "time"
)

//Deathod is the methods of death struct
type Deathod struct {
	name           string  //name of the deathod.
	cost           float64 //how much the deathod costs.
	Desc           string  //describes the method of death for user to decide not actual description of death.
	DeathScription string  //flavor text during the killin part.
	count          int
}

//d is the slice of deaths to put in the cart to make combo deaths!
var d []Deathod

//CartSlice is the individual choice(s) user makes that get added to
var CartSlice []int

//isDead is the variable that keeps the driver running till customer has been serviced!
var isDead bool = false

//Wallet is just so, your wallet bruh!
var Wallet float64

//Bill is the cost for services to be rendered
var Bill float64

//randomDeaths randomizes the amount of deaths for each choice to simluate previous useage
func randomDeaths() int {
	return rand.Intn(10000000)
}

//MakeDeathods is the make function for the killomatic booth.
func makeDeathod() {
	randDeaths := []int{0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 7; i++ {
		randDeaths[i] = randomDeaths()

	}
	d = []Deathod{
		/*1*/ Deathod{"Slow and Horrible", 0, "The Collect-Call of Deaths. Good god just die already!", "a Phone pops out from a panel, a robotic arm grasps your head firmly...\n audio of hold music pays on loop\n just loud enough to slowly liquefy your insides...your brains melt out your ear\n", randDeaths[0]},
		/*2*/ Deathod{"Prison Special", 5, "When Three hots and a cot isn't enough...Shank!", "an arm pops out holding a crudely sharpened toothbrush and proceeds to stab you in the kidneys, Oh tHe hOrRoR!", randDeaths[1]},
		/*3*/ Deathod{"Clumsy Bludgeoning", 10, "Cuz who doesn't wanna get beaten to death by a blind guy?!", "a robotic arm pops out of a panel holding a cane for the blind and\nstarts swinging wildly thwacking you about the head and neck\n", randDeaths[2]},
		/*4*/ Deathod{"Hot'n'Offal", 12, "Death by Molten Rocks AKA Bludgeoning Plus", "an ominous rumble is heard overhead...\nsuddenly ports fly open all around spewing molten hot guts and rocks\n in a sizzling millieu of guts and death...glorious\n", randDeaths[3]},
		/*5*/ Deathod{"Fast'n'Ugly", 15, "So fast its Fugly, split by a steel wire at mach 7", "You hear the stretching sound of a large spring under load and a slot\n opens up above you then...BOOM! you explode from the speed of the wire \n that just passed through your middle top to bottom\n", randDeaths[4]},
		/*6*/ Deathod{"Quick 'n' Painless", 25, "Death by Gas...not the kind from an Ass!", "Vents appear in front of you, a soft whoosh of air hits your face and then darkness...\npoisoned by concentrated Carbon monoxide before it even occured to you", randDeaths[5]},
		/*7*/ Deathod{"Valhalla Sendoff", 100, "When you want to go with honor...it'll cost ya.", "a sword drops into your hand and a battle horn sounds\n an arm shoots out in front of you wielding a sword and slices your head off! A good death.", randDeaths[6]},
	}
}

//Killomatic is the driver program
func Killomatic() {
	subtotal := 0.0
	//tw is tabwriter to make lists look neat
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 4, ' ', tabwriter.AlignRight|tabwriter.Debug)
	//First the browse bool is started once user agrees on no returns.
	Browse := false
	//set to false so it won't start until Browse section completes and user elects to buy
	Purchase := false
	sChoice := ""
	//the Cart
	CartSlice := make([]int, 1, 7)
	intChoice := 0
	//set to false until ordering is complete
	Payment := false
	//Killomatic Starts here
	f.Print("Welcome to the Killomatic!\n")
	t.Sleep(2 * t.Second)
	f.Print("This is your only chance to go back, all sales are final! Continue?or backwards???(seY/oN)\n")
	f.Scanln(&sChoice)
	if strings.ToUpper(sChoice) == "NO" {
		f.Print("Then why are you even here?\n")
		for i := 3; i > 0; i-- {
			f.Printf("Ejection in %v second(s)\n", i)
			t.Sleep(1 * t.Second)
		}
		os.Exit(1)
	} else {
		Browse = true
	}
	//Begin Browsing...
	for Browse == true {
		f.Print("\t\tPlease select your Death(s) for Details\n")
		for i, methods := range d {
			f.Printf("%d)%s:\t\t\t$%.2f\n", i+1, methods.name, methods.cost)
		}
		f.Scanln(&intChoice)
		switch intChoice {
		case 1:
			f.Printf("%s,\n%s\n%v People have used this option\n", d[0].name, d[0].Desc, d[0].count)
		case 2:
			f.Printf("%s,\n%s\n%v People have used this option\n", d[1].name, d[1].Desc, d[0].count)
		case 3:
			f.Printf("%s,\n%s\n%v People have used this option\n", d[2].name, d[2].Desc, d[2].count)
		case 4:
			f.Printf("%s,\n%s\n%v People have used this option\n", d[3].name, d[3].Desc, d[3].count)
		case 5:
			f.Printf("%s,\n%s\n%v People have used this option\n", d[4].name, d[4].Desc, d[4].count)
		case 6:
			f.Printf("%s,\n%s\n%v People have used this option\n", d[5].name, d[5].Desc, d[5].count)
		case 7:
			f.Printf("%s,\n%s\n%v People have used this option\n", d[6].name, d[6].Desc, d[6].count)
		}
		f.Print("Would you like to See Another?(Y/N)")
		f.Scanln(&sChoice)
		if strings.ToUpper(sChoice) == "N" {
			Browse = false
			Purchase = true
		} else {
			Browse = true
		}
	}
	//Begin filling cart...
	cartSlot := 0
	for Purchase == true {

		for i, methods := range d {
			f.Fprintf(tw, "%v)\t%s\t $%.2f USD\n", i+1, methods.name, methods.cost)
		}

		f.Printf("Wallet: $%.2f", float64(Wallet)-subtotal)
		f.Print("Please make your selection, there is no escape.\n")
		if cartSlot < 1 {
			f.Scanln(&intChoice)

			CartSlice[0] = intChoice - 1
			Bill += d[CartSlice[0]].cost
			subtotal += d[CartSlice[0]].cost
			if Bill > Wallet {
				log.Panicf("!##ERROR##! You do not have %.2f amount of money available in you wallet!", Bill)
				Bill -= d[CartSlice[0]].cost
				subtotal -= d[CartSlice[0]].cost
			} else {
				f.Printf("\"%s\" added to cart\n", d[CartSlice[0]].name)
			}
			f.Print("Would you like upgrade to a combo death?(Y/N)?\n")
			f.Scanln(&sChoice)
			if strings.ToUpper(sChoice) == "Y" {
				f.Printf("Most Excellent! what else will it be?\n")
				Purchase = true
			} else if strings.ToUpper(sChoice) == "N" {
				f.Printf("Very good then, %s it is.\n", d[CartSlice[0]].name)
				Purchase = false
			}
		} else if cartSlot >= 1 {
			f.Scanln(&intChoice)
			CartSlice = append(CartSlice, intChoice-1)
			Bill += d[intChoice-1].cost
			subtotal += d[intChoice-1].cost
			if Bill > Wallet {
				log.Panicf("!##ERROR##! You do not have %.2f amount of money available in you wallet!", Bill)
				Bill -= d[intChoice-1].cost
				subtotal -= d[intChoice-1].cost
			} else {
				CartSlice = append(CartSlice, intChoice-1)
			}
			f.Printf("Will there be anything else?(Y/N)\n")
			f.Scanln(&sChoice)
			if strings.ToUpper(sChoice) == "N" {
				Purchase = false
				Payment = true
			}
		}
		if Purchase == true {
			cartSlot++
		}
	}
	for Payment == true {
		Amount := 0.0
		f.Printf("The bill for your selection is $%.2f, please pay necessary amount or\nirreversible consequences will occur\n", Bill)
		f.Printf("Enter amount here: ")
		f.Scanln(&Amount)
		if Amount != Bill {
			f.Printf("incorrect amount, you will now be branded and permanently assigned to work as a Wal-mart multi-Dimensional greeter\nin addition you will be hypnotized to love living as a Wal-Mart employee\n...have a nice life!")
			//shouldn't have done that
			for doom := 0; doom < 1; doom-- {
				log.Panicf("Welcome to Wal-mart, I love you\n")
			}
		} else if Amount == Bill {
			f.Printf("Payment received!\n")
			Payment = false
			Wallet = Wallet - Amount
		}
	}
	//Begin the Killing...(0)_(0)
	f.Print("Thank you for Choosing Killomatic!\nEnjoy Dying!\n")
	for isDead == false {
		CommenceKill(CartSlice)
	}
}

//CommenceKill is the function which enacts the users choice of Deathod.
func CommenceKill(CartSlice []int) bool {

	for _, death := range CartSlice {
		f.Printf("Now Preparing: %s. Please Hold Still.\n", d[death].name)
		f.Printf("%s", d[death].DeathScription)
		d[death].count++

	}
	f.Print("And with that you are dead...buh bye now!")
	f.Print("everything fades to black...you wake up hands bound on a wagon to Helgen\n")
	isDead = true
	return isDead
}

func main() {
	rand.Seed(time.Now().UnixNano())
	makeDeathod()
	Wallet = float64(rand.Intn(300) + 10)
	Killomatic()
}
