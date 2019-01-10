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


func Test_IntIntersection(t *testing.T) {
	a := []int{1,2,3}
	b := []int{2,3,4}
	actual := zapp.IntIntersection(a, b)
	expected := []int{2,3}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}


func Test_IntDiff(t *testing.T) {
	a := []int{1,2,3}
	b := []int{2,3,4}
	actual := zapp.IntDiff(a, b)
	expected := []int{1}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}


	actual2 := zapp.IntDiff(b, a)
	expected2 := []int{4}
	if diff := deep.Equal(actual2, expected2); diff != nil {
		t.Errorf("got: %v\nwant: %v", actual2, expected2)
	}
}