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
