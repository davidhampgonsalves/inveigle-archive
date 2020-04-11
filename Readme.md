
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

