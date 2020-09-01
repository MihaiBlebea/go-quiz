package player

import (
	"fmt"
)

// Computer player
type Computer struct {
	Answers []string
	counter int
}

// Print _
func (c *Computer) Print(output string) {
	fmt.Println(output)
}

// Input _
func (c *Computer) Input() (string, error) {
	defer c.incrementCounter()

	return c.Answers[c.counter], nil
}

func (c *Computer) incrementCounter() {
	c.counter++
}
