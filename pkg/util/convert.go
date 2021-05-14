package util

func ToIntArray(oa []float64, div float64) []int {
	r := make([]int, 0, len(oa))
	for _, i := range oa {
		r = append(r, int(i / div))
	}
	return r
}