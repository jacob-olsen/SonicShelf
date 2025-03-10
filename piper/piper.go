package piper

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func MakeAudio(text string, output string, voice string, playspeed float64) {
	cmd := exec.Command("piper", "--model", voice, "--length_scale", strconv.FormatFloat(playspeed, 'f', 2, 64), "-f", output)
	cmd.Stdin = strings.NewReader(text)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
