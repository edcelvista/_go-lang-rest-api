package lib

import (
	"math/rand"
)

func GetRandomInt() (randn int) {
	randn = rand.Intn(100)
	return
}
