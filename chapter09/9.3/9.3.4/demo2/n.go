package main

import "fmt"

type animal interface {
	saySomething()
}

//----------------------------------

type cat struct { }

func (p *cat) saySomething() {
	fmt.Println("喵喵")
}


type dog struct { }

func (p *dog) saySomething() {
	fmt.Println("汪汪")
}


type sheep struct { }

func (p *sheep) saySomething() {
	fmt.Println("咩咩")
}

//------------------------------------

func newAnimal(tp uint8) (res animal) {
	switch tp {
	case 1:
		res = &cat{}
	case 2:
		res = &dog{}
	case 3:
		res = &sheep{}
	default:
		res = nil
	}
	return
}

//---------------------------------

func main() {
	var a1 = newAnimal(1)
	var a2 = newAnimal(3)
	var a3 = newAnimal(5)

	if a1 != nil {
		a1.saySomething()
	}
	if a2 != nil {
		a2.saySomething()
	}
	if a3 != nil {
		a3.saySomething()
	}
}