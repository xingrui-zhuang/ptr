package ptr

import (
	"fmt"
	"testing"
)

func TestPtr_Elem(t *testing.T) {
	a := []byte{1, 2, 3, 4}
	fmt.Println(*Ptoa[byte](Atop(&a) + 1))
}
