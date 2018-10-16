package zapp

func UniqString(src []string) []string {
	ret := make([]string, 0, len(src))
	srcMap := make(map[string]struct{}, len(src))
	for _, n := range src {
		if _, ok := srcMap[n]; !ok {
			srcMap[n] = struct{}{}
			ret = append(ret, n)
		}
	}
	return ret
}
