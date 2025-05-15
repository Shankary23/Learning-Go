package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rps() {
	rand.Seed(time.Now().UnixNano())
	choices := []string{"R", "P", "S"}
	computer := choices[rand.Intn(len(choices))]
	var choice string
	fmt.Print("Choose either RPS: ")
	fmt.Scanln(&choice)
	fmt.Print(computer, "\n")
	if choice != "R" && choice != "S" && choice != "P" {
		fmt.Print("Invalid Input, Try Again! \n")
		rps()
	} else if choice == computer {
		fmt.Print("Try Again \n")
		rps()
	} else if choice == "R" || choice == "r" && computer == "R" {
		fmt.Print("Try Again \n")
		rps()
	} else if choice == "R" && computer == "S" || choice == "P" && computer == "R" || choice == "S" && computer == "P" {
		fmt.Print("You Win! ")
		return
	} else {
		fmt.Print("You loose!")
		return
	}
}
