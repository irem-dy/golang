package main

import "fmt"

func main() {

//fmt.Println("Merhaba "Türkiye" ")
// burada hata verecektir. neyi kullanmamız gerektiğini neyi 
//yazdırması gerektiğini algılayamıyor
// bunun içinde escape characters (kaçıs karakterleri)
// kullanılacaktır
//bunu string ifadeden ayrıymış gibi niteleyen bir şey değilde 
//stringin içindeki karakter olduğunu söylemememiz gerekmektedir.

//kullnaırken
fmt.Println("merhaba \"türkiye\" ")

fmt.Print("merhaba dunya\n")//newline ln \n burada alt satıra geçmesini sağlar.
fmt.Print("hello world")



}
