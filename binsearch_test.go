package binsearch

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	hello        = "Hello"
	something    = "something"
	disallowed   = "I JUST NEEDED SOME LONG TEST STRINGS"
	someMoreText = "SomeVeryVeryVeryVeryLooooooongKeyThatIsAlsoPossibleInRealProduction"
)

func TestSearch(t *testing.T) {
	container := NewContainer[string]()
	values := map[string]string{
		hello:        "world",
		something:    "i want",
		disallowed:   "this hashmap is gonna take significant space.",
		someMoreText: "idk what to write about here",
	}
	container.SetMap(values)

	t.Run("existing values", func(t *testing.T) {
		for key, value := range values {
			actual, found := container.Get(key)
			if assert.Truef(t, found, "value exists, but not found: %s", key) {
				assert.Equalf(t, value, actual, "got wrong value with key %s", key)
			}
		}
	})

	t.Run("non-existing values", func(t *testing.T) {
		_, found := container.Get("some random non-existing text here")
		require.Falsef(t, found, "the key must not be presented in the container")
	})
}

func TestArrInsert(t *testing.T) {
	t.Run("no reallocation", func(t *testing.T) {
		testset := []int{1, 2, 3, 4}
		numbers := make([]int, len(testset), len(testset)+1)
		copy(numbers, testset)
		require.Equal(t, []int{1, 2, 3, 4}, numbers)
		require.True(t, cap(numbers)-len(numbers) == 1)
		numbers = arrInsert(2, 5, numbers)
		require.Equal(t, []int{1, 2, 5, 3, 4}, numbers)
	})

	t.Run("yes reallocation", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		numbers = arrInsert(2, 5, numbers)
		require.Equal(t, []int{1, 2, 5, 3, 4}, numbers)
	})

	t.Run("into empty slice", func(t *testing.T) {
		arr := arrInsert(0, 1, []int{})
		require.Equal(t, []int{1}, arr)
	})
}
