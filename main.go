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


	// Data type conversion 

	var nilai32 int32 = 3232
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai64)

	var newDataName = "haru"
	var e = newDataName[0] // 104 outputnya 
	var eString = string(e)
		fmt.Println(eString, " <<< ") // output : h

	type NoKTP string

	var ktpHadi NoKTP = "1111111111"

	fmt.Println(ktpHadi)

	var i = 10 
	i += 10
	i++
	fmt.Println(i)


	//====> condition
	
	var name1 = "woila"
	var name2 = "woila"

	var result = name1 == name2

	fmt.Println(result)

		//====> array dibatasin 3 
	var array1  [3]string 
	array1[0] = "nada"
	array1[1] = "nadi"
	array1[2] = "nudu"
		

	fmt.Println(array1)



	var arrayDirect = [3]int{
		22,33,33,
	}
	
	fmt.Println(arrayDirect)

	//====> array method / function 

	// array harus di declare langsung kalau tidak mau ada batasan index array 
	// mau berapapun di isi ini bisa tapi kalau tidak di isi maka array isinya kosong
	var values = [...]int{ 
		22,33,44,22,11,
	}
	fmt.Println(values)


	} 

	

