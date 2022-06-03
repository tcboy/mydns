package util_random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomIntRange Returns an int >= min, < max
func RandomIntRange(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}
