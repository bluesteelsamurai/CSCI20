package main

import (
  "fmt"
  "time"
  "math/rand"
  "os"
)
/*
global variables and what they mean
Mhp--monster hit points
Php--player hit points
damage--damage dealt or healed
room--current location
map--overview of rooms in temple
*/
func MainMenu(){

   _____                         _
  /__   \ ___  _ __ ___   _ __  | |  ___
    / /\// _ \| '_ ` _ \ | '_ \ | | / _ \
   / /  |  __/| | | | | || |_) || ||  __/
   \/    \___||_| |_| |_|| .__/ |_| \___|
                         |_|
             __
      ___   / _|
     / _ \ | |_
    | (_) ||  _|
     \___/ |_|

      ___              _    _
     /   \ ___   __ _ | |_ | |__
    / /\ // _ \ / _` || __|| '_ \
   / /_//|  __/| (_| || |_ | | | |
  /___,'  \___| \__,_| \__||_| |_|







}
//monster function
func Monster(){

}
//player function for
func Player(){

}

func Inventory(){

}

func Item(){

}

func main(){
//start the random number generator for hit or miss values or probability of loot etc
rand.seed(time.Now().UnixNano())
}
