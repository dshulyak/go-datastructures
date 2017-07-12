package ptrie

import (
	"fmt"
	"testing"
)

func TestPtrie(t *testing.T) {
	r1 := makeRoot()
	r2 := r1.insert("10")
	r3 := r2.insert("11")
	fmt.Println(r3.lookup(0), r3.lookup(1))
	r4 := r3.insert("13")
	fmt.Println(r4.lookup(0))
	r5 := r4.insert("15")
	fmt.Println(r5.lookup(3))
	r6 := r5.delete(3).delete(2)
	fmt.Println(r6.listVals(), r6.nei.depth)
}
