package main

import "fmt"

type Human struct {
	name, address string
	age           int
}

type Teacher struct {
	NIP string
	Human
	Identity KTP
}

type Student struct {
	NIS string
	Human
	Identity KTP
}

type KTP struct {
	NIK, ImageURL string
}

func ChangeNameWithoutPointer(h Human) {
	h.name = "Stevani"
}

func ChangeNameWithPointer(h *Human) {
	h.name = "YOKA"
}

func main() {
	var value = Human{"Siska", "Kalimantan", 24}
	var memory = &value

	teacher := Teacher{"1234567890", *memory, KTP{"0987654321", "https://www.google.com/url?sa=i&url=https%3A%2F%2Ftravel.kompas.com%2Fread%2F2021%2F10%2F06%2F061000627%2Ffakta-menarik-dua-warna-bulu-panda-alat-kamuflase-dan-komunikasi%3Fpage%3Dall&psig=AOvVaw088MkX9e6WPjM0lT7Jes89&ust=1641824065140000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCMDb2LPtpPUCFQAAAAAdAAAAABAD"}}

	memory.name = "Stevani"
	student := Student{"0987654321", *memory, KTP{"1234567878", "https://www.google.com/url?sa=i&url=https%3A%2F%2Ftravel.kompas.com%2Fread%2F2021%2F10%2F06%2F061000627%2Ffakta-menarik-dua-warna-bulu-panda-alat-kamuflase-dan-komunikasi%3Fpage%3Dall&psig=AOvVaw088MkX9e6WPjM0lT7Jes89&ust=1641824065140000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCMDb2LPtpPUCFQAAAAAdAAAAABAD"}}

	fmt.Println(teacher)
	fmt.Println(student)

	//method
	ChangeNameWithPointer(&value)
	fmt.Println(value)
	ChangeNameWithoutPointer(value)
	fmt.Println(value)

}
