package main

import (
	"SonicShelf/meta"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("setup folders")
	SetupFolders()
	fmt.Println("setup databaser")
	meta.Setup()
}

func AddVoice(voicePath string, name string) {
	id := meta.AddVoice(name)
	folderPath := "./voice/" + strconv.Itoa(int(id))
	os.MkdirAll(folderPath, 0775)

	os.Rename(voicePath, folderPath+"/voice.onnx")
	os.Rename(voicePath+".json", folderPath+"/voice.onnx.json")
}
