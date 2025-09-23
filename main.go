package main

import (
	"fmt"
)

func print(firstName string, lastName string) {
	fmt.Println("new learning")
	fmt.Println(len("Learn"))
	fmt.Println(firstName, lastName)
}

func basicKnowledge() {
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
		firstName  = "minez"
		secondName = "zota"
		thirdName  = "enedi"
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
	var array1 [3]string
	array1[0] = "nada"
	array1[1] = "nadi"
	array1[2] = "nudu"

	fmt.Println(array1)

	var arrayDirect = [3]int{
		22, 33, 33,
	}

	fmt.Println(arrayDirect)

	//====> array method / function

	// array harus di declare langsung kalau tidak mau ada batasan index array
	// mau berapapun di isi ini bisa tapi kalau tidak di isi maka array isinya kosong
	var values = [...]int{
		22, 33, 44, 22, 11,
	}
	fmt.Println(values)
	// tidak ada function untuk menghapus index dalam array seperti javascript

	// data type slice dan sering sekali dipakai karena ukuran slice dapat berubah
	// cara kerja :
	// ketika slice memanage array , dan ketika penuh maka akan dipindahkan ke array baru
	var names = [...]string{
		"hadi",
		"asta",
		"budi",
		"andi",
		"tami",
	}
	slice := names[2:4] // dimulai dari array index ke 2 dan berakhir di index 4 kalau mau ambil semua kosongkan saja cth : names[:]

	fmt.Println(slice)

	// cek lenth array
	lenghtSlice := len(slice)

	fmt.Println(lenghtSlice)
	// check capacity array
	capSlice := cap(slice)
	fmt.Println(capSlice)

	// method atau function slice banyak mangkanya enak untuk di pakai  ada
	// append(slice,data)  buat slice baru dengan menambahkan di posisi terakhir
	//  make([]typedata, length, capacity) buat slice baru
	//  copy(destination , source) saling dari source

	// dataType Map
	dataMap := map[string]string{
		"name":    "hadi",
		"kelas":   "12",
		"address": "jalan",
	}
	// akses value dengan key
	fmt.Println(dataMap["name"])
	// akses semua yang ada pada map
	fmt.Println(dataMap)

	// function / method map ada banyak
	// make new map with method "make"
	newDataMap := make(map[string]string)

	newDataMap["title"] = "new journey"
	newDataMap["book"] = "new book"
	newDataMap["schedule"] = "new schedule"
	// delete by title
	delete(newDataMap, "book")
	fmt.Println(newDataMap)

	// if condition
	myNames := "nageeeeeeeeee"

	if myNames == "hadi" {
		fmt.Println("nama kamu hadi")
	}
	// if else
	if myNames == "naga" {
		fmt.Println("asta")
	} else {
		fmt.Println("selain asta")
	}
	// else if
	if myNames == "nage" {
		fmt.Println("1")
	} else if myNames == "nagu" {
		fmt.Println("2")
	} else if myNames == "nagi" {
		fmt.Println("3")
	} else {
		fmt.Println("selain ini 4")
	}

	// if short statement
	if length := len(myNames); length > 5 {
		fmt.Println("name length", length)
	} else {
		fmt.Println("name here dibawah 5 characther")
	}

	// switch
	nameSwitch := "hadi"

	switch nameSwitch {
	case "nadi":
		fmt.Println("nadi")
	case "nodi":
		fmt.Println("nodi")
	default:

		fmt.Println("not nodi & nadi")

	}

	// short statement switch
	switch length := len(nameSwitch); length > 5 {
	case true:
		fmt.Println("lebih dari 5")
	case false:
		fmt.Println("kurang dari 5")
	}

	// for loop
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println("ulang sebanyak count")
	}

	for count < 5 {
		fmt.Println("kurang dari 5")
	}

	// for range
	//  arti dari code di bawah ini adalah declare data for range dengan data array yang belum ditentukan isinya berapa tergantung yang di sign di dalam array
	dataForRange := []string{"hadi", "hasta", "nama"}
	// jika i lebih kecil dari panjang index yang ada dalam dataforRange
	for i := 0; i < len(dataForRange); i++ {
		// cetak data for range index ke i
		fmt.Println(dataForRange[i])

		for index, dataForRange := range dataForRange {
			fmt.Println("index", index, "=", dataForRange, " >>>>> ")
		}
		// kalau gak butuh index bisa di jadiin _ aja
		for index, dataForRange := range dataForRange {
			fmt.Println(dataForRange, index, " <<< ")
		}

		// break

		for i := 0; i < 5; i++ {
			if i == 2 {
				break
			}
			fmt.Println("counting ", i)
		}
	}

	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// fmt.Println("ini i ke", i)
			continue
		}
		// fmt.Println("ini i ke woow", i)
	}

}

// test return function
func returnString(kata string) string {
	hello := "halo " + kata
	return hello
}

func multipleString() (string, string) {
	return "hadi", "hasta"
}

func getCompleteDeclaration() (string, string, string) {
    return "hadi", "hadi", "hadi"
}


func main() {
	// basicKnowledge()
	// print("hadi", "hasta")

	// ambil data dari function returnString
	result := returnString("hadi")
	fmt.Println(result)

	// kalau tidak mau di pakai ganti parameter dengan _
	// first, _ := multipleString()
	first, second := multipleString()
	fmt.Println(first, second)

 a, b, c := getCompleteDeclaration()
 fmt.Println( a, b, c )
}
