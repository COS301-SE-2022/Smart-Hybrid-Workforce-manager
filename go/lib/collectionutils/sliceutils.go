package collectionutils

// SequentialSequence returns a sequence of sequential integers in the range [start, end)
// no numbers returned if start >= end
func SequentialSequence(start, end int) []int {
	if start >= end {
		return []int{}
	}
	seq := make([]int, end-start)
	for i := range seq {
		seq[i] = i + start
	}
	return seq
}

// This method finds the intersection between two integer slices
func IntSliceIntersection(slice1, slice2 []int) []int {
	// Gets the intersection of two integer slices
	intersectionMap := make(map[int]bool)
	for _, i := range slice1 {
		intersectionMap[i] = false
	}
	intersection := []int{}
	for _, i := range slice2 {
		if _, ok := intersectionMap[i]; ok {
			intersection = append(intersection, i)
		}
	}
	return intersection
}
