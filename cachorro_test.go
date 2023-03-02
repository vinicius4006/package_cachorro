package cachorro

import "testing"

func TestIdade(t *testing.T) {
	x := Idade(5)
	if x != 35 {
		t.Error("Expected x = 35")
	}
}
