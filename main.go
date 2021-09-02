package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("pwsh", "-nologo", "-noprofile")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, "Write-Host test")
		fmt.Fprintln(stdin, `
		$ErrorActionPreference = "Stop"

if (-NOT $args[0]) {
    Write-Error "No file path provided"
}

$file = Get-ChildItem $args[0]

$filePath = Split-Path $file

$outputFile = (Join-Path -Path $filePath -ChildPath (([io.path]::GetFileNameWithoutExtension($file)) + ".avi"))

$gifRepeat = Read-Host -Prompt "Amount of times to repeat the gif (default: 1)"

$gifSpeed = Read-Host -Prompt "Video speed (higher is slower) (default: 1)"

if (-NOT ([Microsoft.VisualBasic.Information]::IsNumeric($gifRepeat)) -OR -NOT ([Microsoft.VisualBasic.Information]::IsNumeric($gifSpeed))) {
    Write-Error "Non-numeric values"
}

./ffmpeg/bin/ffmpeg.exe -stream_loop $gifRepeat -i $file -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" -filter:v "setpts=$gifSpeed*PTS" $outputFile
		`)
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}
