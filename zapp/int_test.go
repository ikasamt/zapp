package zapp_test

import (
	"github.com/go-test/deep"
	"github.com/ikasamt/zapp/zapp"
	"testing"
)

func Test_MapAtoi(t *testing.T) {
	data := []string{`1`,`2`,`3`,`4`,`5`}
	actual := zapp.MapAtoi(data)
	expected := []int{1,2,3,4,5}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
func Test_MapItoA(t *testing.T) {
	b := []int{1,2,3,4,5}
	actual := zapp.MapItoA(b)
	expected := []string{`1`,`2`,`3`,`4`,`5`}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func Test_JoinIntSliceToString(t *testing.T) {
	src := []int{1,2,3,4,5}
	actual := zapp.JoinIntSliceToString(src)
	expected := `1 2 3 4 5`
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func Test_SplitStringToIntSlice(t *testing.T) {
	src:= `1 2 3 4 5 11 123`
	actual := zapp.SplitStringToIntSlice(src)
	expected := []int{1,2,3,4,5, 11, 123}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}