package virtDev

import (
	"fmt"
	"testing"
)

func TestM(t *testing.T) {
	// KeyNew()
	err := WheelNew()
	fmt.Println("error: ", err)
}
