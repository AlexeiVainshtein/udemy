package constant

import "fmt"

const p = "tuli"

const q int = 42

const my string = "dadi"

const (
	Pi   = 3.14
	lang = "Go"
)

func PrintConstant() {
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(my)
	fmt.Println(Pi)
	fmt.Println(lang)

	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	const (
		d = iota
		e = iota
		f = iota
	)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
