package main

import "os"

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
func FolderExist(path string) bool {
	return FileExist(path)
}
