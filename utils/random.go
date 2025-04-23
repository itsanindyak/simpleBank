package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.New(rand.NewSource(time.Microsecond.Microseconds()))
}

// return random number b/w max and min
func RandomInt(min int64,max int64) int64{
	return min + rand.Int63n(max-min+1)
}
// return random owner with 7 chracter
func RandomOwner() string{
	var sb strings.Builder
	l := len(alphabet)
	
	for range 7 {
		c:= alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()

}

func RandomMoney() int64{
	return RandomInt(0,1000)
}

