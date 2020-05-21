
# Setup

## Install Gallery Lib
`yarn install`
_Then move files from /dist to root (GH pages won't resolve the node_modules for some reason)_.

## Generate Site
Uses go-lang mustache templates.
`go run generate.go` or for dev: `fswatch -0 -drx ./ -Ie "\.html$" | xargs -0 -I {} go run generate.go`

## Photo Thumbnails
From photo folder:
```
find . -exec sh -c 'convert "$1" -resize "400^>" -gravity center -crop 200x200+0+0 -strip "$1.tn"' -- {} \;
```

## Transcode Mov to MP4
```
 ffmpeg -i "iconolatry part 2.mov" -crf 19 -vf yadif -tune film -preset slow -pix_fmt yuv420p -vcodec h264 -acodec mp2 "iconolatry part 2.mp4"
```

## Transcode Videos to Vp9 (too slow)
Based on Google's (recommended settings)[https://developers.google.com/media/vp9/settings/vod].

```
for file in *.mp4; do
  ffmpeg -i $file -threads 4 -vf scale=640x480 -b:v 750k -quality good -speed 0 -crf 33 -c:v libvpx-vp9 -c:a libopus $file.webm
done
```

## Canadian Forces
transcode each part via: `ffmpeg -i "canadian forces - roderick.avi" -crf 19 -vf yadif -tune film -preset slow -pix_fmt yuv420p -vcodec h264 -acodec mp2 "canadian forces - roderick.mp4"`

Then concat via:
list.txt:
`
file 'canadian forces - intro.mp4'
file 'canadian forces - richard.mp4'
file 'canadian forces - friends.mp4'
file 'canadian forces - roderick.mp4'
file 'canadian forces - david.mp4'
`

`ffmpeg -f concat -safe 0 -i list.txt -crf 20 -tune film -preset slow -pix_fmt yuv420p -vcodec h264 -acodec mp2 "canadian forces.mp4"`

