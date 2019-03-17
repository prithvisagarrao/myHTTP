package main

import (
	"testing"
)

func TestCheckURLs(t *testing.T) {

	testUrl:= []string{"http://www.google.com", "twitter.com"}
	expect := []string{"http://www.google.com", "http://twitter.com"}
	got := checkURLs(testUrl)

	for i:=0;i<len(testUrl); i++{
		if expect[i] != got[i]{

			t.Errorf("got '%s' expected '%s'",got[i],expect[i])
		}
	}

}

