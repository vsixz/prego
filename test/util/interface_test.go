package util

import (
	"fmt"
	"github.com/vsixz/prego/util"
	"math/rand"
	"testing"
)

func TestShuffle(t *testing.T) {
	slice := []interface{}{"a", "b", "c", "d", "e", "f"}
	//slice := []interface{}{1, 2, 3, 4, 5, 6}
	fmt.Printf("before func: %p\n", slice)
	util.Shuffle(slice)
	fmt.Printf("after func: %p\n", &slice)
	t.Log(slice)

	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	t.Log(slice)
}
