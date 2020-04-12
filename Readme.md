
# Setup

## Install Gallery Lib
`yarn install`

## Generate Site
Uses go-lang mustache templates.
`go run generate.go` or for dev: `fswatch -0 -drx ./ -Ie "\.html$" | xargs -0 -I {} go run generate.go`

## Photo Thumbnails
From photo folder:
```
find . -exec sh -c 'convert "$1" -resize "400^>" -gravity center -crop 200x200+0+0 -strip "$1.tn"' -- {} \;
```

# Transcoding videos
Based on Google's (recommended settings)[https://developers.google.com/media/vp9/settings/vod].

for file in *.mp4; do
  ffmpeg -i $file -threads 4 -vf scale=640x480 -b:v 750k -quality good -speed 0 -crf 33 -c:v libvpx-vp9 -c:a libopus $file.webm
done
