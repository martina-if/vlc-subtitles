package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/matcornic/addic7ed"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <file> [--launch-vlc | --launch ]\n", os.Args[0])
		return
	}
	filePath, err := videoFilePath()
	if err != nil {
		fmt.Println(err)
	}

	err = downloadSubtitle(filePath)
	if err != nil {
		fmt.Println(err)
	}

	if launchFlagEnabled() {
		launchVlc(filePath)
	}
}

func downloadSubtitle(filePath string) error {
	fileName := path.Base(filePath)

	addicted := addic7ed.New()
	_, subtitle, err := addicted.SearchBest(fileName, "English")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Download the subtitle to a given file name
	err = subtitle.DownloadTo(filePath + ".srt")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func launchFlagEnabled() bool {
	for _, argument := range os.Args[1:] {
		if argument == "--launch-vlc" || argument == "--launch" {
			return true
		}
	}
	return false
}

func videoFilePath() (string, error) {
	for _, argument := range os.Args[1:] {
		if !strings.HasPrefix(argument, "--") {
			return argument, nil
		}
	}
	return "", errors.New(fmt.Sprintf("No filename supplied. Usage: %s <file> [--launch-vlc | --launch ]", os.Args[0]))
}

func launchVlc(filePath string) {
	cmd := exec.Command("/usr/bin/vlc", filePath)
	err := cmd.Run()
	if err != nil {
		fmt.Errorf("%v\n", err)
	}
}
