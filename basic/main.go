package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rpc()
}

func rpc(){

	intro_1 := "---------------Welcome to Rock Paper Scissor----------------"
	intro_2:=  "Press 1 to start the game or any other key to end the game"
	fmt.Println(intro_1)
	fmt.Println(intro_2)
	rock:= "    _______\n---'   ____)\n      (_____)\n      (_____)\n      (____)\n---.__(___)"
	paper:="    _______\n---'   ____)____\n          ______)\n          _______)\n          _______)\n---.__________)"
	scissor:="    _______\n---'   ____)____\n          ______)\n        __________)\n      (____)\n---.__(___)"
	var input int
	fmt.Scan(&input)
	
	for input==1{
		fmt.Println("Press 1 for rock")
		fmt.Println("Press 2 for paper")
		fmt.Println("Press 3 for scissor")
		var comp = rand.Intn(3)+1
		var value  int
		
		fmt.Scan(&value)
		fmt.Println(value)
		switch value {
		case 1:
			fmt.Println("You chose Rock")
			fmt.Println(rock)
		case 2:
			fmt.Printf("You chose Paper\n")
			fmt.Println(paper)
		case 3:
			fmt.Printf("You chose Scissor\n")
			fmt.Println(scissor)
		}
		time.Sleep(2000 * time.Millisecond)
		switch comp {
		case 1 :
			fmt.Printf("Computer chose Rock\n")
			fmt.Println(rock)
		case 2 :
			fmt.Printf("Computer chose Paper\n")
			fmt.Println(paper)
		case 3:
			fmt.Printf("Compute chose Scissor\n")
			fmt.Println(scissor)
			
		}
		time.Sleep(2000 * time.Millisecond)
		if (comp==value){
			fmt.Println("You both tied")
		}else if value==1{
			if(comp==3){
				fmt.Println("---------You won----------")
			}else{
				fmt.Println("You lose")
			}
		}else if value ==2 {
			if(comp==1){
				fmt.Println("---------You won----------")
			}else{
				fmt.Println("You lose")
			}
		}else{
			if(comp==2){
				fmt.Println("---------You won----------")
			}else{
				fmt.Println("You lose")
			}
		}
		fmt.Println()
		fmt.Println(intro_2)
		fmt.Scan(&input)
	}
}