package zapp

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func MapAtoi(ss []string) []int {
	tmp := []int{}
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		tmp = append(tmp, i)
	}
	return tmp
}

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true

		}
	}
	return false
}

func uniqInt(list []int) []int {
	isExists := map[int]bool{}
	uniqList := []int{}
	for _, i := range list {
		if !isExists[i] {
			uniqList = append(uniqList, i)
			isExists[i] = true
		}
	}
	return uniqList
}

func MapItoA(ints []int) []string {
	tmp := []string{}
	for _, i := range ints {
		tmp = append(tmp, fmt.Sprintf("%d", i))
	}
	return tmp
}

func JoinIntSliceToString(ss []int) string {
	StrSlice := MapItoA(ss)
	return strings.Join(StrSlice, " ")
}

func SplitStringToIntSlice(s string) []int {
	strs := strings.Split(s, " ")
	IDs := MapAtoi(strs)
	sort.Ints(IDs)
	return IDs
}


func IntersectionInt(a, b []int) (ret []int) {
	m := make(map[int]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			ret = append(ret, item)
		}
	}
	return
}