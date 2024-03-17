package main

import "fmt"

func main() {
	a:=10
	b:=15
	c:=a+b
	//fmt.Println("A sayisi ile B sayisinin toplami 
	//c sayisidir")
 // şeklinde yazdırmak için
 fmt.Println(a,"sayisi ile ",b, "sayisinin toplami" ,c, "sayisidir")
// şeklinde yazılabilir fakat çok uzun
// bunun yerine printf(format)fonksiyonunu kullanacağız

fmt.Printf(" %d sayisi ile %d sayisinin toplami %d dir",a,b,c)
// integer bir sayiyi bir string içerisinde yazdırmak istiyorsak %d kullanırız.


}
