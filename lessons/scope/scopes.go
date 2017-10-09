package scope

import "fmt"

func Test() (string){
	return my()
}

func MyFanc() (int) {
	fmt.Println(myBol)
	return test()
}

func test()(int)  {
	return x
}