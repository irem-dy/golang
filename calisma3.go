// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	//kullanıcıdan girdi nasıl alırız ona bakalım
	var yas float64
	fmt.Print("bir sayi giriniz:")
	//fmt.Scan() fonksiyonuyla kullanıcıdan girdi alırız.
	fmt.Scan(&yas)
	fmt.Println(yas)
	// burada output imlecinin yan tarafta olmasını
	//istiyorsam println yerine print kullanabilirim.

}
