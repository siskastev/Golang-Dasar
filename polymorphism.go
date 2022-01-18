package main

import "fmt"

type Smartphone interface {
	GetInfo()
}

type Xiomi struct {
	name, color, camera, storage string
}

type Oppo struct {
	name, color, camera, storage string
}

func (x *Xiomi) GetInfo() string {
	return "name" + x.name + "," + x.color + "," + x.camera + "," + x.storage
}

func (o *Oppo) GetInfo() string {
	return "name" + o.name + "," + o.color + "," + o.camera + "," + o.storage
}

func main() {
	o := Oppo{"Renno", "Red", "Kamera Depan/Belakang : 8MP/13MP+2MP", "RAM 8GB/ROM 16GB"}
	x := Xiomi{"Redmi", "Red", "Kamera Depan/Belakang : 16MP/12MP+5MP", "RAM 4GB/ROM 64GB"}

	fmt.Println(x.GetInfo())
	fmt.Println(o.GetInfo())

}
