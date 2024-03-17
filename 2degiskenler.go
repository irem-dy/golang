package main
import "fmt"
func main(){

var number int //var = variable(değişken) number isim olarak kullanıyoruz 
//değişken ismi ve değişken türü şeklinde tanımlarız.
number=15
fmt.Println(number)
// üstteki dolaylı yoldan bir tanımlamayken biz bunu doğrudanda 
//tanımlayabiliriz.
var number int=17 
//şeklinde dogrudanda tanımlayabiliriz.
//yalnız bu da aynı şekilde sürekli yazabileceğimiz bir 
//tür değil.

number:17 
// bunu yazarken var number int dememeliyiz number 17.2 de diyemeyiz.
//tam sayı olarak tanımladıysak ona göre bir sayı yazabiliriz.
//number:17 yazdığımzda sağdaki yazılan şey ne ise ona göre bir değişken 
//tanımlamasını istedik 
//direkt bu şekilde de tanımlayabiliriz.
// hangi türde olduğunu integer olduğunu üstte söylemiştik
// number:17 dediğimiz kısımda 

}