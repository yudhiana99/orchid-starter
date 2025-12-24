package common

func MoveElementFromTo[T any](s []T, from, to int) []T {
	if from < 0 || from >= len(s) || to < 0 || to >= len(s) || from == to {
		return s // or handle error
	}

	// Save the element
	elem := s[from]

	// Remove it
	s = append(s[:from], s[from+1:]...)

	// Insert it at the new position
	s = append(s[:to], append([]T{elem}, s[to:]...)...)

	return s
}
