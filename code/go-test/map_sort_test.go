package test

import (
	"fmt"
	"sort"
	"testing"
)

func TestMapSort(t *testing.T) {

	basket := map[string]interface{}{"orange": "test", "apple": 7,
		"mango": 4, "strawberry": "name"}

	keys := make([]string, 0, len(basket))

	for k := range basket {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, basket[k])
	}
}

func TestName(t *testing.T) {
	aesKey := []byte{254, 217, 36, 108, 118, 29, 210, 115, 26, 71, 105, 99, 139, 14, 31, 100}
	aeslv := []byte{160, 222, 53, 79, 168, 40, 93, 62, 84, 33, 195, 182, 63, 138, 205, 30}
	fmt.Println(string(aesKey))
	fmt.Println(string(aeslv))
}
