package utils

func RemoveFromSliceByIndex[T any](slice []T, idx int) []T {
	return append(slice[:idx], slice[idx+1:]...)
}