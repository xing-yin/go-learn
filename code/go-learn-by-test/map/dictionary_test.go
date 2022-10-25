package main

import "testing"

/**
 * @Author: Alan Yin
 * @Date: 2022/10/25
 * @Desc:
 */

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is a test"

		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertValue(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is a test"
		dictionary := Dictionary{key: value}

		err := dictionary.Add(key, "new test value")

		assertError(t, err, ErrWordExists)
		assertValue(t, dictionary, key, value)
	})
}

func assertValue(t *testing.T, dictionary Dictionary, key string, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added value:", err)
	}
	if value != got {
		t.Errorf("got %s want %s", got, value)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
