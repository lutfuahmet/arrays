package arrays

import (
	"math"
)

// InArray checks if a value is in an array
func InArray[T comparable](src []T, item T) bool {
	for i := range src {
		if src[i] == item {
			return true
		}
	}
	return false
}

// ArrayChunk splits an array into chunks
func ArrayChunk[T any](src []T, size int) [][]T {
	if size < 1 {
		return nil
	}
	length := len(src)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var result [][]T
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		result = append(result, src[i*size:end])
		i++
	}
	return result
}

// ArrayUnique removes duplicate values in an array
func ArrayUnique[T any](src []T) []T {
	var result []T
	var elemMap = make(map[T]bool)
	for i := range src {
		if !elemMap[src[i]] {
			result = append(result, src[i])
			elemMap[src[i]] = true
		}
	}
	return result
}

// ArrayUniqueMerge merges two arrays and removes duplicate values
func ArrayUniqueMerge[T any](arrays ...[]T) []T {
	var result []T
	var elemMap = make(map[T]bool)
	for _, array := range arrays {
		for i := range array {
			if !elemMap[array[i]] {
				result = append(result, array[i])
				elemMap[array[i]] = true
			}
		}
	}
	return result
}

// ArrayDiff returns the values from src that are not present in other arrays
func ArrayDiff[T comparable](src []T, arrays ...[]T) []T {
	var elemMap = make(map[T]bool)
	for _, array := range arrays {
		for i := range array {
			elemMap[array[i]] = true
		}
	}
	var result []T
	for i := range src {
		if !elemMap[src[i]] {
			result = append(result, src[i])
		}
	}
	return result
}

// ArrayFilter filters the array with the callback function
func ArrayFilter[T any](src []T, filter func(item T) bool) []T {
	var result []T
	for i := range src {
		if filter(src[i]) {
			result = append(result, src[i])
		}
	}
	return result
}
