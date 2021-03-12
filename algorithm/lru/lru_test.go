package lru

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLRU(t *testing.T) {

	lruCache := NewLRUCache(10)

	for i := 0; i < 10; i++ {
		j := fmt.Sprintf("%d", i)
		lruCache.Put(j, j)
	}

	require.Equal(t, "9", lruCache.Get("9"))

	lruCache.Put("10", "10")
	require.Equal(t, "10", lruCache.Get("10"))

	require.Equal(t, "10", lruCache.head.next.value)

	lruCache.Put("0", "100")
	require.Equal(t, "100", lruCache.head.next.value)
	require.Equal(t, "100", lruCache.Get("0"))

}
