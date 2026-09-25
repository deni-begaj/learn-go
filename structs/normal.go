package structs

import "fmt"

type Car struct {
	Color      string
	Brand      string
	EngineType string
	MaxSpeed   int16
	Position   int16
}

func (c Car) display() {
	fmt.Printf("This is %v, %v color\n", c.Brand, c.Color)
}

func (c *Car) drive(gas int16) {
	oldPos := c.Position
	c.Position += gas
	fmt.Printf("Car moved from pos %v to %v\n", oldPos, c.Position)
}

func (c *Car) reverse(gas int8) {
	gas16 := int16(gas)
	if c.Position == 0 {
		fmt.Println("Car already at start position, cannot move further back")
		return
	}
	newPos := c.Position - gas16
	if newPos < 0 {
		c.Position = 0
		fmt.Println("Car cannot move further back than start position")
		return
	}
	c.Position -= gas16
}

func StructsNormal() {
	fmt.Println("\n------### Structs > Normal ###------")
	defer fmt.Println("------------------------------------")

	bmw := Car{"blue", "BMW", "combustion", 215, 0}
	bmw.display()
	bmw.drive(5)
	bmw.drive(12)
	bmw.reverse(10)
	bmw.reverse(9)
}
