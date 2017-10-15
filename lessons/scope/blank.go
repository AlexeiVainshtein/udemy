package scope

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func Blank() {

	res, err := http.Get("http://www.google.com")
	if err != nil {
		panic("in the panic")
	}

	res.Body.Close()

	res1, _ := http.Get("http://www.google.com")
	page, _ := ioutil.ReadAll(res1.Body)
	fmt.Println("Printing")
	fmt.Println(len(page))
	fmt.Printf("%s", page)
	fmt.Println("done")

	res1.Body.Close()

	page2, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		fmt.Println("In the err2 " + err2.Error())
	}
	fmt.Println("After the err check")
	fmt.Println(len(page2))

}
