package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func getffmpeg() {
	// Retrieve exe from the go file
	data, err := Asset("ffmpeg/ffmpeg.exe")

	if err != nil {
		log.Fatal(err)
	}

	// Write the exe to cache
	os.WriteFile("./ffmpeg.exe", data, 0644)
}

func main() {
	// Get OS Cache
	cachedir, err := os.UserCacheDir()

	// FFmpeg location
	ffmpegLoc := path.Join(cachedir, "ffmpeg.exe")

	if _, err := os.Stat(ffmpegLoc); os.IsNotExist(err) {
		fmt.Println("Cache file does not exist, creating...")
		getffmpeg()
	}

	// Terminate if argument is not given
	if len(os.Args) != 2 {
		log.Fatal("Error: You must provide a file to convert")
	}

	// Get filepath from the program startup arguments
	filePath := os.Args[1]

	// Make the new filepath identical to the old one but with .avi instead of the previous extension
	newFilePath := strings.TrimSuffix(filePath, path.Ext(filePath)) + ".avi"

	// Set default values of 1
	var gifRepeat int = 1
	var gifSpeed string = "1.00"

	fmt.Print("Amount of times to repeat the gif (default: 1): ")
	fmt.Scanln(&gifRepeat)

	fmt.Print("Video speed (higher is slower) (default: 1): ")
	fmt.Scanln(&gifSpeed)

	fmt.Println("Starting Conversion...")
	command := strings.Fields(ffmpegLoc + " -y -stream_loop " + strconv.Itoa(gifRepeat) + " -i " + filePath + " -movflags faststart -pix_fmt yuv420p -vf setpts=" + gifSpeed + "*PTS " + newFilePath)
	cmd := exec.Command(command[0], command[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmderror := cmd.Run()
	if cmderror != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		os.Remove(newFilePath)
		return
	}
	fmt.Println("Conversion Complete! File located at: " + newFilePath)
}
