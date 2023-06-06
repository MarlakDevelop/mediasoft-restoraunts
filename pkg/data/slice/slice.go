package slice

func Merge[T any](slices ...[]T) []T {
	var result []T
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

func MergeArray[T any](slices [][]T) []T {
	var result []T
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

func MakeElemsUnique[T any](slice []T, cb func(index int, value T, slice []T) bool) []T {
	var result []T
	for _, elem := range slice {
		if IndexOf[T](slice, func(index int, value T, slice []T) bool {
			return cb(index, value, slice)
		}) == -1 {
			result = append(result, elem)
		}
	}
	return result
}

func IndexOf[T any](slice []T, cb func(index int, value T, slice []T) bool) int {
	for i, elem := range slice {
		if cb(i, elem, slice) {
			return i
		}
	}
	return -1
}

func Filter[T any](slice []T, cb func(index int, value T, slice []T) bool) []T {
	var result []T
	for i, elem := range slice {
		if cb(i, elem, slice) {
			result = append(result, elem)
		}
	}
	return result
}

func Map[OldT any, NewT any](slice []OldT, cb func(index int, value OldT, slice []OldT) NewT) []NewT {
	var result []NewT
	for i, elem := range slice {
		result = append(result, cb(i, elem, slice))
	}
	return result
}

func HasDuplicates[T string | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint](slice []T) bool {
	var inSlice []T
	for _, elem := range slice {
		if IndexOf[T](inSlice, func(index int, value T, slice []T) bool {
			return value == elem
		}) != -1 {
			return true
		} else {
			inSlice = append(inSlice, elem)
		}
	}
	return false
}

func IsOneInOther[T string | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint](slice []T, otherSlice []T) bool {
	for _, elem := range slice {
		if IndexOf[T](otherSlice, func(index int, value T, slice []T) bool {
			return value == elem
		}) == -1 {
			return false
		}
	}
	return true
}
