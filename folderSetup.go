package main

import "os"

func SetupFolders() {
	if FolderExist("./temp") {
		os.RemoveAll("./temp")
	}
	os.Mkdir("./temp", 0775)

	if !FolderExist("./voice") {
		os.Mkdir("./voice", 0775)
	}
}
