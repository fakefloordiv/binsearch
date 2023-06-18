package binsearch

import "testing"

func BenchmarkMap_Get(b *testing.B) {
	m := NewContainer[string]()
	m.SetMap(map[string]string{
		hello:        "world",
		something:    "i want",
		disallowed:   "this hashmap is gonna take significant space.",
		someMoreText: "idk what to write about here",
	})

	b.Run("Lookup Hello", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = m.Get(hello)
		}
	})

	b.Run("Lookup something", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = m.Get(something)
		}
	})

	b.Run("Lookup disallowed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = m.Get(disallowed)
		}
	})

	b.Run("Lookup someMoreText", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = m.Get(someMoreText)
		}
	})

	b.Run("Lookup short unknown key", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = m.Get("sm.")
		}
	})

	b.Run("Lookup long unknown key", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = m.Get("some key that doesn't really exists, but we don't care")
		}
	})
}

func BenchmarkStrhash(b *testing.B) {
	var (
		hello        = []byte("Hello")
		something    = []byte("something")
		disallowed   = []byte("I JUST NEEDED SOME LONG TEST STRINGS")
		someMoreText = []byte("SomeVeryVeryVeryVeryLooooooongKeyThatIsAlsoPossibleInRealProduction")
	)

	b.Run("Strhash hello", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strhash(hello)
		}
	})

	b.Run("Strhash something", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strhash(something)
		}
	})

	b.Run("Strhash disallowed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strhash(disallowed)
		}
	})

	b.Run("Strhash someMoreText", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strhash(someMoreText)
		}
	})
}
