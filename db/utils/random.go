package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int32 {
	return int32(min + rand.Int63n(max-min+1))
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomList(n int) []int32 {
	items := make([]int32, n)

	for i := 0; i < n; i++ {
		k := int32(n)
		items = append(items, rand.Int31n(k-1))
	}
	return items
}

func RandomTime() time.Time {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	return t
}
