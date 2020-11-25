package maps

import "testing"

func TestSearc(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertEquals(t, got, expected)
	})

	t.Run("unknown work", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		assertDefinitions(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinitions(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinitions(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("Expected word %q to be deleted", word)
	}
}

func assertEquals(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func assertError(t *testing.T, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("expected error %q but got %q", expected, got)
	}
}

func assertDefinitions(t *testing.T, dictionary Dictionary, word, defintion string) {
	t.Helper()

	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("expected to find added word:", err)
	}

	assertEquals(t, got, defintion)
}
