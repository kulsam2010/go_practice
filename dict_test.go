package main

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "definition of test")
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "definition of test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("test1")
		assertError(t, err, ErrNotFound)
	})

	t.Run("existing word", func(t *testing.T) {
		err := dictionary.Add("test", "new defination")
		assertError(t, err, ErrExistingKey)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test update", func(t *testing.T) {
		word := "test"
		definition := "definition of test"
		newDefinition := "new definition of test"

		dictionary := Dictionary{}
		dictionary.Add(word, definition)
		dictionary.Update(word, newDefinition)
		got, _ := dictionary.Search(word)
		assertStrings(t, got, newDefinition)
	})

	t.Run("test update key doesn't exist ", func(t *testing.T) {
		word := "test"
		definition := "definition of test"
		newDefinition := "new definition of test"

		dictionary := Dictionary{}
		dictionary.Add(word, definition)
		err := dictionary.Update("abc", newDefinition)
		assertError(t, err, ErrNotFound)
	})

}

func TestDelete(t *testing.T) {
	t.Run("test delete", func(t *testing.T) {
		dictionary := Dictionary{}
		testKey := "test"
		dictionary.Add(testKey, "test def")
		err := dictionary.Delete(testKey)
		assetNoError(t, err)
		_, err = dictionary.Search(testKey)
		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
