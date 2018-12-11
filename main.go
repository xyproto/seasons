package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/xyproto/wallpaper"
)

func main() {
	imageFilename := ""
	if len(os.Args) > 1 {
		imageFilename = os.Args[1]
	} else {
		if matches, err := filepath.Glob("/usr/share/pixmaps/*.png"); err != nil {
			panic(err)
		} else {
			rand.Seed(time.Now().Unix())
			randomIndex := rand.Int() % len(matches)
			imageFilename = matches[randomIndex]
		}
	}
	absImageFilename, err := filepath.Abs(imageFilename)
	if err == nil {
		imageFilename = absImageFilename
	}
	if _, err := os.Stat(imageFilename); os.IsNotExist(err) {
		// File does not exist
		panic(err)
	}
	fmt.Println("Setting background image to: " + imageFilename)
	if err := wallpaper.Set(imageFilename); err != nil {
		panic(err)
	}
}
