package player

import (
	"bufio"
	"fmt"
	"os"
)

// Human player
type Human struct {
}

// Print _
func (h *Human) Print(output string) {
	fmt.Println(output)
}

// Input _
func (h *Human) Input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
