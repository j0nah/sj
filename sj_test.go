package sj

import (
	"fmt"
	"testing"
)

//func TestInvalidJson(t *testing.T) {
//	haystack := `
//	{
//		"foo" : {
//			"baz" : "bob"
//
//	}
//`
//	_, err := search("", haystack)
//
//	if err == nil {
//		t.Fatal(err.Error())
//	} else {
//		fmt.Println(err.Error())
//	}
//}
//
//func TestSimple(t *testing.T) {
//	haystack := `
//	{
//		"foo" : {
//			"baz" : "bob"
//		}
//	}
//`
//	needle := "baz"
//	expected := `{"baz":"bob"}`
//	actual, err := search(needle, haystack)
//
//	if err != nil {
//		t.Fatal(err.Error())
//	}
//
//	if expected != *actual {
//		t.Errorf("Expected: %s Got: %s \n", expected, *actual)
//	}
//}

func TestSimple(t *testing.T) {
	haystack := `
		{
			"baz":"bob",
			"foo": {
				"baz" : "gumble",
				"lumble" : {
					"baz" : {
						"hello" : "world"
					}
				}
			}
		}
`
	needle := "baz"
	val, _ := Search(needle, haystack)

	for _, v := range *val {
		fmt.Println(v)
	}

	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//
	//if expected != *actual {
	//	t.Errorf("Expected: %s Got: %s \n", expected, *actual)
	//}
}
