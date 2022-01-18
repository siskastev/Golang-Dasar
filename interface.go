package main

import "fmt"

type Person struct {
	name string
}

type InfoInterface interface {
	DoInfo()
}

func Show(name string) InfoInterface {
	return &Person{name: name}
}

func (p *Person) DoInfo() {}

func main() {
	p := Person{"Siska"}
	fmt.Println(Show(p.name))

}
