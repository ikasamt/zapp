package zapp_test

import (
	"testing"

	"github.com/ikasamt/zapp/zapp"
)

func Test_RuneCount(t *testing.T) {
	actual := zapp.RuneCount(`😃`)
	expected := 1
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	actual = zapp.RuneCount(`ほげ`)
	expected = 2
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	actual = zapp.RuneCount(`aほげa`)
	expected = 4
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	actual = zapp.RuneCount(`aa`)
	expected = 2
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
