package word_count

import (
	"reflect"
	"testing"
)

func TestRemovePunct(t * testing.T) {

	s := "This, is! a. string? I; like' This` string."
	stIdeal := "This is a string I like This string"
	stTest := RemovePunct(s)


	if stTest != stIdeal {

		t.Errorf("expected string of %s, but got string %s\n", stIdeal, stTest)
	}
}

func TestWordCount(t * testing.T) {

	s := "This is a string? I like This string."

	mpTest := WordCount(s)
	mpTruth := make(map[string]int)

	mpTruth["this"] = 2
	mpTruth["is"] = 1
	mpTruth["a"] = 1
	mpTruth["string"] = 2
	mpTruth["i"] = 1
	mpTruth["like"] = 1

	if !reflect.DeepEqual(mpTest, mpTruth) {

		t.Errorf("expected map of %+v, but got map: %+v", mpTruth, mpTest)

	}

}