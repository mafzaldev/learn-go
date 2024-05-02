package main

import "fmt"

type animal interface {
	callHuman() string
}

type cat struct {
	color      string
	brainCells uint8
	human
}

type dog struct {
	color      string
	brainCells uint8
	human
}

type human struct {
	name string
}

// Cats are disrespectful
func (c cat) callHuman() string {
	return c.name + "!!!"
}

// Dogs are respectful
func (d dog) callHuman() string {
	return "Hey " + d.name + "."
}

func speak(a animal) {
	fmt.Println(a.callHuman())
}

func main() {
	var orangeCat cat = cat{"Orange", 0, human{"Karen"}}
	fmt.Printf("Color: %v, Brain Cells: %v, Human: %v \n", orangeCat.color, orangeCat.brainCells, orangeCat.name)
	var shibaInu dog = dog{"Orange", 5, human{"Bob"}}
	fmt.Printf("Color: %v, Brain Cells: %v, Human: %v \n", shibaInu.color, shibaInu.brainCells, shibaInu.name)
	speak(orangeCat)
	speak(shibaInu)
}
