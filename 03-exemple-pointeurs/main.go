package main

import "fmt"

func UpdateVal(val string) { // Ne fonctionne pas car passe par copie
	val = "Louis"
}

func UpdateValPtr(val *string) { // Ne fonctionne pas car passe par copie
	*val = "Simon"
}

func main() {
	i := 1
	var p *int = &i

	fmt.Printf("i=%v\n", i)
	fmt.Printf("p=%v\n", p)
	fmt.Printf("p=%v\n", *p)
	fmt.Println("----------------")

	s := "Bob"
	var sPtr *string = &s
	s2 := *sPtr

	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("----------------")

	*sPtr = "Alice"
	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("----------------")

	UpdateVal(s)
	fmt.Printf("s=%v\n", s)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Println("----------------")

	UpdateValPtr(&s)
	fmt.Printf("s=%v\n", s)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Println("----------------")
}
