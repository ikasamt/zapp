package zapp_test

import (
	"github.com/ikasamt/zapp/zapp"
	"testing"
)

func Test_RuneCount(t *testing.T) {
	actual := zapp.RuneCount(`ほ`)
	expected := 1
	if actual != expected{
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	actual = zapp.RuneCount(`ほげ`)
	expected = 2
	if actual != expected{
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	actual = zapp.RuneCount(`aほげa`)
	expected = 4
	if actual != expected{
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}

	actual = zapp.RuneCount(`aa`)
	expected = 2
	if actual != expected{
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}


