package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	ks, vs := make([]int, 0, len(input)), make([]string, 0, len(input))
	for k := range input {
		ks = append(ks, k)
	}
	sort.Sort(sort.IntSlice(ks))
	for _, v := range ks {
		vs = append(vs, input[v])
	}
	return vs
}
