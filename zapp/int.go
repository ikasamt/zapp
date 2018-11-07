package zapp

import "strconv"

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
