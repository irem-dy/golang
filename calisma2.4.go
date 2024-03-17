package main

import "fmt"

func main() {
	var sayi int = 15
	fmt.Println(sayi)
	var metin string
	metin = "merhaba dunya 123.6"
	fmt.Println(metin)
	// şimdi şöyleki burada " "
	//arasına aldığımız sayilar artik birer karakter olmaktadır
	// yani herhangi bir şekilde integer tipine bağlı sayıyla
	//aritmetik bir ifade yapamayız.
	// fmt.println(5+metin) şeklinde yazdığımız bir şey
	//bize çıktı vermeyecektir.
	// aynı şekilde ("5"+metin) şeklinde yazdığımız ifadede bize
	// bunları string yapısında aldığı için yan yana yazdırarak verecektir.

}
