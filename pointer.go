package main

import "fmt"

func main() {
	//activities 1
	//Buatlah sebuah variabel(b) yang menyimpan alamat memori dari variabel lain(a)
	//dan ubahlah nilai dari variabel tersebut
	var a int = 1
	var b *int = &a
	*b = 5
	//fmt.Println(a)
	fmt.Println(*b)

	//Buatlah variabel lain(c) yang mengambil nilai dari variabel pertama(a)
	//dan ubahlah nilainya, namun variabel pertama(a) tidak boleh ikut berubah
	var c int = a
	c = 4
	fmt.Println(a)

	//Tampilkan nilai kedua variabel (b,c) tersebut dalam console
	fmt.Println(*b)
	fmt.Println(c)

	//Tulis program yang dapat menukar dua bilangan bulat (x := 1; y := 2;
	//swap(&x, &y) akan menghasilkan x=2 dan y=1
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
