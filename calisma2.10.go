// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	//type casting
	// ımplicit veride kayıp yok

	//explicit veride kayıp olduğunda
	//değişiklik olduğunda ortaya çıkar

	var tamsayi int = 5
	var reelsayi float64 = 6.7
	tamsayi = int(reelsayi)
	fmt.Println(tamsayi)
	// burada float değeri integer bir değer yaptik
	reelsayi=float64(tamsayi)
	// burada demek istediği floata çevir sonrasında 
	//yaz.

	// aslında baktığımızda bir değer kaybı söz konusu değil
	// değerler arası dönüşüm yaparken başında belirtmemizi istiyor.
	
}
