package sj

import (
	"fmt"
	"testing"
)

func TestInvalidJson(t *testing.T) {
	haystack := `
	{
		"foo" : {
			"baz" : "bob"
		
	}
`
	_, err := search("", haystack)

	if err == nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(err.Error())
	}
}

func TestSimple(t *testing.T) {
	haystack := `
	{
		"foo" : {
			"baz" : "bob"
		}
	}
`
	needle := "baz"
	expected := `{"baz":"bob"}`
	actual, err := search(needle, haystack)

	if err != nil {
		t.Fatal(err.Error())
	}

	if expected != *actual {
		t.Errorf("Expected: %s Got: %s \n", expected, *actual)
	}
}
