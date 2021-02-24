package util

import (
	"math/rand"
	"time"
)

func Shuffle(slice []interface{}) {
	rand.Seed(time.Now().Unix())
	for len(slice) > 0 {
		n := len(slice)
		randIndex := rand.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
