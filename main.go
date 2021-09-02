package main
// fart
import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	// Terminate if argument is not given
	if len(os.Args) != 2 {
		log.Fatal("Error: You must provide a file to convert")
	}

	filePath := os.Args[1]

	var gifRepeat float32 = 1
	var gifSpeed float32 = 1

	fmt.Print("Amount of times to repeat the gif (default: 1): ")
	fmt.Scan(&gifRepeat)

	fmt.Print("Video speed (higher is slower) (default: 1): ")
	fmt.Scan(&gifSpeed)

	fmt.Println(gifRepeat, gifSpeed)

	cmd := exec.Command("./ffmpeg/bin/ffmpeg.exe", " -stream_loop"+strconv.Itoa(gifRepeat)+"-i $file -movflags faststart -pix_fmt yuv420p -vf \"scale=trunc(iw/2)*2:trunc(ih/2)*2\" -filter:v \"setpts=$gifSpeed*PTS\" $outputFile")

	cmd.Run()

	// cmd := exec.Command("pwsh", "-nologo", "-noprofile")
	// stdin, err := cmd.StdinPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// go func() {
	// 	defer stdin.Close()
	// 	fmt.Fprintln(stdin, `Write-Host test1;
	// 	Write-Host test2;
	// 	Read-Host -Prompt "Amount of times to repeat the gif (default: 1)"`)

	// 	// 		fmt.Fprintln(stdin, `
	// 	// 		$ErrorActionPreference = "Stop"

	// 	// if (-NOT $args[0]) {
	// 	//     Write-Error "No file path provided"
	// 	// }

	// 	// $file = Get-ChildItem $args[0]

	// 	// $filePath = Split-Path $file

	// 	// $outputFile = (Join-Path -Path $filePath -ChildPath (([io.path]::GetFileNameWithoutExtension($file)) + ".avi"))

	// 	// $gifRepeat = Read-Host -Prompt "Amount of times to repeat the gif (default: 1)"

	// 	// $gifSpeed = Read-Host -Prompt "Video speed (higher is slower) (default: 1)"

	// 	// if (-NOT ([Microsoft.VisualBasic.Information]::IsNumeric($gifRepeat)) -OR -NOT ([Microsoft.VisualBasic.Information]::IsNumeric($gifSpeed))) {
	// 	//     Write-Error "Non-numeric values"
	// 	// }

	// 	// ./ffmpeg/bin/ffmpeg.exe -stream_loop $gifRepeat -i $file -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" -filter:v "setpts=$gifSpeed*PTS" $outputFile
	// 	// 		`)
	// }()
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", out)
}
