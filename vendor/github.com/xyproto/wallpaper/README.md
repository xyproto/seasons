# Wallpaper

A Go module for changing the desktop wallpaper, for any windowmanager.

If your windowmanager is not yet supported, please submit a pull request or file an issue.

## Example Use

Setting a random PNG image from `/usr/share/pixmaps/` as the background image:

```go
package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/xyproto/wallpaper"
)

func main() {
	imageFilename := ""
	if matches, err := filepath.Glob("/usr/share/pixmaps/*.png"); err != nil {
		panic(err)
	} else {
		rand.Seed(time.Now().Unix())
		randomIndex := rand.Int() % len(matches)
		imageFilename = matches[randomIndex]
	}
	fmt.Println("Setting background image to: " + imageFilename)
	if err := wallpaper.Set(imageFilename); err != nil {
		panic(err)
	}
}
```

## Installation of the `getwallpaper` utility

    go get -u github.com/xyproto/wallpaper/cmd/setwallpaper

## General Info

* License: MIT
* Version: 1.0.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
