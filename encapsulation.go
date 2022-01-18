package main

import (
	"fmt"
	"golang-dasar/Model"
)

type Movie struct {
	name, genre string
}

func main() {

	h := []Movie{
		{
			"Hotel Transylvania",
			"Family, Adventure, Cartoon",
		},
		{
			"Annabel",
			"Horror",
		},
		{
			"Matrix",
			"Action, Mystery",
		},
	}

	m := Model.Meta{}
	var total int = len(h)
	m.SetTotalRecord(total)
	m.GetTotalRecord()
	m.SetTotalPage(2)
	m.GetTotalPage()
	m.SetPage(1)
	m.GetPage()

	r := Model.Response{}
	r.Status = true
	r.Data = h
	r.Message = "OK"
	r.Meta = m

	fmt.Println(r)

}
