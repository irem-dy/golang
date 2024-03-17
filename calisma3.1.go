// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	// reel sayinin tam ve ondalıklı kısmını hesaplayalım
	// kullanıcı 7.6 sayisi girdi diyelim örneğin tam kısmı 7
	//ondalıklı kısmı 0.6 olarak yazalım
	var sayi float64
	fmt.Println("bir sayi giriniz:")
	fmt.Scan(&sayi)
	tamkisim := int(sayi)
	fmt.Println(tamkisim)
    reelkisim:=sayi-float64(tamkisim)
    // burada sayidan tam kısmı çıkarırken 
	//tam kısmı sayi float türünde olduğu için 
	//tam kısmıda float64 değişkenine çevirerek çıkarma 
	//işlemi yapacaktır.
	fmt.Printf("tam kisim: %d ",tamkisim)
    fmt.Printf("Reel kısım: %.2f",reelkisim)
}
