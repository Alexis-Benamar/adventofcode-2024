package utils

func ArrayCopy(orig []string) []string {
	newArray := make([]string, len(orig))
	copy(newArray, orig[:])
	return newArray
}