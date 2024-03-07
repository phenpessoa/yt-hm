package hashers

import (
	"strconv"
	"testing"
)

func TestStringHasher(t *testing.T) {
	const n = 1000
	for i := range n {
		hasher := StringHasher(strconv.FormatInt(int64(i), 10))
		hash := hasher.Hash()
		for range 100 {
			h := hasher.Hash()
			if h != hash {
				t.Errorf("different hashs, expected %d, got %d", hash, h)
			}
		}
	}
}
