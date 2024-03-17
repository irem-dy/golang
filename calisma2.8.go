package main

import "fmt"

func main() {
	// bu örnekte float tipine ait bir sayiyi formatlayacağız.
	num := 7.2
	fmt.Printf("sectigimiz sayi: %f", num)
	// burada sayiyi yazdıktan sonra 7.20000 gibi bir değer yazdırdı
	// bunuda düzeltmek adına üstte % yazdıgımız kısmın basına 1 ya da 2 gibi değer vereceğiz
	// %.2f gibi burada yazdığımız 2 ya da 1 virgülden sonra ne kadar değeri
	// görmek istediğimizle ilgildir.

}
