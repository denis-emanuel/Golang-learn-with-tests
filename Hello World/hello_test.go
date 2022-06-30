package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper() // when the test fails the line number reported will be in our function call rather than inside our test helper
				   // comment to see what happens. 
				   // If commented error line is 12
				   // If not commented error line is wherever the error happened.
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello to people", func(t *testing.T){
		got := Hello("Chris", "")
		want := "Hello, Chris"
	
		assertCorrectMessage(t, got, want);
	})
	t.Run("Say 'Hello, World' by default is no name is specified", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want);
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Rita", "Spanish")
		want := "Hola, Rita";

		assertCorrectMessage(t, got, want);
	})
}