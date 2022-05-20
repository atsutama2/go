package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	p := Person{
		Name: "名前",
		Age:  20,
	}

	fmt.Printf("最初のp :%+v\n", p)

	p2 := p
	p2.Name = "高橋"
	p2.Age = 21

	// pではなく
	fmt.Printf("p2で木村に書き換えを行ったはずのp :%+v\n", p)

	p3 := &p
	p3.Name = "田中"
	p3.Age = 22

	fmt.Printf("p3で田中に書き換えを行ったp :%+v\n", p)

	name := "日村"
	namePoint := &name
	fmt.Printf("namePoint 1 :%v\n", namePoint)
	fmt.Printf("namePoint 2 :%v\n", *namePoint)
}
