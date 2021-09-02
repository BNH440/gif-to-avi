$gifRepeat = Read-Host -Prompt "Amount of times to repeat the gif (default: 1)"

$gifSpeed = Read-Host -Prompt "Video speed (higher is slower) (default: 1)"

./ffmpeg/bin/ffmpeg.exe -stream_loop $gifRepeat -i file.gif -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" -filter:v "setpts=$gifSpeed*PTS" file.avi