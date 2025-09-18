package main

import "fmt"

func print() {
	fmt.Println("new learning")
	fmt.Println(len("Learn"))
}

func main() {
	var name string
	name = "minez"
	fmt.Println(name)


	var kendaraan = "motor"
	fmt.Println(kendaraan)

	
	// famouse declare 
	a := "first declare"

	a = "second declare"
	fmt.Println(a)

	// multiple declare
	var (
		firstName = "minez"
		secondName = "zota"
		thirdName = "enedi"
	)

		fmt.Println(firstName, secondName, thirdName)


	// const if not use no error but cannot re-declare
	const newName string = "hadi"
	const lastName = "asta"

}

