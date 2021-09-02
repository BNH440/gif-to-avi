# gif-to-avi

## Commands

`ffmpeg -stream_loop 15 -i file.gif -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" -filter:v "setpts=4*PTS" file.avi`

`-stream_loop X`: X is the amount of times to loop the gif
`filter:v "setpts=X*PTS"`: X is the speed multiplier (higher is slower)
