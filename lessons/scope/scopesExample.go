package scope

import "fmt"

var TestVar = "test var"
var myBol bool
var x = 44
 func my()(string) {
	a := "Testt"
	return TestVar + a
}

func Bool()(bool) {
	return myBol
}

func My() {
	fmt.Println("From My another printing")
	fmt.Printf("%T what?? \n", myBol)
	fmt.Printf("%v what?? \n", myBol)
}
