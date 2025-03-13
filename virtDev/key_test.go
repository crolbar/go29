package virtDev

import (
	"fmt"
	"testing"
)

func TestM(t *testing.T) {
	// KeyNew()
	kb, err := NewVirtKeyboard()
	if err != nil {
		fmt.Println(err)
		return
	}

	kb.PressKey(KEY_A)
	kb.PressKey(KEY_H)
	kb.PressKey(KEY_E)
	kb.PressKey(KEY_L)
	kb.PressKey(KEY_L)
	kb.PressKey(KEY_O)
	kb.PressKey(KEY_D)

	select{}

	kb.DestroyDev()
}
