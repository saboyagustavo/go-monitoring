package utils

import "os"

func ClearInputBuffer() {
	var buffer [100]byte
	os.Stdin.Read(buffer[:])
}
