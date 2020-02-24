package main

import (
  "fmt"
  // "strings"
  // "time"
  // "math/rand"
)

type allArmor struct {
  name string
  stat  int
}

type allWeapons struct {
  name string
  stat int
}

type baseStats struct{
  name string
  stat int
}

func buildequip(){
baseStats:=[]base{
  base{name:"health", stat:10},
  base{name:"attack", stat:10},
  base{name:"defense", stat:10},
  }
allArmor:=[]armor{
  armor{name:"gloves", stat:10},
  armor{name:"helm", stat:10},
  armor{name:"chest", stat:10},
  armor{name:"boots", stat:10},
  }
allWeapons:=[]weaps{
  weaps{name:"sword", stat:10},
  weaps{name:"axe", stat:10},
  weaps{name:"mace", stat:10},
  }
}




var pickup, d int
var item string
var righthand string
var offhand string
var helm string
var chest string
var boots string
var gloves string
var delete int
var equip int
var doing int
var bag []string
var equipment []string

func attaining(){
  for i:=0;i<10;i++{
    fmt.Scanln(&pickup)
    if pickup==1 {
      fmt.Printf("You put %s into your bag of holding.\n", item)
      bag=append(bag, item)
      break
  } else if pickup==2 {
      fmt.Printf("You leave the item where you found it.\n")
      break
  } else {
      fmt.Printf("Take it or leave it, make a real choice!!\n")
  }
}
}

func discard() {
  fmt.Printf("What item do you want to discard? 1-10\n")
  for i, bag:=range bag{
    fmt.Printf("%d) %s\n",i+1,bag)
  }
  fmt.Scanln(&delete)
  if delete==1{
   	  bag=append(bag[:delete-1], bag[delete:]...)
  } else if delete==2{
      bag=append(bag[:delete-1], bag[delete:]...)
  } else if delete==3{
      bag=append(bag[:delete-1], bag[delete:]...)
  } else if delete==4{
      bag=append(bag[:delete-1], bag[delete:]...)
  } else if delete==5{
      bag=append(bag[:delete-1], bag[delete:]...)
  } else if delete==6{
      bag=append(bag[:delete-1], bag[delete:]...)
  } else if delete==7{
      bag=append(bag[:delete-1], bag[delete:]...)
  } else if delete==8{
      bag=append(bag[:delete-1], bag[delete:]...)
  } else if delete==9{
      bag=append(bag[:delete-1], bag[delete:]...)
  } else if delete==10{
      bag=append(bag[:delete-1], bag[delete:]...)
  } else{
      fmt.Printf("Your bag is not that big.\n")
  }
  fmt.Printf("Your bag now contains\n")
   	for i, bag:=range bag {
   		fmt.Printf("%d) %v\n", i+1, bag)
   	}
}

func equipped(){
  fmt.Printf("What do you want to equip?\n")
  fmt.Printf("You have %v in your inventory.\n", bag)
  fmt.Scanln(&equip)
    if equip==1{
        equipment=append(equipment, bag[equip-1])
     	  bag=append(bag[:equip-1], bag[equip:]...)
    } else if equip==2{
        equipment=append(equipment, bag[equip-1])
        bag=append(bag[:delete-1], bag[delete:]...)
    } else if equip==3{
      equipment=append(equipment, bag[equip-1])
      bag=append(bag[:delete-1], bag[delete:]...)
    } else if equip==4{
      equipment=append(equipment, bag[equip-1])
      bag=append(bag[:delete-1], bag[delete:]...)
    } else if equip==5{
      equipment=append(equipment, bag[equip-1])
      bag=append(bag[:delete-1], bag[delete:]...)
    } else if equip==6{
      equipment=append(equipment, bag[equip-1])
      bag=append(bag[:delete-1], bag[delete:]...)
    }
}

// func weaponrack(){
//   choice:=[3]string{}
// }

func main() {
  fmt.Printf("Your bag contains %v.\n", bag)
  fmt.Printf("You wake up in a forest and your head hurts.\n")
  fmt.Printf("You wander around and find a glove on the ground.\n")

  item="glove"
  helm="helm of defense"
  fmt.Printf("Do you want to pick up %s\n", item)

  attaining()
  fmt.Printf("You look a little further and find an axe in a tree stump.\n")
  item="axe"
  fmt.Printf("Do you take the axe?\n")
  attaining()
  fmt.Printf("Your bag of holding contains %v\n", bag)

  equipped()
  fmt.Printf("You have %v euipped.\n", equipment)
  // if doing==1{
  //   //north
  // } else if doing==2{
  //   //east
  // } else if doing==3{
  //   //south
  // } else if doing==4{
  //   //west
  // } else if doing==5{
  //     discard()
  // } else if doing==6{
  //     equipped()
  // }
  fmt.Printf("Your bag of holding contains %v\n", bag)
}
