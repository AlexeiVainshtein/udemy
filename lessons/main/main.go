package main

import (
	"fmt"
	"log"
	"github.com/AlexeiVainshtein/udemy/lessons/mypackage"
	"github.com/AlexeiVainshtein/udemy/lessons/scope"
	"github.com/AlexeiVainshtein/udemy/lessons/constant"
)

func main() {
	println("test")
	fmt.Println("ttt")
	log.Print("From log")
	fmt.Println("From numeral:")
	fmt.Println(42)
	fmt.Println("From binary:")
	fmt.Printf("decimal %d - binary %b - hexedeckcimal %x \n", 42, 42, 42)
	fmt.Printf("decimal %d - binary %b - hexedeckcimal %#x \n", 42, 42, 42)
	fmt.Printf("decimal %d - binary %b - hexedeckcimal %#X \n", 42, 42, 42)

	fmt.Println("Now it's time for the loop")

	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}

	fmt.Println(mypackage.Test())
	fmt.Println(mypackage.MyFanc())
	fmt.Println(mypackage.Bool())
	mypackage.My()
	mypackage.Test()
	{
		fmt.Println("Inner scope")
	}

	funcExpression := func () int {
		const name= iota
		return name
	}

	fmt.Println(funcExpression())
	st := mypackage.TestVar
	fmt.Println(st)
	fmt.Println("################")
	fmt.Println("Scopes")
	fmt.Println(scope.MyFanc())
	fmt.Println(scope.Bool())
	scope.Blank()
	fmt.Println("################")
	fmt.Println("################")
	fmt.Println("Constants:")
	constant.PrintConstant()
	fmt.Println("################")

}
