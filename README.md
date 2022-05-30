# Gif To Avi

A program that converts gif files into avi's using ffmpeg and go.

## Commands

`ffmpeg -stream_loop 15 -i file.gif -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" -filter:v "setpts=4*PTS" file.avi`

`-stream_loop X`: X is the amount of times to loop the gif

`filter:v "setpts=X*PTS"`: X is the speed multiplier (higher is slower)

### Add ffmpeg binary

`go-bindata -o ffmpeg.go ffmpeg/...`

## Original Powershell Script

```powershell
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
```
